// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrorCode int32

const (
	ErrorCode_Unknown  ErrorCode = 0
	ErrorCode_Ok       ErrorCode = 1
	ErrorCode_Failed   ErrorCode = 2
	ErrorCode_Occupied ErrorCode = 3
)

var ErrorCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
	3: "Occupied",
}

var ErrorCode_value = map[string]int32{
	"Unknown":  0,
	"Ok":       1,
	"Failed":   2,
	"Occupied": 3,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

type WorkspaceSourceType int32

const (
	WorkspaceSourceType_BLUEPRINT          WorkspaceSourceType = 0
	WorkspaceSourceType_SCHEMATICS_SANDBOX WorkspaceSourceType = 1
	WorkspaceSourceType_SCHEMATICS         WorkspaceSourceType = 2
)

var WorkspaceSourceType_name = map[int32]string{
	0: "BLUEPRINT",
	1: "SCHEMATICS_SANDBOX",
	2: "SCHEMATICS",
}

var WorkspaceSourceType_value = map[string]int32{
	"BLUEPRINT":          0,
	"SCHEMATICS_SANDBOX": 1,
	"SCHEMATICS":         2,
}

func (x WorkspaceSourceType) String() string {
	return proto.EnumName(WorkspaceSourceType_name, int32(x))
}

func (WorkspaceSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

type STATE int32

const (
	STATE_UNKNOWN  STATE = 0
	STATE_PENDING  STATE = 1
	STATE_RUNNING  STATE = 2
	STATE_READY    STATE = 3
	STATE_COMPLETE STATE = 4
	STATE_FAIL     STATE = 5
	STATE_TIMEOUT  STATE = 6
	STATE_STOPPED  STATE = 7
)

var STATE_name = map[int32]string{
	0: "UNKNOWN",
	1: "PENDING",
	2: "RUNNING",
	3: "READY",
	4: "COMPLETE",
	5: "FAIL",
	6: "TIMEOUT",
	7: "STOPPED",
}

var STATE_value = map[string]int32{
	"UNKNOWN":  0,
	"PENDING":  1,
	"RUNNING":  2,
	"READY":    3,
	"COMPLETE": 4,
	"FAIL":     5,
	"TIMEOUT":  6,
	"STOPPED":  7,
}

func (x STATE) String() string {
	return proto.EnumName(STATE_name, int32(x))
}

func (STATE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

type TFSTATE int32

const (
	TFSTATE_TFFAILURE TFSTATE = 0
	TFSTATE_ACTIVE    TFSTATE = 1
	TFSTATE_INACTIVE  TFSTATE = 2
)

var TFSTATE_name = map[int32]string{
	0: "TFFAILURE",
	1: "ACTIVE",
	2: "INACTIVE",
}

var TFSTATE_value = map[string]int32{
	"TFFAILURE": 0,
	"ACTIVE":    1,
	"INACTIVE":  2,
}

func (x TFSTATE) String() string {
	return proto.EnumName(TFSTATE_name, int32(x))
}

func (TFSTATE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

type InitActionRequest struct {
	OrchestratorID       *OrchestratorId     `protobuf:"bytes,1,opt,name=OrchestratorID,proto3" json:"OrchestratorID,omitempty"`
	RequestID            string              `protobuf:"bytes,2,opt,name=RequestID,proto3" json:"RequestID,omitempty"`
	WorkspaceSource      WorkspaceSourceType `protobuf:"varint,3,opt,name=WorkspaceSource,proto3,enum=rpc.WorkspaceSourceType" json:"WorkspaceSource,omitempty"`
	WorkspaceCRN         string              `protobuf:"bytes,4,opt,name=WorkspaceCRN,proto3" json:"WorkspaceCRN,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *InitActionRequest) Reset()         { *m = InitActionRequest{} }
func (m *InitActionRequest) String() string { return proto.CompactTextString(m) }
func (*InitActionRequest) ProtoMessage()    {}
func (*InitActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *InitActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitActionRequest.Unmarshal(m, b)
}
func (m *InitActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitActionRequest.Marshal(b, m, deterministic)
}
func (m *InitActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitActionRequest.Merge(m, src)
}
func (m *InitActionRequest) XXX_Size() int {
	return xxx_messageInfo_InitActionRequest.Size(m)
}
func (m *InitActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitActionRequest proto.InternalMessageInfo

func (m *InitActionRequest) GetOrchestratorID() *OrchestratorId {
	if m != nil {
		return m.OrchestratorID
	}
	return nil
}

func (m *InitActionRequest) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

func (m *InitActionRequest) GetWorkspaceSource() WorkspaceSourceType {
	if m != nil {
		return m.WorkspaceSource
	}
	return WorkspaceSourceType_BLUEPRINT
}

func (m *InitActionRequest) GetWorkspaceCRN() string {
	if m != nil {
		return m.WorkspaceCRN
	}
	return ""
}

type InitActionResponse struct {
	JobID                *ActionServiceId `protobuf:"bytes,1,opt,name=JobID,proto3" json:"JobID,omitempty"`
	Status               *Status          `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Error                *ErrorDetails    `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *InitActionResponse) Reset()         { *m = InitActionResponse{} }
func (m *InitActionResponse) String() string { return proto.CompactTextString(m) }
func (*InitActionResponse) ProtoMessage()    {}
func (*InitActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *InitActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitActionResponse.Unmarshal(m, b)
}
func (m *InitActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitActionResponse.Marshal(b, m, deterministic)
}
func (m *InitActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitActionResponse.Merge(m, src)
}
func (m *InitActionResponse) XXX_Size() int {
	return xxx_messageInfo_InitActionResponse.Size(m)
}
func (m *InitActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitActionResponse proto.InternalMessageInfo

func (m *InitActionResponse) GetJobID() *ActionServiceId {
	if m != nil {
		return m.JobID
	}
	return nil
}

func (m *InitActionResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *InitActionResponse) GetError() *ErrorDetails {
	if m != nil {
		return m.Error
	}
	return nil
}

type ErrorDetails struct {
	Message              string    `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 ErrorCode `protobuf:"varint,2,opt,name=Code,proto3,enum=rpc.ErrorCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ErrorDetails) Reset()         { *m = ErrorDetails{} }
func (m *ErrorDetails) String() string { return proto.CompactTextString(m) }
func (*ErrorDetails) ProtoMessage()    {}
func (*ErrorDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *ErrorDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorDetails.Unmarshal(m, b)
}
func (m *ErrorDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorDetails.Marshal(b, m, deterministic)
}
func (m *ErrorDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDetails.Merge(m, src)
}
func (m *ErrorDetails) XXX_Size() int {
	return xxx_messageInfo_ErrorDetails.Size(m)
}
func (m *ErrorDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDetails proto.InternalMessageInfo

func (m *ErrorDetails) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorDetails) GetCode() ErrorCode {
	if m != nil {
		return m.Code
	}
	return ErrorCode_Unknown
}

type OrchestratorId struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrchestratorId) Reset()         { *m = OrchestratorId{} }
func (m *OrchestratorId) String() string { return proto.CompactTextString(m) }
func (*OrchestratorId) ProtoMessage()    {}
func (*OrchestratorId) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *OrchestratorId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrchestratorId.Unmarshal(m, b)
}
func (m *OrchestratorId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrchestratorId.Marshal(b, m, deterministic)
}
func (m *OrchestratorId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrchestratorId.Merge(m, src)
}
func (m *OrchestratorId) XXX_Size() int {
	return xxx_messageInfo_OrchestratorId.Size(m)
}
func (m *OrchestratorId) XXX_DiscardUnknown() {
	xxx_messageInfo_OrchestratorId.DiscardUnknown(m)
}

var xxx_messageInfo_OrchestratorId proto.InternalMessageInfo

func (m *OrchestratorId) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type ActionServiceId struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionServiceId) Reset()         { *m = ActionServiceId{} }
func (m *ActionServiceId) String() string { return proto.CompactTextString(m) }
func (*ActionServiceId) ProtoMessage()    {}
func (*ActionServiceId) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{4}
}

func (m *ActionServiceId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionServiceId.Unmarshal(m, b)
}
func (m *ActionServiceId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionServiceId.Marshal(b, m, deterministic)
}
func (m *ActionServiceId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionServiceId.Merge(m, src)
}
func (m *ActionServiceId) XXX_Size() int {
	return xxx_messageInfo_ActionServiceId.Size(m)
}
func (m *ActionServiceId) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionServiceId.DiscardUnknown(m)
}

var xxx_messageInfo_ActionServiceId proto.InternalMessageInfo

func (m *ActionServiceId) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Status struct {
	State                STATE    `protobuf:"varint,1,opt,name=State,proto3,enum=rpc.STATE" json:"State,omitempty"`
	TfState              TFSTATE  `protobuf:"varint,2,opt,name=TfState,proto3,enum=rpc.TFSTATE" json:"TfState,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{5}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetState() STATE {
	if m != nil {
		return m.State
	}
	return STATE_UNKNOWN
}

func (m *Status) GetTfState() TFSTATE {
	if m != nil {
		return m.TfState
	}
	return TFSTATE_TFFAILURE
}

func init() {
	proto.RegisterEnum("rpc.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("rpc.WorkspaceSourceType", WorkspaceSourceType_name, WorkspaceSourceType_value)
	proto.RegisterEnum("rpc.STATE", STATE_name, STATE_value)
	proto.RegisterEnum("rpc.TFSTATE", TFSTATE_name, TFSTATE_value)
	proto.RegisterType((*InitActionRequest)(nil), "rpc.InitActionRequest")
	proto.RegisterType((*InitActionResponse)(nil), "rpc.InitActionResponse")
	proto.RegisterType((*ErrorDetails)(nil), "rpc.ErrorDetails")
	proto.RegisterType((*OrchestratorId)(nil), "rpc.OrchestratorId")
	proto.RegisterType((*ActionServiceId)(nil), "rpc.ActionServiceId")
	proto.RegisterType((*Status)(nil), "rpc.Status")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x6e, 0xd2, 0xaf, 0xe5, 0x6d, 0x97, 0x79, 0xef, 0xd0, 0xa8, 0x10, 0x87, 0x12, 0x24, 0x98,
	0x7a, 0xd8, 0x21, 0xdc, 0x80, 0x4b, 0xd6, 0xa4, 0xe0, 0xd1, 0x26, 0x95, 0xe3, 0x32, 0x38, 0xa1,
	0x2e, 0x35, 0x50, 0x75, 0xaa, 0x83, 0x93, 0x82, 0xf8, 0x13, 0xfc, 0x3c, 0x7e, 0x0f, 0xb2, 0xd3,
	0x7d, 0x75, 0xdc, 0xf2, 0x7c, 0xf8, 0xf5, 0xf3, 0xc4, 0x36, 0x38, 0x2a, 0xcf, 0x4e, 0x73, 0x25,
	0x4b, 0x89, 0x75, 0x95, 0x67, 0xde, 0x5f, 0x0b, 0x0e, 0xe9, 0x7a, 0x59, 0x06, 0x59, 0xb9, 0x94,
	0x6b, 0x26, 0x7e, 0x6c, 0x44, 0x51, 0xe2, 0x1b, 0x70, 0x13, 0x95, 0x7d, 0x17, 0x45, 0xa9, 0xe6,
	0xa5, 0x54, 0x34, 0xec, 0x59, 0x7d, 0xeb, 0xa4, 0xe3, 0x1f, 0x9d, 0xea, 0xe5, 0xf7, 0xa4, 0x05,
	0xdb, 0xb1, 0xe2, 0x53, 0x70, 0xb6, 0x73, 0x68, 0xd8, 0xb3, 0xfb, 0xd6, 0x89, 0xc3, 0x6e, 0x09,
	0x3c, 0x83, 0x83, 0x0b, 0xa9, 0x56, 0x45, 0x3e, 0xcf, 0x44, 0x2a, 0x37, 0x2a, 0x13, 0xbd, 0x7a,
	0xdf, 0x3a, 0x71, 0xfd, 0x9e, 0x99, 0xbd, 0xa3, 0xf1, 0xdf, 0xb9, 0x60, 0xbb, 0x0b, 0xd0, 0x83,
	0xee, 0x0d, 0x35, 0x64, 0x71, 0xaf, 0x61, 0x36, 0xb9, 0xc7, 0x79, 0x7f, 0x2c, 0xc0, 0xbb, 0xc5,
	0x8a, 0x5c, 0xae, 0x0b, 0x81, 0x03, 0x68, 0x9e, 0xcb, 0xcb, 0x9b, 0x42, 0x8f, 0xcc, 0xa6, 0x95,
	0x27, 0x15, 0xea, 0xe7, 0x32, 0x13, 0x74, 0xc1, 0x2a, 0x0b, 0x3e, 0x87, 0x56, 0x5a, 0xce, 0xcb,
	0x4d, 0x61, 0x5a, 0x74, 0xfc, 0x8e, 0x31, 0x57, 0x14, 0xdb, 0x4a, 0xf8, 0x12, 0x9a, 0x91, 0x52,
	0x52, 0x99, 0x16, 0x1d, 0xff, 0xd0, 0x78, 0x0c, 0x13, 0x8a, 0x72, 0xbe, 0xbc, 0x2a, 0x58, 0xa5,
	0x7b, 0x63, 0xe8, 0xde, 0xa5, 0xb1, 0x07, 0xed, 0x89, 0x28, 0x8a, 0xf9, 0x37, 0x61, 0xb2, 0x38,
	0xec, 0x1a, 0xa2, 0x07, 0x8d, 0xa1, 0x5c, 0x08, 0xb3, 0xab, 0xeb, 0xbb, 0xb7, 0x13, 0x35, 0xcb,
	0x8c, 0xe6, 0xf5, 0x77, 0x4e, 0x68, 0x81, 0x2e, 0xd8, 0xdb, 0x5a, 0x0e, 0xb3, 0x69, 0xe8, 0x3d,
	0x83, 0x83, 0x9d, 0x5e, 0x0f, 0x2c, 0xec, 0xba, 0x20, 0xf6, 0xa1, 0xa9, 0xbf, 0xaa, 0x28, 0xae,
	0x0f, 0x55, 0x53, 0x1e, 0xf0, 0x88, 0x55, 0x02, 0xbe, 0x80, 0x36, 0xff, 0x5a, 0x79, 0xaa, 0x5c,
	0x5d, 0xe3, 0xe1, 0xa3, 0xca, 0x75, 0x2d, 0x0e, 0x5e, 0x83, 0x73, 0x93, 0x15, 0x3b, 0xd0, 0x9e,
	0xad, 0x57, 0x6b, 0xf9, 0x6b, 0x4d, 0x6a, 0xd8, 0x02, 0x3b, 0x59, 0x11, 0x0b, 0x01, 0x5a, 0xa3,
	0xf9, 0xf2, 0x4a, 0x2c, 0x88, 0x8d, 0x5d, 0xd8, 0x4b, 0xb2, 0x6c, 0x93, 0x2f, 0xc5, 0x82, 0xd4,
	0x07, 0x63, 0x38, 0xfa, 0xcf, 0xf9, 0xe3, 0x3e, 0x38, 0x67, 0xe3, 0x59, 0x34, 0x65, 0x34, 0xe6,
	0xa4, 0x86, 0xc7, 0x80, 0xe9, 0xf0, 0x7d, 0x34, 0x09, 0x38, 0x1d, 0xa6, 0x5f, 0xd2, 0x20, 0x0e,
	0xcf, 0x92, 0x4f, 0xc4, 0x42, 0x17, 0xe0, 0x96, 0x27, 0xf6, 0x60, 0x05, 0x4d, 0x93, 0xcd, 0xa4,
	0x88, 0x3f, 0xc4, 0xc9, 0x45, 0x4c, 0x6a, 0x1a, 0x4c, 0xa3, 0x38, 0xa4, 0xf1, 0x3b, 0x62, 0x69,
	0xc0, 0x66, 0x71, 0xac, 0x81, 0x8d, 0x0e, 0x34, 0x59, 0x14, 0x84, 0x9f, 0x49, 0x5d, 0xc7, 0x1a,
	0x26, 0x93, 0xe9, 0x38, 0xe2, 0x11, 0x69, 0xe0, 0x1e, 0x34, 0x46, 0x01, 0x1d, 0x93, 0xa6, 0xf6,
	0x73, 0x3a, 0x89, 0x92, 0x19, 0x27, 0x2d, 0x0d, 0x52, 0x9e, 0x4c, 0xa7, 0x51, 0x48, 0xda, 0x03,
	0x1f, 0xda, 0xdb, 0x5f, 0xa1, 0xe3, 0xf2, 0x91, 0x5e, 0x30, 0x63, 0x11, 0xa9, 0xe9, 0xba, 0xc1,
	0x90, 0xd3, 0x8f, 0x11, 0xb1, 0xf4, 0x5c, 0x1a, 0x6f, 0x91, 0xed, 0x4f, 0x60, 0xff, 0xde, 0x09,
	0xe1, 0x5b, 0x68, 0xeb, 0x2b, 0x7b, 0x2e, 0x2f, 0xf1, 0xd8, 0xfc, 0xdd, 0x07, 0x2f, 0xf3, 0xc9,
	0xe3, 0x07, 0x7c, 0x75, 0xb1, 0xbd, 0xda, 0x65, 0xcb, 0x3c, 0xeb, 0x57, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xa9, 0xd9, 0x6a, 0x5d, 0xe3, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ActionServiceClient is the client API for ActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActionServiceClient interface {
	InitJob(ctx context.Context, in *InitActionRequest, opts ...grpc.CallOption) (*InitActionResponse, error)
}

type actionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActionServiceClient(cc grpc.ClientConnInterface) ActionServiceClient {
	return &actionServiceClient{cc}
}

func (c *actionServiceClient) InitJob(ctx context.Context, in *InitActionRequest, opts ...grpc.CallOption) (*InitActionResponse, error) {
	out := new(InitActionResponse)
	err := c.cc.Invoke(ctx, "/rpc.ActionService/InitJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionServiceServer is the server API for ActionService service.
type ActionServiceServer interface {
	InitJob(context.Context, *InitActionRequest) (*InitActionResponse, error)
}

// UnimplementedActionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedActionServiceServer struct {
}

func (*UnimplementedActionServiceServer) InitJob(ctx context.Context, req *InitActionRequest) (*InitActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitJob not implemented")
}

func RegisterActionServiceServer(s *grpc.Server, srv ActionServiceServer) {
	s.RegisterService(&_ActionService_serviceDesc, srv)
}

func _ActionService_InitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).InitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ActionService/InitJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).InitJob(ctx, req.(*InitActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ActionService",
	HandlerType: (*ActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitJob",
			Handler:    _ActionService_InitJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
