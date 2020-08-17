module example.com/grpcLB

go 1.14

require (
	example.com/grpcLB/rpc v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.29.1
)

replace example.com/grpcLB/rpc => ./rpc
