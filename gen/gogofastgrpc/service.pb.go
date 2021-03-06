// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

/*
	Package service is a generated protocol buffer package.

	It is generated from these files:
		service.proto

	It has these top-level messages:
		PutReq
		PutRes
		GetReq
		GetRes
		WriteReq
		WriteRes
		ReadReq
		ReadRes
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PutReq struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Blob []byte `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (m *PutReq) Reset()                    { *m = PutReq{} }
func (m *PutReq) String() string            { return proto.CompactTextString(m) }
func (*PutReq) ProtoMessage()               {}
func (*PutReq) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{0} }

func (m *PutReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PutReq) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

type PutRes struct {
}

func (m *PutRes) Reset()                    { *m = PutRes{} }
func (m *PutRes) String() string            { return proto.CompactTextString(m) }
func (*PutRes) ProtoMessage()               {}
func (*PutRes) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{1} }

type GetReq struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{2} }

func (m *GetReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetRes struct {
	Blob []byte `protobuf:"bytes,1,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (m *GetRes) Reset()                    { *m = GetRes{} }
func (m *GetRes) String() string            { return proto.CompactTextString(m) }
func (*GetRes) ProtoMessage()               {}
func (*GetRes) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{3} }

func (m *GetRes) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

type WriteReq struct {
	// Types that are valid to be assigned to Phase:
	//	*WriteReq_Name
	//	*WriteReq_Blob
	Phase isWriteReq_Phase `protobuf_oneof:"phase"`
}

func (m *WriteReq) Reset()                    { *m = WriteReq{} }
func (m *WriteReq) String() string            { return proto.CompactTextString(m) }
func (*WriteReq) ProtoMessage()               {}
func (*WriteReq) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{4} }

type isWriteReq_Phase interface {
	isWriteReq_Phase()
	MarshalTo([]byte) (int, error)
	Size() int
}

type WriteReq_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}
type WriteReq_Blob struct {
	Blob []byte `protobuf:"bytes,2,opt,name=blob,proto3,oneof"`
}

func (*WriteReq_Name) isWriteReq_Phase() {}
func (*WriteReq_Blob) isWriteReq_Phase() {}

func (m *WriteReq) GetPhase() isWriteReq_Phase {
	if m != nil {
		return m.Phase
	}
	return nil
}

func (m *WriteReq) GetName() string {
	if x, ok := m.GetPhase().(*WriteReq_Name); ok {
		return x.Name
	}
	return ""
}

func (m *WriteReq) GetBlob() []byte {
	if x, ok := m.GetPhase().(*WriteReq_Blob); ok {
		return x.Blob
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WriteReq) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WriteReq_OneofMarshaler, _WriteReq_OneofUnmarshaler, _WriteReq_OneofSizer, []interface{}{
		(*WriteReq_Name)(nil),
		(*WriteReq_Blob)(nil),
	}
}

func _WriteReq_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WriteReq)
	// phase
	switch x := m.Phase.(type) {
	case *WriteReq_Name:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Name)
	case *WriteReq_Blob:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.Blob)
	case nil:
	default:
		return fmt.Errorf("WriteReq.Phase has unexpected type %T", x)
	}
	return nil
}

func _WriteReq_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WriteReq)
	switch tag {
	case 1: // phase.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Phase = &WriteReq_Name{x}
		return true, err
	case 2: // phase.blob
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Phase = &WriteReq_Blob{x}
		return true, err
	default:
		return false, nil
	}
}

func _WriteReq_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WriteReq)
	// phase
	switch x := m.Phase.(type) {
	case *WriteReq_Name:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *WriteReq_Blob:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Blob)))
		n += len(x.Blob)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type WriteRes struct {
}

func (m *WriteRes) Reset()                    { *m = WriteRes{} }
func (m *WriteRes) String() string            { return proto.CompactTextString(m) }
func (*WriteRes) ProtoMessage()               {}
func (*WriteRes) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{5} }

type ReadReq struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BufSize uint32 `protobuf:"varint,2,opt,name=buf_size,json=bufSize,proto3" json:"buf_size,omitempty"`
}

func (m *ReadReq) Reset()                    { *m = ReadReq{} }
func (m *ReadReq) String() string            { return proto.CompactTextString(m) }
func (*ReadReq) ProtoMessage()               {}
func (*ReadReq) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{6} }

func (m *ReadReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadReq) GetBufSize() uint32 {
	if m != nil {
		return m.BufSize
	}
	return 0
}

type ReadRes struct {
	Blob []byte `protobuf:"bytes,1,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (m *ReadRes) Reset()                    { *m = ReadRes{} }
func (m *ReadRes) String() string            { return proto.CompactTextString(m) }
func (*ReadRes) ProtoMessage()               {}
func (*ReadRes) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{7} }

func (m *ReadRes) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

func init() {
	proto.RegisterType((*PutReq)(nil), "service.PutReq")
	proto.RegisterType((*PutRes)(nil), "service.PutRes")
	proto.RegisterType((*GetReq)(nil), "service.GetReq")
	proto.RegisterType((*GetRes)(nil), "service.GetRes")
	proto.RegisterType((*WriteReq)(nil), "service.WriteReq")
	proto.RegisterType((*WriteRes)(nil), "service.WriteRes")
	proto.RegisterType((*ReadReq)(nil), "service.ReadReq")
	proto.RegisterType((*ReadRes)(nil), "service.ReadRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Blober service

type BloberClient interface {
	Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*PutRes, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error)
	Write(ctx context.Context, opts ...grpc.CallOption) (Blober_WriteClient, error)
	Read(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (Blober_ReadClient, error)
}

type bloberClient struct {
	cc *grpc.ClientConn
}

func NewBloberClient(cc *grpc.ClientConn) BloberClient {
	return &bloberClient{cc}
}

func (c *bloberClient) Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*PutRes, error) {
	out := new(PutRes)
	err := grpc.Invoke(ctx, "/service.Blober/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloberClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRes, error) {
	out := new(GetRes)
	err := grpc.Invoke(ctx, "/service.Blober/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bloberClient) Write(ctx context.Context, opts ...grpc.CallOption) (Blober_WriteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Blober_serviceDesc.Streams[0], c.cc, "/service.Blober/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &bloberWriteClient{stream}
	return x, nil
}

type Blober_WriteClient interface {
	Send(*WriteReq) error
	CloseAndRecv() (*WriteRes, error)
	grpc.ClientStream
}

type bloberWriteClient struct {
	grpc.ClientStream
}

func (x *bloberWriteClient) Send(m *WriteReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bloberWriteClient) CloseAndRecv() (*WriteRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bloberClient) Read(ctx context.Context, in *ReadReq, opts ...grpc.CallOption) (Blober_ReadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Blober_serviceDesc.Streams[1], c.cc, "/service.Blober/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &bloberReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Blober_ReadClient interface {
	Recv() (*ReadRes, error)
	grpc.ClientStream
}

type bloberReadClient struct {
	grpc.ClientStream
}

func (x *bloberReadClient) Recv() (*ReadRes, error) {
	m := new(ReadRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Blober service

type BloberServer interface {
	Put(context.Context, *PutReq) (*PutRes, error)
	Get(context.Context, *GetReq) (*GetRes, error)
	Write(Blober_WriteServer) error
	Read(*ReadReq, Blober_ReadServer) error
}

func RegisterBloberServer(s *grpc.Server, srv BloberServer) {
	s.RegisterService(&_Blober_serviceDesc, srv)
}

func _Blober_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloberServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Blober/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloberServer).Put(ctx, req.(*PutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blober_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BloberServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Blober/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BloberServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blober_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BloberServer).Write(&bloberWriteServer{stream})
}

type Blober_WriteServer interface {
	SendAndClose(*WriteRes) error
	Recv() (*WriteReq, error)
	grpc.ServerStream
}

type bloberWriteServer struct {
	grpc.ServerStream
}

func (x *bloberWriteServer) SendAndClose(m *WriteRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bloberWriteServer) Recv() (*WriteReq, error) {
	m := new(WriteReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Blober_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BloberServer).Read(m, &bloberReadServer{stream})
}

type Blober_ReadServer interface {
	Send(*ReadRes) error
	grpc.ServerStream
}

type bloberReadServer struct {
	grpc.ServerStream
}

func (x *bloberReadServer) Send(m *ReadRes) error {
	return x.ServerStream.SendMsg(m)
}

var _Blober_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Blober",
	HandlerType: (*BloberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Blober_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Blober_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Write",
			Handler:       _Blober_Write_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Read",
			Handler:       _Blober_Read_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}

func (m *PutReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PutReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Blob) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Blob)))
		i += copy(dAtA[i:], m.Blob)
	}
	return i, nil
}

func (m *PutRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PutRes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *GetRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Blob) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Blob)))
		i += copy(dAtA[i:], m.Blob)
	}
	return i, nil
}

func (m *WriteReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Phase != nil {
		nn1, err := m.Phase.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *WriteReq_Name) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintService(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	return i, nil
}
func (m *WriteReq_Blob) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Blob != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Blob)))
		i += copy(dAtA[i:], m.Blob)
	}
	return i, nil
}
func (m *WriteRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteRes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ReadReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.BufSize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintService(dAtA, i, uint64(m.BufSize))
	}
	return i, nil
}

func (m *ReadRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadRes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Blob) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Blob)))
		i += copy(dAtA[i:], m.Blob)
	}
	return i, nil
}

func encodeFixed64Service(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Service(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PutReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Blob)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *PutRes) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *GetRes) Size() (n int) {
	var l int
	_ = l
	l = len(m.Blob)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *WriteReq) Size() (n int) {
	var l int
	_ = l
	if m.Phase != nil {
		n += m.Phase.Size()
	}
	return n
}

func (m *WriteReq_Name) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovService(uint64(l))
	return n
}
func (m *WriteReq_Blob) Size() (n int) {
	var l int
	_ = l
	if m.Blob != nil {
		l = len(m.Blob)
		n += 1 + l + sovService(uint64(l))
	}
	return n
}
func (m *WriteRes) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ReadReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.BufSize != 0 {
		n += 1 + sovService(uint64(m.BufSize))
	}
	return n
}

func (m *ReadRes) Size() (n int) {
	var l int
	_ = l
	l = len(m.Blob)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func sovService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PutReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PutReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blob = append(m.Blob[:0], dAtA[iNdEx:postIndex]...)
			if m.Blob == nil {
				m.Blob = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PutRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PutRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blob = append(m.Blob[:0], dAtA[iNdEx:postIndex]...)
			if m.Blob == nil {
				m.Blob = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WriteReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WriteReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phase = &WriteReq_Name{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Phase = &WriteReq_Blob{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WriteRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WriteRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufSize", wireType)
			}
			m.BufSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BufSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blob = append(m.Blob[:0], dAtA[iNdEx:postIndex]...)
			if m.Blob == nil {
				m.Blob = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("service.proto", fileDescriptorService) }

var fileDescriptorService = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x0c, 0xb8,
	0xd8, 0x02, 0x4a, 0x4b, 0x82, 0x52, 0x0b, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x52, 0x4e, 0x7e, 0x92, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0xad, 0xc4, 0x01, 0xd5, 0x51, 0xac, 0x24, 0xc3, 0xc5, 0xe6,
	0x9e, 0x8a, 0x4b, 0x2f, 0x5c, 0xb6, 0x18, 0x6e, 0x0a, 0x23, 0x92, 0x29, 0xf6, 0x5c, 0x1c, 0xe1,
	0x45, 0x99, 0x25, 0xa9, 0x20, 0xdd, 0x22, 0xc8, 0xba, 0x3d, 0x18, 0xa0, 0x76, 0x8b, 0x20, 0xdb,
	0x0d, 0x12, 0x05, 0xf1, 0x9c, 0xd8, 0xb9, 0x58, 0x0b, 0x32, 0x12, 0x8b, 0x53, 0x95, 0xb8, 0xe0,
	0x06, 0x14, 0x2b, 0x59, 0x70, 0xb1, 0x07, 0xa5, 0x26, 0xa6, 0xe0, 0xf2, 0x85, 0x24, 0x17, 0x47,
	0x52, 0x69, 0x5a, 0x7c, 0x71, 0x66, 0x55, 0x2a, 0xd8, 0x34, 0xde, 0x20, 0xf6, 0xa4, 0xd2, 0xb4,
	0xe0, 0xcc, 0xaa, 0x54, 0x25, 0x59, 0x98, 0x4e, 0xac, 0xae, 0x34, 0xda, 0xc9, 0xc8, 0xc5, 0xe6,
	0x94, 0x93, 0x9f, 0x94, 0x5a, 0x24, 0xa4, 0xce, 0xc5, 0x1c, 0x50, 0x5a, 0x22, 0xc4, 0xaf, 0x07,
	0x0b, 0x48, 0x48, 0xb0, 0x49, 0xa1, 0x09, 0x14, 0x83, 0x14, 0xba, 0xa7, 0x22, 0x2b, 0x84, 0x84,
	0x91, 0x14, 0x9a, 0x40, 0xb1, 0x90, 0x3e, 0x17, 0x2b, 0xd8, 0x07, 0x42, 0x82, 0x70, 0x19, 0x58,
	0x90, 0x48, 0x61, 0x08, 0x15, 0x6b, 0x30, 0x0a, 0xe9, 0x70, 0xb1, 0x80, 0x1c, 0x2b, 0x24, 0x00,
	0x97, 0x84, 0xfa, 0x5a, 0x0a, 0x5d, 0xa4, 0xd8, 0x80, 0xd1, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x1c,
	0xf7, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xb6, 0x38, 0x3e, 0x0c, 0x02, 0x00, 0x00,
}
