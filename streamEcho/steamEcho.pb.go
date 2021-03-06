// Code generated by protoc-gen-go. DO NOT EDIT.
// source: steamEcho.proto

/*
Package streamEchoService is a generated protocol buffer package.

It is generated from these files:
	steamEcho.proto

It has these top-level messages:
	EchoRequest
	EchoReply
*/
package streamEchoService

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

// the request message containing the echoMessage
type EchoRequest struct {
	RequestMessage string `protobuf:"bytes,1,opt,name=requestMessage" json:"requestMessage,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoRequest) GetRequestMessage() string {
	if m != nil {
		return m.RequestMessage
	}
	return ""
}

// the response message
type EchoReply struct {
	EplayMessage string `protobuf:"bytes,1,opt,name=eplayMessage" json:"eplayMessage,omitempty"`
	TimeNow      string `protobuf:"bytes,2,opt,name=timeNow" json:"timeNow,omitempty"`
}

func (m *EchoReply) Reset()                    { *m = EchoReply{} }
func (m *EchoReply) String() string            { return proto.CompactTextString(m) }
func (*EchoReply) ProtoMessage()               {}
func (*EchoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EchoReply) GetEplayMessage() string {
	if m != nil {
		return m.EplayMessage
	}
	return ""
}

func (m *EchoReply) GetTimeNow() string {
	if m != nil {
		return m.TimeNow
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "streamEchoService.EchoRequest")
	proto.RegisterType((*EchoReply)(nil), "streamEchoService.EchoReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echoer service

type EchoerClient interface {
	Echo(ctx context.Context, opts ...grpc.CallOption) (Echoer_EchoClient, error)
}

type echoerClient struct {
	cc *grpc.ClientConn
}

func NewEchoerClient(cc *grpc.ClientConn) EchoerClient {
	return &echoerClient{cc}
}

func (c *echoerClient) Echo(ctx context.Context, opts ...grpc.CallOption) (Echoer_EchoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Echoer_serviceDesc.Streams[0], c.cc, "/streamEchoService.Echoer/Echo", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoerEchoClient{stream}
	return x, nil
}

type Echoer_EchoClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoReply, error)
	grpc.ClientStream
}

type echoerEchoClient struct {
	grpc.ClientStream
}

func (x *echoerEchoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoerEchoClient) Recv() (*EchoReply, error) {
	m := new(EchoReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Echoer service

type EchoerServer interface {
	Echo(Echoer_EchoServer) error
}

func RegisterEchoerServer(s *grpc.Server, srv EchoerServer) {
	s.RegisterService(&_Echoer_serviceDesc, srv)
}

func _Echoer_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoerServer).Echo(&echoerEchoServer{stream})
}

type Echoer_EchoServer interface {
	Send(*EchoReply) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoerEchoServer struct {
	grpc.ServerStream
}

func (x *echoerEchoServer) Send(m *EchoReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoerEchoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Echoer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "streamEchoService.Echoer",
	HandlerType: (*EchoerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Echo",
			Handler:       _Echoer_Echo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "steamEcho.proto",
}

func init() { proto.RegisterFile("steamEcho.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2e, 0x49, 0x4d,
	0xcc, 0x75, 0x4d, 0xce, 0xc8, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2c, 0x2e, 0x29,
	0x82, 0x8a, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x2a, 0x99, 0x72, 0x71, 0x83, 0xb8, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x6a, 0x5c, 0x7c, 0x45, 0x10, 0xa6, 0x6f, 0x6a, 0x71,
	0x71, 0x62, 0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9a, 0xa8, 0x92, 0x27, 0x17,
	0x27, 0x44, 0x5b, 0x41, 0x4e, 0xa5, 0x90, 0x12, 0x17, 0x4f, 0x6a, 0x41, 0x4e, 0x62, 0x25, 0xaa,
	0x16, 0x14, 0x31, 0x21, 0x09, 0x2e, 0xf6, 0x92, 0xcc, 0xdc, 0x54, 0xbf, 0xfc, 0x72, 0x09, 0x26,
	0xb0, 0x34, 0x8c, 0x6b, 0x14, 0xc2, 0xc5, 0x06, 0x32, 0x2a, 0xb5, 0x48, 0xc8, 0x8b, 0x8b, 0x05,
	0xc4, 0x12, 0x92, 0xd3, 0xc3, 0x70, 0xa7, 0x1e, 0x92, 0x23, 0xa5, 0x64, 0x70, 0xca, 0x17, 0xe4,
	0x54, 0x2a, 0x31, 0x68, 0x30, 0x1a, 0x30, 0x3a, 0xa9, 0x71, 0x61, 0x7a, 0xd6, 0x89, 0x3f, 0x18,
	0x2e, 0x14, 0x00, 0x0a, 0x90, 0x00, 0xc6, 0x24, 0x36, 0x70, 0xc8, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xf5, 0xe9, 0x44, 0xfe, 0x2c, 0x01, 0x00, 0x00,
}
