// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculator_pb/calculator.proto

package calcpb

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

type Sum struct {
	Num1                 int64    `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2                 int64    `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sum) Reset()         { *m = Sum{} }
func (m *Sum) String() string { return proto.CompactTextString(m) }
func (*Sum) ProtoMessage()    {}
func (*Sum) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{0}
}

func (m *Sum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sum.Unmarshal(m, b)
}
func (m *Sum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sum.Marshal(b, m, deterministic)
}
func (m *Sum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sum.Merge(m, src)
}
func (m *Sum) XXX_Size() int {
	return xxx_messageInfo_Sum.Size(m)
}
func (m *Sum) XXX_DiscardUnknown() {
	xxx_messageInfo_Sum.DiscardUnknown(m)
}

var xxx_messageInfo_Sum proto.InternalMessageInfo

func (m *Sum) GetNum1() int64 {
	if m != nil {
		return m.Num1
	}
	return 0
}

func (m *Sum) GetNum2() int64 {
	if m != nil {
		return m.Num2
	}
	return 0
}

type SumRequest struct {
	Sum                  *Sum     `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{1}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetSum() *Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

type SumResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{2}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type DecomposeRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecomposeRequest) Reset()         { *m = DecomposeRequest{} }
func (m *DecomposeRequest) String() string { return proto.CompactTextString(m) }
func (*DecomposeRequest) ProtoMessage()    {}
func (*DecomposeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{3}
}

func (m *DecomposeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecomposeRequest.Unmarshal(m, b)
}
func (m *DecomposeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecomposeRequest.Marshal(b, m, deterministic)
}
func (m *DecomposeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecomposeRequest.Merge(m, src)
}
func (m *DecomposeRequest) XXX_Size() int {
	return xxx_messageInfo_DecomposeRequest.Size(m)
}
func (m *DecomposeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecomposeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecomposeRequest proto.InternalMessageInfo

func (m *DecomposeRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type DecomposeResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecomposeResponse) Reset()         { *m = DecomposeResponse{} }
func (m *DecomposeResponse) String() string { return proto.CompactTextString(m) }
func (*DecomposeResponse) ProtoMessage()    {}
func (*DecomposeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b44e1cc294e667d, []int{4}
}

func (m *DecomposeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecomposeResponse.Unmarshal(m, b)
}
func (m *DecomposeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecomposeResponse.Marshal(b, m, deterministic)
}
func (m *DecomposeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecomposeResponse.Merge(m, src)
}
func (m *DecomposeResponse) XXX_Size() int {
	return xxx_messageInfo_DecomposeResponse.Size(m)
}
func (m *DecomposeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecomposeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecomposeResponse proto.InternalMessageInfo

func (m *DecomposeResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Sum)(nil), "calculator.Sum")
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*DecomposeRequest)(nil), "calculator.DecomposeRequest")
	proto.RegisterType((*DecomposeResponse)(nil), "calculator.DecomposeResponse")
}

func init() {
	proto.RegisterFile("calculator/calculator_pb/calculator.proto", fileDescriptor_5b44e1cc294e667d)
}

var fileDescriptor_5b44e1cc294e667d = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x47, 0x30, 0xe3, 0x0b, 0x92, 0x90, 0x78, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x08, 0x11, 0x25, 0x5d, 0x2e, 0xe6, 0xe0, 0xd2, 0x5c, 0x21,
	0x21, 0x2e, 0x96, 0xbc, 0xd2, 0x5c, 0x43, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x30, 0x1b,
	0x2a, 0x66, 0x24, 0xc1, 0x04, 0x17, 0x33, 0x52, 0xd2, 0xe7, 0xe2, 0x0a, 0x2e, 0xcd, 0x0d, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe4, 0x62, 0x2e, 0x2e, 0xcd, 0x05, 0x6b, 0xe2, 0x36,
	0xe2, 0xd7, 0x43, 0xb2, 0x08, 0xa4, 0x08, 0x24, 0xa7, 0xa4, 0xca, 0xc5, 0x0d, 0xd6, 0x50, 0x5c,
	0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x02, 0xb5,
	0x09, 0xca, 0x53, 0xd2, 0xe2, 0x12, 0x70, 0x49, 0x4d, 0xce, 0xcf, 0x2d, 0xc8, 0x2f, 0x4e, 0x85,
	0x99, 0x2e, 0xc6, 0xc5, 0x96, 0x57, 0x9a, 0x9b, 0x94, 0x5a, 0x04, 0x53, 0x0b, 0xe1, 0x29, 0x69,
	0x73, 0x09, 0x22, 0xa9, 0xc5, 0x6f, 0xb0, 0xd1, 0x74, 0x46, 0x2e, 0x41, 0x67, 0xb8, 0xbb, 0x82,
	0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xcc, 0x20, 0xbe, 0x16, 0x43, 0x77, 0x32, 0xc4, 0x66,
	0x29, 0x71, 0x0c, 0x71, 0xa8, 0x2d, 0x5e, 0x5c, 0x9c, 0x70, 0xab, 0x85, 0x64, 0x90, 0x55, 0xa1,
	0xbb, 0x5e, 0x4a, 0x16, 0x87, 0x2c, 0xc4, 0x24, 0x03, 0x46, 0x27, 0x8e, 0x28, 0x36, 0x90, 0x8a,
	0x82, 0xa4, 0x24, 0x36, 0x70, 0xb4, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x4c, 0x68,
	0x41, 0xc3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Decompose(ctx context.Context, in *DecomposeRequest, opts ...grpc.CallOption) (CalculatorService_DecomposeClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Decompose(ctx context.Context, in *DecomposeRequest, opts ...grpc.CallOption) (CalculatorService_DecomposeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/Decompose", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceDecomposeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_DecomposeClient interface {
	Recv() (*DecomposeResponse, error)
	grpc.ClientStream
}

type calculatorServiceDecomposeClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceDecomposeClient) Recv() (*DecomposeResponse, error) {
	m := new(DecomposeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Decompose(*DecomposeRequest, CalculatorService_DecomposeServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) Decompose(req *DecomposeRequest, srv CalculatorService_DecomposeServer) error {
	return status.Errorf(codes.Unimplemented, "method Decompose not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Decompose_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DecomposeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).Decompose(m, &calculatorServiceDecomposeServer{stream})
}

type CalculatorService_DecomposeServer interface {
	Send(*DecomposeResponse) error
	grpc.ServerStream
}

type calculatorServiceDecomposeServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceDecomposeServer) Send(m *DecomposeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Decompose",
			Handler:       _CalculatorService_Decompose_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator/calculator_pb/calculator.proto",
}
