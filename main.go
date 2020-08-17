package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

	pb "example.com/grpcLB/rpc"
)

// DialGrpc -
// TODO - Add timeout and retry logic for dialling to grpc
func dialGrpc(grpcContext context.Context, serviceAddress string) (*grpc.ClientConn, error) {

	grpcContextWithTimeout, cancel := context.WithTimeout(grpcContext,
		time.Duration(5)*time.Minute)
	defer cancel()
	grpcClientOpts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	grpcClientOpts = append(grpcClientOpts, grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                time.Duration(20) * time.Second,
		Timeout:             time.Duration(20) * time.Millisecond,
		PermitWithoutStream: false,
	}))
	// 1. Get connection to schematics job - GRPC Server
	fmt.Println("Attempting to get grpc connection ", "Connection Parameters:", serviceAddress)
	return grpc.DialContext(grpcContextWithTimeout, serviceAddress, grpcClientOpts...)
}

func main() {

	serviceAddress := os.Args[1]
	grpcContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	grpcContext = metadata.AppendToOutgoingContext(grpcContext, "ContextKeyRequestID", "dummytest_forstopcancel")

	grpcCallOption := grpc.WaitForReady(true)

	grpcConn, grpcErr := dialGrpc(grpcContext, serviceAddress)

	if grpcErr != nil {
		fmt.Println("Grpc conn error:", grpcErr)
	}
	fmt.Println("GRPC Connection to Job details:", "target", grpcConn.Target(),
		"state", grpcConn.GetState().String())

	fmt.Println("GRPC Connection to Job details:", "target",
		"state", grpcConn)
	addr, ok := peer.FromContext(grpcContext)
	// fmt.Println(addr.Addr, addr.AuthInfo, ok)
	fmt.Println(addr, ok)
	// Get job GRPC Client to invoke GRPC Server
	jobGrpcClient := pb.NewActionServiceClient(grpcConn)
	orchestratorID := pb.OrchestratorId{ID: "test_handler"}

	//TODO: Have to put a retry logic in case of failure
	fmt.Println("orchestratorID", orchestratorID)
	initActionRequest := pb.InitActionRequest{OrchestratorID: &orchestratorID, RequestID: "dummytest_forstopcancel"}
	initRes, initErr := jobGrpcClient.InitJob(grpcContext, &initActionRequest, grpcCallOption)
	fmt.Println("Init job response:", initRes)
	if initErr != nil {
		fmt.Println("Error while initializing the job", initErr)
	}

	if initRes != nil && initRes.GetError().GetCode() == pb.ErrorCode_Occupied {
		fmt.Println("The job is occupied, closing this connection and trying again for %dth time")
	}

	if initRes == nil || initRes.GetError().GetCode() == pb.ErrorCode_Failed {
		fmt.Println("Error while initializing the job")
	}
}
