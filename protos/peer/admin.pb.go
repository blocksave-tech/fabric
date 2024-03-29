// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/admin.proto

/*
Package peer is a generated protocol buffer package.

It is generated from these files:
	peer/admin.proto
	peer/chaincode.proto
	peer/chaincode_event.proto
	peer/chaincode_shim.proto
	peer/configuration.proto
	peer/events.proto
	peer/peer.proto
	peer/proposal.proto
	peer/proposal_response.proto
	peer/query.proto
	peer/resources.proto
	peer/signed_cc_dep_spec.proto
	peer/transaction.proto

It has these top-level messages:
	ServerStatus
	LogLevelRequest
	LogLevelResponse
	AdminOperation
	ChaincodeID
	ChaincodeInput
	ChaincodeSpec
	ChaincodeDeploymentSpec
	ChaincodeInvocationSpec
	LifecycleEvent
	ChaincodeEvent
	ChaincodeMessage
	GetState
	PutState
	DelState
	GetStateByRange
	GetQueryResult
	GetHistoryForKey
	QueryStateNext
	QueryStateClose
	QueryResultBytes
	QueryResponse
	AnchorPeers
	AnchorPeer
	APIResource
	ACLs
	ChaincodeReg
	Interest
	Register
	Rejection
	Unregister
	FilteredBlock
	FilteredTransaction
	FilteredTransactionActions
	FilteredChaincodeAction
	SignedEvent
	Event
	DeliverResponse
	PeerID
	PeerEndpoint
	SignedProposal
	Proposal
	ChaincodeHeaderExtension
	ChaincodeProposalPayload
	ChaincodeAction
	ProposalResponse
	Response
	ProposalResponsePayload
	Endorsement
	ChaincodeQueryResponse
	ChaincodeInfo
	ChannelQueryResponse
	ChannelInfo
	ChaincodeIdentifier
	ChaincodeValidation
	VSCCArgs
	ChaincodeEndorsement
	ConfigTree
	SignedChaincodeDeploymentSpec
	SignedTransaction
	ProcessedTransaction
	Transaction
	TransactionAction
	ChaincodeActionPayload
	ChaincodeEndorsedAction
*/
package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import common "github.com/sinochem-tech/fabric/protos/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServerStatus_StatusCode int32

const (
	ServerStatus_UNDEFINED ServerStatus_StatusCode = 0
	ServerStatus_STARTED   ServerStatus_StatusCode = 1
	ServerStatus_STOPPED   ServerStatus_StatusCode = 2
	ServerStatus_PAUSED    ServerStatus_StatusCode = 3
	ServerStatus_ERROR     ServerStatus_StatusCode = 4
	ServerStatus_UNKNOWN   ServerStatus_StatusCode = 5
)

var ServerStatus_StatusCode_name = map[int32]string{
	0: "UNDEFINED",
	1: "STARTED",
	2: "STOPPED",
	3: "PAUSED",
	4: "ERROR",
	5: "UNKNOWN",
}
var ServerStatus_StatusCode_value = map[string]int32{
	"UNDEFINED": 0,
	"STARTED":   1,
	"STOPPED":   2,
	"PAUSED":    3,
	"ERROR":     4,
	"UNKNOWN":   5,
}

func (x ServerStatus_StatusCode) String() string {
	return proto.EnumName(ServerStatus_StatusCode_name, int32(x))
}
func (ServerStatus_StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ServerStatus struct {
	Status ServerStatus_StatusCode `protobuf:"varint,1,opt,name=status,enum=protos.ServerStatus_StatusCode" json:"status,omitempty"`
}

func (m *ServerStatus) Reset()                    { *m = ServerStatus{} }
func (m *ServerStatus) String() string            { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()               {}
func (*ServerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServerStatus) GetStatus() ServerStatus_StatusCode {
	if m != nil {
		return m.Status
	}
	return ServerStatus_UNDEFINED
}

type LogLevelRequest struct {
	LogModule string `protobuf:"bytes,1,opt,name=log_module,json=logModule" json:"log_module,omitempty"`
	LogLevel  string `protobuf:"bytes,2,opt,name=log_level,json=logLevel" json:"log_level,omitempty"`
}

func (m *LogLevelRequest) Reset()                    { *m = LogLevelRequest{} }
func (m *LogLevelRequest) String() string            { return proto.CompactTextString(m) }
func (*LogLevelRequest) ProtoMessage()               {}
func (*LogLevelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogLevelRequest) GetLogModule() string {
	if m != nil {
		return m.LogModule
	}
	return ""
}

func (m *LogLevelRequest) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

type LogLevelResponse struct {
	LogModule string `protobuf:"bytes,1,opt,name=log_module,json=logModule" json:"log_module,omitempty"`
	LogLevel  string `protobuf:"bytes,2,opt,name=log_level,json=logLevel" json:"log_level,omitempty"`
}

func (m *LogLevelResponse) Reset()                    { *m = LogLevelResponse{} }
func (m *LogLevelResponse) String() string            { return proto.CompactTextString(m) }
func (*LogLevelResponse) ProtoMessage()               {}
func (*LogLevelResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogLevelResponse) GetLogModule() string {
	if m != nil {
		return m.LogModule
	}
	return ""
}

func (m *LogLevelResponse) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

type AdminOperation struct {
	// Types that are valid to be assigned to Content:
	//	*AdminOperation_LogReq
	Content isAdminOperation_Content `protobuf_oneof:"content"`
}

func (m *AdminOperation) Reset()                    { *m = AdminOperation{} }
func (m *AdminOperation) String() string            { return proto.CompactTextString(m) }
func (*AdminOperation) ProtoMessage()               {}
func (*AdminOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isAdminOperation_Content interface{ isAdminOperation_Content() }

type AdminOperation_LogReq struct {
	LogReq *LogLevelRequest `protobuf:"bytes,1,opt,name=logReq,oneof"`
}

func (*AdminOperation_LogReq) isAdminOperation_Content() {}

func (m *AdminOperation) GetContent() isAdminOperation_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *AdminOperation) GetLogReq() *LogLevelRequest {
	if x, ok := m.GetContent().(*AdminOperation_LogReq); ok {
		return x.LogReq
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AdminOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AdminOperation_OneofMarshaler, _AdminOperation_OneofUnmarshaler, _AdminOperation_OneofSizer, []interface{}{
		(*AdminOperation_LogReq)(nil),
	}
}

func _AdminOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AdminOperation)
	// content
	switch x := m.Content.(type) {
	case *AdminOperation_LogReq:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogReq); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AdminOperation.Content has unexpected type %T", x)
	}
	return nil
}

func _AdminOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AdminOperation)
	switch tag {
	case 1: // content.logReq
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogLevelRequest)
		err := b.DecodeMessage(msg)
		m.Content = &AdminOperation_LogReq{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AdminOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AdminOperation)
	// content
	switch x := m.Content.(type) {
	case *AdminOperation_LogReq:
		s := proto.Size(x.LogReq)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ServerStatus)(nil), "protos.ServerStatus")
	proto.RegisterType((*LogLevelRequest)(nil), "protos.LogLevelRequest")
	proto.RegisterType((*LogLevelResponse)(nil), "protos.LogLevelResponse")
	proto.RegisterType((*AdminOperation)(nil), "protos.AdminOperation")
	proto.RegisterEnum("protos.ServerStatus_StatusCode", ServerStatus_StatusCode_name, ServerStatus_StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Admin service

type AdminClient interface {
	GetStatus(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*ServerStatus, error)
	StartServer(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*ServerStatus, error)
	GetModuleLogLevel(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogLevelResponse, error)
	SetModuleLogLevel(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogLevelResponse, error)
	RevertLogLevels(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) GetStatus(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*ServerStatus, error) {
	out := new(ServerStatus)
	err := grpc.Invoke(ctx, "/protos.Admin/GetStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) StartServer(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*ServerStatus, error) {
	out := new(ServerStatus)
	err := grpc.Invoke(ctx, "/protos.Admin/StartServer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetModuleLogLevel(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogLevelResponse, error) {
	out := new(LogLevelResponse)
	err := grpc.Invoke(ctx, "/protos.Admin/GetModuleLogLevel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SetModuleLogLevel(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*LogLevelResponse, error) {
	out := new(LogLevelResponse)
	err := grpc.Invoke(ctx, "/protos.Admin/SetModuleLogLevel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) RevertLogLevels(ctx context.Context, in *common.Envelope, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/protos.Admin/RevertLogLevels", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminServer interface {
	GetStatus(context.Context, *common.Envelope) (*ServerStatus, error)
	StartServer(context.Context, *common.Envelope) (*ServerStatus, error)
	GetModuleLogLevel(context.Context, *common.Envelope) (*LogLevelResponse, error)
	SetModuleLogLevel(context.Context, *common.Envelope) (*LogLevelResponse, error)
	RevertLogLevels(context.Context, *common.Envelope) (*google_protobuf.Empty, error)
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetStatus(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_StartServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).StartServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/StartServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).StartServer(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetModuleLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetModuleLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/GetModuleLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetModuleLogLevel(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_SetModuleLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).SetModuleLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/SetModuleLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).SetModuleLogLevel(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_RevertLogLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).RevertLogLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Admin/RevertLogLevels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).RevertLogLevels(ctx, req.(*common.Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _Admin_GetStatus_Handler,
		},
		{
			MethodName: "StartServer",
			Handler:    _Admin_StartServer_Handler,
		},
		{
			MethodName: "GetModuleLogLevel",
			Handler:    _Admin_GetModuleLogLevel_Handler,
		},
		{
			MethodName: "SetModuleLogLevel",
			Handler:    _Admin_SetModuleLogLevel_Handler,
		},
		{
			MethodName: "RevertLogLevels",
			Handler:    _Admin_RevertLogLevels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer/admin.proto",
}

func init() { proto.RegisterFile("peer/admin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x6e, 0x0b, 0xed, 0xc8, 0xe9, 0xd8, 0x82, 0x41, 0x50, 0x75, 0x42, 0xa0, 0x5c, 0xc1, 0x8d,
	0x23, 0x8a, 0xd0, 0xae, 0xb8, 0x68, 0x49, 0x18, 0x88, 0x2d, 0xad, 0x9c, 0x55, 0x08, 0x24, 0x34,
	0xa5, 0xed, 0x99, 0x57, 0xe1, 0xc4, 0x99, 0xe3, 0x54, 0xda, 0xeb, 0xf0, 0x1c, 0x3c, 0x1c, 0x8a,
	0x9d, 0x68, 0x13, 0xf4, 0x06, 0xf5, 0xea, 0xf8, 0x1c, 0x7f, 0xdf, 0x67, 0x9f, 0x3f, 0x70, 0x73,
	0x44, 0xe5, 0x27, 0xab, 0x74, 0x9d, 0xd1, 0x5c, 0x49, 0x2d, 0x49, 0xcf, 0x98, 0x62, 0x78, 0xc4,
	0xa5, 0xe4, 0x02, 0x7d, 0xe3, 0x2e, 0xca, 0x4b, 0x1f, 0xd3, 0x5c, 0xdf, 0x58, 0xd0, 0xf0, 0xf1,
	0x52, 0xa6, 0xa9, 0xcc, 0x7c, 0x6b, 0x6c, 0xd0, 0xfb, 0xd5, 0x86, 0xfd, 0x18, 0xd5, 0x06, 0x55,
	0xac, 0x13, 0x5d, 0x16, 0xe4, 0x18, 0x7a, 0x85, 0x39, 0x0d, 0xda, 0x2f, 0xdb, 0xaf, 0x0e, 0x46,
	0x2f, 0x2c, 0xb0, 0xa0, 0x77, 0x51, 0xd4, 0x9a, 0x0f, 0x72, 0x85, 0xac, 0x86, 0x7b, 0xdf, 0x00,
	0x6e, 0xa3, 0xe4, 0x21, 0x38, 0xf3, 0x28, 0x08, 0x3f, 0x7e, 0x8e, 0xc2, 0xc0, 0x6d, 0x91, 0x3e,
	0xec, 0xc5, 0xe7, 0x63, 0x76, 0x1e, 0x06, 0x6e, 0xdb, 0x3a, 0xd3, 0xd9, 0x2c, 0x0c, 0xdc, 0x0e,
	0x01, 0xe8, 0xcd, 0xc6, 0xf3, 0x38, 0x0c, 0xdc, 0x7b, 0xc4, 0x81, 0x6e, 0xc8, 0xd8, 0x94, 0xb9,
	0xf7, 0x2b, 0xcc, 0x3c, 0xfa, 0x12, 0x4d, 0xbf, 0x46, 0x6e, 0xd7, 0x3b, 0x83, 0xc3, 0x53, 0xc9,
	0x4f, 0x71, 0x83, 0x82, 0xe1, 0x75, 0x89, 0x85, 0x26, 0xcf, 0x01, 0x84, 0xe4, 0x17, 0xa9, 0x5c,
	0x95, 0x02, 0xcd, 0x57, 0x1d, 0xe6, 0x08, 0xc9, 0xcf, 0x4c, 0x80, 0x1c, 0x41, 0xe5, 0x5c, 0x88,
	0x8a, 0x32, 0xe8, 0x98, 0xdb, 0x07, 0xa2, 0x96, 0xf0, 0x22, 0x70, 0x6f, 0xe5, 0x8a, 0x5c, 0x66,
	0x05, 0xee, 0xa8, 0x77, 0x30, 0xae, 0x9a, 0x31, 0xcd, 0x51, 0x25, 0x7a, 0x2d, 0x33, 0xf2, 0x06,
	0x7a, 0x42, 0x72, 0x86, 0xd7, 0x46, 0xa9, 0x3f, 0x7a, 0xd6, 0x14, 0xf1, 0xaf, 0x34, 0x3e, 0xb5,
	0x58, 0x0d, 0x9c, 0x38, 0xb0, 0xb7, 0x94, 0x99, 0xc6, 0x4c, 0x8f, 0x7e, 0x77, 0xa0, 0x6b, 0x04,
	0xc9, 0x3b, 0x70, 0x4e, 0x50, 0xd7, 0x9d, 0x71, 0x69, 0xdd, 0xb9, 0x30, 0xdb, 0xa0, 0x90, 0x39,
	0x0e, 0x9f, 0x6c, 0xeb, 0x8d, 0xd7, 0x22, 0xc7, 0xd0, 0x8f, 0x75, 0xa2, 0xb4, 0x0d, 0xff, 0x07,
	0x71, 0x0c, 0x8f, 0x4e, 0x50, 0xdb, 0x9c, 0x9b, 0xaf, 0x6e, 0xa1, 0x0f, 0xfe, 0x4d, 0xc7, 0x96,
	0xd1, 0x4a, 0xc4, 0x3b, 0x4a, 0xbc, 0x87, 0x43, 0x86, 0x1b, 0x54, 0xba, 0xb9, 0xdb, 0x96, 0xfb,
	0x53, 0x6a, 0x67, 0x9d, 0x36, 0xb3, 0x4e, 0xc3, 0x6a, 0xd6, 0xbd, 0xd6, 0xe4, 0x07, 0x78, 0x52,
	0x71, 0x7a, 0x75, 0x93, 0xa3, 0x12, 0xb8, 0xe2, 0xa8, 0xe8, 0x65, 0xb2, 0x50, 0xeb, 0x65, 0xf3,
	0x64, 0xb5, 0x3e, 0x93, 0x7d, 0x53, 0xe1, 0x59, 0xb2, 0xfc, 0x99, 0x70, 0xfc, 0xfe, 0x9a, 0xaf,
	0xf5, 0x55, 0xb9, 0xa8, 0x5e, 0xf1, 0xef, 0x10, 0x7d, 0x4b, 0xb4, 0xfb, 0x54, 0xf8, 0x15, 0x71,
	0x61, 0x77, 0xed, 0xed, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x20, 0x26, 0x0b, 0x86, 0x03,
	0x00, 0x00,
}
