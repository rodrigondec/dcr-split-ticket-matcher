// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package dcrticketmatcher is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	FindMatchesRequest
	FindMatchesResponse
	GenerateTicketRequest
	GenerateTicketResponse
	PublishTicketRequest
	PublishTicketResponse
*/
package dcrticketmatcher

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type FindMatchesRequest struct {
	Amount uint64 `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
}

func (m *FindMatchesRequest) Reset()                    { *m = FindMatchesRequest{} }
func (m *FindMatchesRequest) String() string            { return proto.CompactTextString(m) }
func (*FindMatchesRequest) ProtoMessage()               {}
func (*FindMatchesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FindMatchesRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type FindMatchesResponse struct {
	SessionId uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	Amount    uint64 `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Fee       uint64 `protobuf:"varint,3,opt,name=fee" json:"fee,omitempty"`
}

func (m *FindMatchesResponse) Reset()                    { *m = FindMatchesResponse{} }
func (m *FindMatchesResponse) String() string            { return proto.CompactTextString(m) }
func (*FindMatchesResponse) ProtoMessage()               {}
func (*FindMatchesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FindMatchesResponse) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *FindMatchesResponse) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *FindMatchesResponse) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type GenerateTicketRequest struct {
	SessionId        uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	CommitmentOutput []byte `protobuf:"bytes,2,opt,name=commitment_output,json=commitmentOutput,proto3" json:"commitment_output,omitempty"`
	ChangeOutput     []byte `protobuf:"bytes,3,opt,name=change_output,json=changeOutput,proto3" json:"change_output,omitempty"`
}

func (m *GenerateTicketRequest) Reset()                    { *m = GenerateTicketRequest{} }
func (m *GenerateTicketRequest) String() string            { return proto.CompactTextString(m) }
func (*GenerateTicketRequest) ProtoMessage()               {}
func (*GenerateTicketRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GenerateTicketRequest) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *GenerateTicketRequest) GetCommitmentOutput() []byte {
	if m != nil {
		return m.CommitmentOutput
	}
	return nil
}

func (m *GenerateTicketRequest) GetChangeOutput() []byte {
	if m != nil {
		return m.ChangeOutput
	}
	return nil
}

type GenerateTicketResponse struct {
	Transaction []byte `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	OutputIndex int32  `protobuf:"varint,2,opt,name=output_index,json=outputIndex" json:"output_index,omitempty"`
}

func (m *GenerateTicketResponse) Reset()                    { *m = GenerateTicketResponse{} }
func (m *GenerateTicketResponse) String() string            { return proto.CompactTextString(m) }
func (*GenerateTicketResponse) ProtoMessage()               {}
func (*GenerateTicketResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GenerateTicketResponse) GetTransaction() []byte {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *GenerateTicketResponse) GetOutputIndex() int32 {
	if m != nil {
		return m.OutputIndex
	}
	return 0
}

type PublishTicketRequest struct {
	SessionId  uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	SplitTx    []byte `protobuf:"bytes,2,opt,name=split_tx,json=splitTx,proto3" json:"split_tx,omitempty"`
	TicketTxin []byte `protobuf:"bytes,3,opt,name=ticket_txin,json=ticketTxin,proto3" json:"ticket_txin,omitempty"`
}

func (m *PublishTicketRequest) Reset()                    { *m = PublishTicketRequest{} }
func (m *PublishTicketRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishTicketRequest) ProtoMessage()               {}
func (*PublishTicketRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PublishTicketRequest) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *PublishTicketRequest) GetSplitTx() []byte {
	if m != nil {
		return m.SplitTx
	}
	return nil
}

func (m *PublishTicketRequest) GetTicketTxin() []byte {
	if m != nil {
		return m.TicketTxin
	}
	return nil
}

type PublishTicketResponse struct {
	TicketTx []byte `protobuf:"bytes,1,opt,name=ticket_tx,json=ticketTx,proto3" json:"ticket_tx,omitempty"`
}

func (m *PublishTicketResponse) Reset()                    { *m = PublishTicketResponse{} }
func (m *PublishTicketResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishTicketResponse) ProtoMessage()               {}
func (*PublishTicketResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PublishTicketResponse) GetTicketTx() []byte {
	if m != nil {
		return m.TicketTx
	}
	return nil
}

func init() {
	proto.RegisterType((*FindMatchesRequest)(nil), "dcrticketmatcher.FindMatchesRequest")
	proto.RegisterType((*FindMatchesResponse)(nil), "dcrticketmatcher.FindMatchesResponse")
	proto.RegisterType((*GenerateTicketRequest)(nil), "dcrticketmatcher.GenerateTicketRequest")
	proto.RegisterType((*GenerateTicketResponse)(nil), "dcrticketmatcher.GenerateTicketResponse")
	proto.RegisterType((*PublishTicketRequest)(nil), "dcrticketmatcher.PublishTicketRequest")
	proto.RegisterType((*PublishTicketResponse)(nil), "dcrticketmatcher.PublishTicketResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SplitTicketMatcherService service

type SplitTicketMatcherServiceClient interface {
	FindMatches(ctx context.Context, in *FindMatchesRequest, opts ...grpc.CallOption) (*FindMatchesResponse, error)
	GenerateTicket(ctx context.Context, in *GenerateTicketRequest, opts ...grpc.CallOption) (*GenerateTicketResponse, error)
	PublishTicket(ctx context.Context, in *PublishTicketRequest, opts ...grpc.CallOption) (*PublishTicketResponse, error)
}

type splitTicketMatcherServiceClient struct {
	cc *grpc.ClientConn
}

func NewSplitTicketMatcherServiceClient(cc *grpc.ClientConn) SplitTicketMatcherServiceClient {
	return &splitTicketMatcherServiceClient{cc}
}

func (c *splitTicketMatcherServiceClient) FindMatches(ctx context.Context, in *FindMatchesRequest, opts ...grpc.CallOption) (*FindMatchesResponse, error) {
	out := new(FindMatchesResponse)
	err := grpc.Invoke(ctx, "/dcrticketmatcher.SplitTicketMatcherService/FindMatches", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *splitTicketMatcherServiceClient) GenerateTicket(ctx context.Context, in *GenerateTicketRequest, opts ...grpc.CallOption) (*GenerateTicketResponse, error) {
	out := new(GenerateTicketResponse)
	err := grpc.Invoke(ctx, "/dcrticketmatcher.SplitTicketMatcherService/GenerateTicket", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *splitTicketMatcherServiceClient) PublishTicket(ctx context.Context, in *PublishTicketRequest, opts ...grpc.CallOption) (*PublishTicketResponse, error) {
	out := new(PublishTicketResponse)
	err := grpc.Invoke(ctx, "/dcrticketmatcher.SplitTicketMatcherService/PublishTicket", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SplitTicketMatcherService service

type SplitTicketMatcherServiceServer interface {
	FindMatches(context.Context, *FindMatchesRequest) (*FindMatchesResponse, error)
	GenerateTicket(context.Context, *GenerateTicketRequest) (*GenerateTicketResponse, error)
	PublishTicket(context.Context, *PublishTicketRequest) (*PublishTicketResponse, error)
}

func RegisterSplitTicketMatcherServiceServer(s *grpc.Server, srv SplitTicketMatcherServiceServer) {
	s.RegisterService(&_SplitTicketMatcherService_serviceDesc, srv)
}

func _SplitTicketMatcherService_FindMatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMatchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SplitTicketMatcherServiceServer).FindMatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcrticketmatcher.SplitTicketMatcherService/FindMatches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SplitTicketMatcherServiceServer).FindMatches(ctx, req.(*FindMatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SplitTicketMatcherService_GenerateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SplitTicketMatcherServiceServer).GenerateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcrticketmatcher.SplitTicketMatcherService/GenerateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SplitTicketMatcherServiceServer).GenerateTicket(ctx, req.(*GenerateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SplitTicketMatcherService_PublishTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SplitTicketMatcherServiceServer).PublishTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcrticketmatcher.SplitTicketMatcherService/PublishTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SplitTicketMatcherServiceServer).PublishTicket(ctx, req.(*PublishTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SplitTicketMatcherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dcrticketmatcher.SplitTicketMatcherService",
	HandlerType: (*SplitTicketMatcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindMatches",
			Handler:    _SplitTicketMatcherService_FindMatches_Handler,
		},
		{
			MethodName: "GenerateTicket",
			Handler:    _SplitTicketMatcherService_GenerateTicket_Handler,
		},
		{
			MethodName: "PublishTicket",
			Handler:    _SplitTicketMatcherService_PublishTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdb, 0x6e, 0xda, 0x40,
	0x10, 0x15, 0xd0, 0x52, 0x18, 0x9b, 0x8a, 0x6e, 0x0b, 0x02, 0xaa, 0xaa, 0xd4, 0xbd, 0x80, 0xd4,
	0x8a, 0x87, 0xb6, 0xdf, 0xd0, 0x8a, 0x87, 0x28, 0x91, 0xe1, 0x29, 0x52, 0xe2, 0x98, 0xf5, 0x24,
	0xac, 0x82, 0x77, 0x8d, 0x77, 0x1d, 0xf9, 0x13, 0xf2, 0x55, 0xf9, 0xb6, 0x88, 0xdd, 0x25, 0x31,
	0x17, 0xc9, 0xca, 0x9b, 0xe7, 0xf8, 0xcc, 0x9c, 0x73, 0x3c, 0x63, 0x68, 0x86, 0x09, 0x9b, 0x24,
	0xa9, 0x50, 0x82, 0xb4, 0x23, 0x9a, 0x2a, 0x46, 0x6f, 0x51, 0xc5, 0xa1, 0xa2, 0x4b, 0x4c, 0xbd,
	0x5f, 0x40, 0xfe, 0x31, 0x1e, 0x9d, 0xe8, 0x52, 0xfa, 0xb8, 0xce, 0x50, 0x2a, 0xd2, 0x85, 0x7a,
	0x18, 0x8b, 0x8c, 0xab, 0x5e, 0x65, 0x58, 0x19, 0xbf, 0xf2, 0x6d, 0xe5, 0x5d, 0xc2, 0xfb, 0x1d,
	0xb6, 0x4c, 0x04, 0x97, 0x48, 0x3e, 0x01, 0x48, 0x94, 0x92, 0x09, 0x1e, 0xb0, 0xc8, 0xb6, 0x34,
	0x2d, 0x32, 0x8d, 0x0a, 0xd3, 0xaa, 0xc5, 0x69, 0xa4, 0x0d, 0xb5, 0x6b, 0xc4, 0x5e, 0x4d, 0x83,
	0x9b, 0x47, 0xef, 0xbe, 0x02, 0x9d, 0xff, 0xc8, 0x31, 0x0d, 0x15, 0xce, 0xb5, 0xcf, 0xad, 0xa3,
	0x12, 0x89, 0x9f, 0xf0, 0x8e, 0x8a, 0x38, 0x66, 0x2a, 0x46, 0xae, 0x02, 0x91, 0xa9, 0x24, 0x33,
	0x6a, 0xae, 0xdf, 0x7e, 0x7e, 0x71, 0xaa, 0x71, 0xf2, 0x15, 0x5a, 0x74, 0x19, 0xf2, 0x1b, 0xdc,
	0x12, 0x6b, 0x9a, 0xe8, 0x1a, 0xd0, 0x90, 0xbc, 0x0b, 0xe8, 0xee, 0x3b, 0xb1, 0x69, 0x87, 0xe0,
	0xa8, 0x34, 0xe4, 0x32, 0xa4, 0x8a, 0x09, 0xae, 0xbd, 0xb8, 0x7e, 0x11, 0x22, 0x5f, 0xc0, 0x35,
	0x93, 0x03, 0xc6, 0x23, 0xcc, 0xb5, 0x91, 0xd7, 0xbe, 0x63, 0xb0, 0xe9, 0x06, 0xf2, 0xd6, 0xf0,
	0xe1, 0x2c, 0x5b, 0xac, 0x98, 0x5c, 0xbe, 0x28, 0x67, 0x1f, 0x1a, 0x32, 0x59, 0x31, 0x15, 0xa8,
	0xdc, 0xc6, 0x7b, 0xa3, 0xeb, 0x79, 0x4e, 0x3e, 0x83, 0x63, 0x56, 0x1b, 0xa8, 0x9c, 0x71, 0x9b,
	0x09, 0x0c, 0x34, 0xcf, 0x19, 0xf7, 0xfe, 0x42, 0x67, 0x4f, 0xd2, 0x06, 0xfa, 0x08, 0xcd, 0xa7,
	0x4e, 0x1b, 0xa7, 0xb1, 0xed, 0xfb, 0xfd, 0x50, 0x85, 0xfe, 0x4c, 0x4b, 0x68, 0xc4, 0xac, 0x3e,
	0x9d, 0x61, 0x7a, 0xc7, 0x28, 0x92, 0x73, 0x70, 0x0a, 0x07, 0x41, 0xbe, 0x4d, 0xf6, 0x0f, 0x6c,
	0x72, 0x78, 0x5d, 0x83, 0xef, 0x25, 0x2c, 0x6b, 0x8b, 0xc2, 0xdb, 0xdd, 0x0d, 0x90, 0xd1, 0x61,
	0xe3, 0xd1, 0x6b, 0x19, 0x8c, 0xcb, 0x89, 0x56, 0xe4, 0x0a, 0x5a, 0x3b, 0x1f, 0x85, 0xfc, 0x38,
	0x6c, 0x3d, 0xb6, 0xa8, 0xc1, 0xa8, 0x94, 0x67, 0x14, 0x16, 0x75, 0xfd, 0xeb, 0xfd, 0x79, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0x04, 0x2b, 0x86, 0x41, 0x87, 0x03, 0x00, 0x00,
}
