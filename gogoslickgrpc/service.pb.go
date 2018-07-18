// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	PutReq
	PutRes
	StreamReq
	StreamRes
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

type StreamReq struct {
	// Types that are valid to be assigned to Phase:
	//	*StreamReq_Name
	//	*StreamReq_Blob
	Phase isStreamReq_Phase `protobuf_oneof:"phase"`
}

func (m *StreamReq) Reset()                    { *m = StreamReq{} }
func (m *StreamReq) String() string            { return proto.CompactTextString(m) }
func (*StreamReq) ProtoMessage()               {}
func (*StreamReq) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{2} }

type isStreamReq_Phase interface {
	isStreamReq_Phase()
	MarshalTo([]byte) (int, error)
	Size() int
}

type StreamReq_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}
type StreamReq_Blob struct {
	Blob []byte `protobuf:"bytes,2,opt,name=blob,proto3,oneof"`
}

func (*StreamReq_Name) isStreamReq_Phase() {}
func (*StreamReq_Blob) isStreamReq_Phase() {}

func (m *StreamReq) GetPhase() isStreamReq_Phase {
	if m != nil {
		return m.Phase
	}
	return nil
}

func (m *StreamReq) GetName() string {
	if x, ok := m.GetPhase().(*StreamReq_Name); ok {
		return x.Name
	}
	return ""
}

func (m *StreamReq) GetBlob() []byte {
	if x, ok := m.GetPhase().(*StreamReq_Blob); ok {
		return x.Blob
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamReq) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamReq_OneofMarshaler, _StreamReq_OneofUnmarshaler, _StreamReq_OneofSizer, []interface{}{
		(*StreamReq_Name)(nil),
		(*StreamReq_Blob)(nil),
	}
}

func _StreamReq_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamReq)
	// phase
	switch x := m.Phase.(type) {
	case *StreamReq_Name:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Name)
	case *StreamReq_Blob:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.Blob)
	case nil:
	default:
		return fmt.Errorf("StreamReq.Phase has unexpected type %T", x)
	}
	return nil
}

func _StreamReq_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamReq)
	switch tag {
	case 1: // phase.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Phase = &StreamReq_Name{x}
		return true, err
	case 2: // phase.blob
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Phase = &StreamReq_Blob{x}
		return true, err
	default:
		return false, nil
	}
}

func _StreamReq_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamReq)
	// phase
	switch x := m.Phase.(type) {
	case *StreamReq_Name:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *StreamReq_Blob:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Blob)))
		n += len(x.Blob)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StreamRes struct {
}

func (m *StreamRes) Reset()                    { *m = StreamRes{} }
func (m *StreamRes) String() string            { return proto.CompactTextString(m) }
func (*StreamRes) ProtoMessage()               {}
func (*StreamRes) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{3} }

func init() {
	proto.RegisterType((*PutReq)(nil), "service.PutReq")
	proto.RegisterType((*PutRes)(nil), "service.PutRes")
	proto.RegisterType((*StreamReq)(nil), "service.StreamReq")
	proto.RegisterType((*StreamRes)(nil), "service.StreamRes")
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
	Stream(ctx context.Context, opts ...grpc.CallOption) (Blober_StreamClient, error)
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

func (c *bloberClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Blober_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Blober_serviceDesc.Streams[0], c.cc, "/service.Blober/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &bloberStreamClient{stream}
	return x, nil
}

type Blober_StreamClient interface {
	Send(*StreamReq) error
	CloseAndRecv() (*StreamRes, error)
	grpc.ClientStream
}

type bloberStreamClient struct {
	grpc.ClientStream
}

func (x *bloberStreamClient) Send(m *StreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bloberStreamClient) CloseAndRecv() (*StreamRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Blober service

type BloberServer interface {
	Put(context.Context, *PutReq) (*PutRes, error)
	Stream(Blober_StreamServer) error
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

func _Blober_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BloberServer).Stream(&bloberStreamServer{stream})
}

type Blober_StreamServer interface {
	SendAndClose(*StreamRes) error
	Recv() (*StreamReq, error)
	grpc.ServerStream
}

type bloberStreamServer struct {
	grpc.ServerStream
}

func (x *bloberStreamServer) SendAndClose(m *StreamRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bloberStreamServer) Recv() (*StreamReq, error) {
	m := new(StreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Blober_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Blober",
	HandlerType: (*BloberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Blober_Put_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Blober_Stream_Handler,
			ClientStreams: true,
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

func (m *StreamReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamReq) MarshalTo(dAtA []byte) (int, error) {
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

func (m *StreamReq_Name) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintService(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	return i, nil
}
func (m *StreamReq_Blob) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Blob != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Blob)))
		i += copy(dAtA[i:], m.Blob)
	}
	return i, nil
}
func (m *StreamRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamRes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
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

func (m *StreamReq) Size() (n int) {
	var l int
	_ = l
	if m.Phase != nil {
		n += m.Phase.Size()
	}
	return n
}

func (m *StreamReq_Name) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovService(uint64(l))
	return n
}
func (m *StreamReq_Blob) Size() (n int) {
	var l int
	_ = l
	if m.Blob != nil {
		l = len(m.Blob)
		n += 1 + l + sovService(uint64(l))
	}
	return n
}
func (m *StreamRes) Size() (n int) {
	var l int
	_ = l
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
func (m *StreamReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StreamReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamReq: illegal tag %d (wire type %d)", fieldNum, wire)
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
			m.Phase = &StreamReq_Name{string(dAtA[iNdEx:postIndex])}
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
			m.Phase = &StreamReq_Blob{v}
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
func (m *StreamRes) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StreamRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamRes: illegal tag %d (wire type %d)", fieldNum, wire)
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
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x0c, 0xb8,
	0xd8, 0x02, 0x4a, 0x4b, 0x82, 0x52, 0x0b, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x52, 0x4e, 0x7e, 0x92, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0xad, 0xc4, 0x01, 0xd5, 0x51, 0xac, 0xe4, 0xc0, 0xc5, 0x19,
	0x5c, 0x52, 0x94, 0x9a, 0x98, 0x0b, 0xd2, 0x2e, 0x82, 0xac, 0xdd, 0x83, 0x01, 0x6a, 0x80, 0x08,
	0xb2, 0x01, 0x20, 0x51, 0x10, 0xcf, 0x89, 0x9d, 0x8b, 0xb5, 0x20, 0x23, 0xb1, 0x38, 0x55, 0x89,
	0x1b, 0x61, 0x42, 0xb1, 0x51, 0x2a, 0x17, 0x9b, 0x53, 0x4e, 0x7e, 0x52, 0x6a, 0x91, 0x90, 0x3a,
	0x17, 0x73, 0x40, 0x69, 0x89, 0x10, 0xbf, 0x1e, 0xcc, 0xd1, 0x10, 0x27, 0x4a, 0xa1, 0x09, 0x14,
	0x0b, 0x19, 0x71, 0xb1, 0x41, 0xf4, 0x0b, 0x09, 0xc1, 0xa5, 0xe0, 0x4e, 0x92, 0xc2, 0x14, 0x2b,
	0xd6, 0x60, 0x74, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x87, 0x89, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x9c, 0xdc, 0xb7, 0x0b, 0x24, 0x01, 0x00, 0x00,
}
