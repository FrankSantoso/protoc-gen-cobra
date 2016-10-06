// Code generated by protoc-gen-gogo.
// source: timer.proto
// DO NOT EDIT!

package pb

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

type TickRequest struct {
	Interval int32 `protobuf:"varint,1,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (m *TickRequest) Reset()                    { *m = TickRequest{} }
func (m *TickRequest) String() string            { return proto.CompactTextString(m) }
func (*TickRequest) ProtoMessage()               {}
func (*TickRequest) Descriptor() ([]byte, []int) { return fileDescriptorTimer, []int{0} }

type TickResponse struct {
	Time string `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (m *TickResponse) Reset()                    { *m = TickResponse{} }
func (m *TickResponse) String() string            { return proto.CompactTextString(m) }
func (*TickResponse) ProtoMessage()               {}
func (*TickResponse) Descriptor() ([]byte, []int) { return fileDescriptorTimer, []int{1} }

func init() {
	proto.RegisterType((*TickRequest)(nil), "pb.TickRequest")
	proto.RegisterType((*TickResponse)(nil), "pb.TickResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Timer service

type TimerClient interface {
	Tick(ctx context.Context, in *TickRequest, opts ...grpc.CallOption) (Timer_TickClient, error)
}

type timerClient struct {
	cc *grpc.ClientConn
}

func NewTimerClient(cc *grpc.ClientConn) TimerClient {
	return &timerClient{cc}
}

func (c *timerClient) Tick(ctx context.Context, in *TickRequest, opts ...grpc.CallOption) (Timer_TickClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Timer_serviceDesc.Streams[0], c.cc, "/pb.Timer/Tick", opts...)
	if err != nil {
		return nil, err
	}
	x := &timerTickClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Timer_TickClient interface {
	Recv() (*TickResponse, error)
	grpc.ClientStream
}

type timerTickClient struct {
	grpc.ClientStream
}

func (x *timerTickClient) Recv() (*TickResponse, error) {
	m := new(TickResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Timer service

type TimerServer interface {
	Tick(*TickRequest, Timer_TickServer) error
}

func RegisterTimerServer(s *grpc.Server, srv TimerServer) {
	s.RegisterService(&_Timer_serviceDesc, srv)
}

func _Timer_Tick_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TickRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TimerServer).Tick(m, &timerTickServer{stream})
}

type Timer_TickServer interface {
	Send(*TickResponse) error
	grpc.ServerStream
}

type timerTickServer struct {
	grpc.ServerStream
}

func (x *timerTickServer) Send(m *TickResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Timer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Timer",
	HandlerType: (*TimerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tick",
			Handler:       _Timer_Tick_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptorTimer,
}

func (m *TickRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TickRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Interval != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintTimer(data, i, uint64(m.Interval))
	}
	return i, nil
}

func (m *TickResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TickResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Time) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintTimer(data, i, uint64(len(m.Time)))
		i += copy(data[i:], m.Time)
	}
	return i, nil
}

func encodeFixed64Timer(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Timer(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTimer(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *TickRequest) Size() (n int) {
	var l int
	_ = l
	if m.Interval != 0 {
		n += 1 + sovTimer(uint64(m.Interval))
	}
	return n
}

func (m *TickResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovTimer(uint64(l))
	}
	return n
}

func sovTimer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTimer(x uint64) (n int) {
	return sovTimer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TickRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TickRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
			}
			m.Interval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Interval |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimer(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimer
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
func (m *TickResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TickResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimer(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimer
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
func skipTimer(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTimer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTimer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipTimer(data[start:])
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
	ErrInvalidLengthTimer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimer   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("timer.proto", fileDescriptorTimer) }

var fileDescriptorTimer = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xc9, 0xcc, 0x4d,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe4, 0xe2, 0x0e,
	0xc9, 0x4c, 0xce, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0xc8, 0xcc,
	0x2b, 0x49, 0x2d, 0x2a, 0x4b, 0xcc, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf3, 0x95,
	0x94, 0xb8, 0x78, 0x20, 0x4a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x84, 0xb8, 0x58, 0x40,
	0xa6, 0x81, 0xd5, 0x71, 0x06, 0x81, 0xd9, 0x46, 0x26, 0x5c, 0xac, 0x21, 0x20, 0x1b, 0x84, 0xb4,
	0xb9, 0x58, 0x40, 0x8a, 0x85, 0xf8, 0xf5, 0x0a, 0x92, 0xf4, 0x90, 0x6c, 0x90, 0x12, 0x40, 0x08,
	0x40, 0xcc, 0x31, 0x60, 0x74, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0xbb, 0xd0, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xa0, 0xa1, 0x58, 0x36, 0xb0, 0x00, 0x00, 0x00,
}