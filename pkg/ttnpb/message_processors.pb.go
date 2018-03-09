// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/message_processors.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProcessUplinkMessageRequest struct {
	DeviceModel EndDeviceModel `protobuf:"bytes,1,opt,name=device_model,json=deviceModel" json:"device_model"`
	Message     UplinkMessage  `protobuf:"bytes,2,opt,name=message" json:"message"`
	Parameter   string         `protobuf:"bytes,3,opt,name=parameter,proto3" json:"parameter,omitempty"`
}

func (m *ProcessUplinkMessageRequest) Reset()         { *m = ProcessUplinkMessageRequest{} }
func (m *ProcessUplinkMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessUplinkMessageRequest) ProtoMessage()    {}
func (*ProcessUplinkMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorMessageProcessors, []int{0}
}

func (m *ProcessUplinkMessageRequest) GetDeviceModel() EndDeviceModel {
	if m != nil {
		return m.DeviceModel
	}
	return EndDeviceModel{}
}

func (m *ProcessUplinkMessageRequest) GetMessage() UplinkMessage {
	if m != nil {
		return m.Message
	}
	return UplinkMessage{}
}

func (m *ProcessUplinkMessageRequest) GetParameter() string {
	if m != nil {
		return m.Parameter
	}
	return ""
}

type ProcessDownlinkMessageRequest struct {
	DeviceModel EndDeviceModel  `protobuf:"bytes,1,opt,name=device_model,json=deviceModel" json:"device_model"`
	Message     DownlinkMessage `protobuf:"bytes,2,opt,name=message" json:"message"`
	Parameter   string          `protobuf:"bytes,3,opt,name=parameter,proto3" json:"parameter,omitempty"`
}

func (m *ProcessDownlinkMessageRequest) Reset()         { *m = ProcessDownlinkMessageRequest{} }
func (m *ProcessDownlinkMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessDownlinkMessageRequest) ProtoMessage()    {}
func (*ProcessDownlinkMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorMessageProcessors, []int{1}
}

func (m *ProcessDownlinkMessageRequest) GetDeviceModel() EndDeviceModel {
	if m != nil {
		return m.DeviceModel
	}
	return EndDeviceModel{}
}

func (m *ProcessDownlinkMessageRequest) GetMessage() DownlinkMessage {
	if m != nil {
		return m.Message
	}
	return DownlinkMessage{}
}

func (m *ProcessDownlinkMessageRequest) GetParameter() string {
	if m != nil {
		return m.Parameter
	}
	return ""
}

func init() {
	proto.RegisterType((*ProcessUplinkMessageRequest)(nil), "ttn.v3.ProcessUplinkMessageRequest")
	golang_proto.RegisterType((*ProcessUplinkMessageRequest)(nil), "ttn.v3.ProcessUplinkMessageRequest")
	proto.RegisterType((*ProcessDownlinkMessageRequest)(nil), "ttn.v3.ProcessDownlinkMessageRequest")
	golang_proto.RegisterType((*ProcessDownlinkMessageRequest)(nil), "ttn.v3.ProcessDownlinkMessageRequest")
}
func (this *ProcessUplinkMessageRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ProcessUplinkMessageRequest)
	if !ok {
		that2, ok := that.(ProcessUplinkMessageRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ProcessUplinkMessageRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ProcessUplinkMessageRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ProcessUplinkMessageRequest but is not nil && this == nil")
	}
	if !this.DeviceModel.Equal(&that1.DeviceModel) {
		return fmt.Errorf("DeviceModel this(%v) Not Equal that(%v)", this.DeviceModel, that1.DeviceModel)
	}
	if !this.Message.Equal(&that1.Message) {
		return fmt.Errorf("Message this(%v) Not Equal that(%v)", this.Message, that1.Message)
	}
	if this.Parameter != that1.Parameter {
		return fmt.Errorf("Parameter this(%v) Not Equal that(%v)", this.Parameter, that1.Parameter)
	}
	return nil
}
func (this *ProcessUplinkMessageRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProcessUplinkMessageRequest)
	if !ok {
		that2, ok := that.(ProcessUplinkMessageRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.DeviceModel.Equal(&that1.DeviceModel) {
		return false
	}
	if !this.Message.Equal(&that1.Message) {
		return false
	}
	if this.Parameter != that1.Parameter {
		return false
	}
	return true
}
func (this *ProcessDownlinkMessageRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ProcessDownlinkMessageRequest)
	if !ok {
		that2, ok := that.(ProcessDownlinkMessageRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ProcessDownlinkMessageRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ProcessDownlinkMessageRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ProcessDownlinkMessageRequest but is not nil && this == nil")
	}
	if !this.DeviceModel.Equal(&that1.DeviceModel) {
		return fmt.Errorf("DeviceModel this(%v) Not Equal that(%v)", this.DeviceModel, that1.DeviceModel)
	}
	if !this.Message.Equal(&that1.Message) {
		return fmt.Errorf("Message this(%v) Not Equal that(%v)", this.Message, that1.Message)
	}
	if this.Parameter != that1.Parameter {
		return fmt.Errorf("Parameter this(%v) Not Equal that(%v)", this.Parameter, that1.Parameter)
	}
	return nil
}
func (this *ProcessDownlinkMessageRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProcessDownlinkMessageRequest)
	if !ok {
		that2, ok := that.(ProcessDownlinkMessageRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.DeviceModel.Equal(&that1.DeviceModel) {
		return false
	}
	if !this.Message.Equal(&that1.Message) {
		return false
	}
	if this.Parameter != that1.Parameter {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UplinkMessageProcessor service

type UplinkMessageProcessorClient interface {
	Process(ctx context.Context, in *ProcessUplinkMessageRequest, opts ...grpc.CallOption) (*UplinkMessage, error)
}

type uplinkMessageProcessorClient struct {
	cc *grpc.ClientConn
}

func NewUplinkMessageProcessorClient(cc *grpc.ClientConn) UplinkMessageProcessorClient {
	return &uplinkMessageProcessorClient{cc}
}

func (c *uplinkMessageProcessorClient) Process(ctx context.Context, in *ProcessUplinkMessageRequest, opts ...grpc.CallOption) (*UplinkMessage, error) {
	out := new(UplinkMessage)
	err := grpc.Invoke(ctx, "/ttn.v3.UplinkMessageProcessor/Process", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UplinkMessageProcessor service

type UplinkMessageProcessorServer interface {
	Process(context.Context, *ProcessUplinkMessageRequest) (*UplinkMessage, error)
}

func RegisterUplinkMessageProcessorServer(s *grpc.Server, srv UplinkMessageProcessorServer) {
	s.RegisterService(&_UplinkMessageProcessor_serviceDesc, srv)
}

func _UplinkMessageProcessor_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessUplinkMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UplinkMessageProcessorServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.v3.UplinkMessageProcessor/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UplinkMessageProcessorServer).Process(ctx, req.(*ProcessUplinkMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UplinkMessageProcessor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.v3.UplinkMessageProcessor",
	HandlerType: (*UplinkMessageProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _UplinkMessageProcessor_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/TheThingsNetwork/ttn/api/message_processors.proto",
}

// Client API for DownlinkMessageProcessor service

type DownlinkMessageProcessorClient interface {
	Process(ctx context.Context, in *ProcessDownlinkMessageRequest, opts ...grpc.CallOption) (*DownlinkMessage, error)
}

type downlinkMessageProcessorClient struct {
	cc *grpc.ClientConn
}

func NewDownlinkMessageProcessorClient(cc *grpc.ClientConn) DownlinkMessageProcessorClient {
	return &downlinkMessageProcessorClient{cc}
}

func (c *downlinkMessageProcessorClient) Process(ctx context.Context, in *ProcessDownlinkMessageRequest, opts ...grpc.CallOption) (*DownlinkMessage, error) {
	out := new(DownlinkMessage)
	err := grpc.Invoke(ctx, "/ttn.v3.DownlinkMessageProcessor/Process", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DownlinkMessageProcessor service

type DownlinkMessageProcessorServer interface {
	Process(context.Context, *ProcessDownlinkMessageRequest) (*DownlinkMessage, error)
}

func RegisterDownlinkMessageProcessorServer(s *grpc.Server, srv DownlinkMessageProcessorServer) {
	s.RegisterService(&_DownlinkMessageProcessor_serviceDesc, srv)
}

func _DownlinkMessageProcessor_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessDownlinkMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownlinkMessageProcessorServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.v3.DownlinkMessageProcessor/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownlinkMessageProcessorServer).Process(ctx, req.(*ProcessDownlinkMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DownlinkMessageProcessor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.v3.DownlinkMessageProcessor",
	HandlerType: (*DownlinkMessageProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _DownlinkMessageProcessor_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/TheThingsNetwork/ttn/api/message_processors.proto",
}

func (m *ProcessUplinkMessageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessUplinkMessageRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMessageProcessors(dAtA, i, uint64(m.DeviceModel.Size()))
	n1, err := m.DeviceModel.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintMessageProcessors(dAtA, i, uint64(m.Message.Size()))
	n2, err := m.Message.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Parameter) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessageProcessors(dAtA, i, uint64(len(m.Parameter)))
		i += copy(dAtA[i:], m.Parameter)
	}
	return i, nil
}

func (m *ProcessDownlinkMessageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessDownlinkMessageRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMessageProcessors(dAtA, i, uint64(m.DeviceModel.Size()))
	n3, err := m.DeviceModel.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x12
	i++
	i = encodeVarintMessageProcessors(dAtA, i, uint64(m.Message.Size()))
	n4, err := m.Message.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Parameter) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessageProcessors(dAtA, i, uint64(len(m.Parameter)))
		i += copy(dAtA[i:], m.Parameter)
	}
	return i, nil
}

func encodeVarintMessageProcessors(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedProcessUplinkMessageRequest(r randyMessageProcessors, easy bool) *ProcessUplinkMessageRequest {
	this := &ProcessUplinkMessageRequest{}
	v1 := NewPopulatedEndDeviceModel(r, easy)
	this.DeviceModel = *v1
	v2 := NewPopulatedUplinkMessage(r, easy)
	this.Message = *v2
	this.Parameter = randStringMessageProcessors(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedProcessDownlinkMessageRequest(r randyMessageProcessors, easy bool) *ProcessDownlinkMessageRequest {
	this := &ProcessDownlinkMessageRequest{}
	v3 := NewPopulatedEndDeviceModel(r, easy)
	this.DeviceModel = *v3
	v4 := NewPopulatedDownlinkMessage(r, easy)
	this.Message = *v4
	this.Parameter = randStringMessageProcessors(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyMessageProcessors interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMessageProcessors(r randyMessageProcessors) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMessageProcessors(r randyMessageProcessors) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneMessageProcessors(r)
	}
	return string(tmps)
}
func randUnrecognizedMessageProcessors(r randyMessageProcessors, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMessageProcessors(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMessageProcessors(dAtA []byte, r randyMessageProcessors, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMessageProcessors(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateMessageProcessors(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateMessageProcessors(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMessageProcessors(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMessageProcessors(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMessageProcessors(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMessageProcessors(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ProcessUplinkMessageRequest) Size() (n int) {
	var l int
	_ = l
	l = m.DeviceModel.Size()
	n += 1 + l + sovMessageProcessors(uint64(l))
	l = m.Message.Size()
	n += 1 + l + sovMessageProcessors(uint64(l))
	l = len(m.Parameter)
	if l > 0 {
		n += 1 + l + sovMessageProcessors(uint64(l))
	}
	return n
}

func (m *ProcessDownlinkMessageRequest) Size() (n int) {
	var l int
	_ = l
	l = m.DeviceModel.Size()
	n += 1 + l + sovMessageProcessors(uint64(l))
	l = m.Message.Size()
	n += 1 + l + sovMessageProcessors(uint64(l))
	l = len(m.Parameter)
	if l > 0 {
		n += 1 + l + sovMessageProcessors(uint64(l))
	}
	return n
}

func sovMessageProcessors(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessageProcessors(x uint64) (n int) {
	return sovMessageProcessors((x << 1) ^ uint64((int64(x) >> 63)))
}
func (m *ProcessUplinkMessageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageProcessors
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
			return fmt.Errorf("proto: ProcessUplinkMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessUplinkMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceModel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageProcessors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageProcessors
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeviceModel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageProcessors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageProcessors
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageProcessors
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
				return ErrInvalidLengthMessageProcessors
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parameter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessageProcessors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessageProcessors
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
func (m *ProcessDownlinkMessageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageProcessors
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
			return fmt.Errorf("proto: ProcessDownlinkMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessDownlinkMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceModel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageProcessors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageProcessors
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeviceModel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageProcessors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageProcessors
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageProcessors
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
				return ErrInvalidLengthMessageProcessors
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parameter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessageProcessors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessageProcessors
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
func skipMessageProcessors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessageProcessors
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
					return 0, ErrIntOverflowMessageProcessors
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
					return 0, ErrIntOverflowMessageProcessors
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
				return 0, ErrInvalidLengthMessageProcessors
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessageProcessors
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
				next, err := skipMessageProcessors(dAtA[start:])
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
	ErrInvalidLengthMessageProcessors = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessageProcessors   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/message_processors.proto", fileDescriptorMessageProcessors)
}
func init() {
	golang_proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/message_processors.proto", fileDescriptorMessageProcessors)
}

var fileDescriptorMessageProcessors = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x3f, 0x6c, 0xd3, 0x40,
	0x14, 0xc6, 0xef, 0x01, 0x6a, 0xd5, 0x2b, 0x93, 0x25, 0x4a, 0x64, 0xe0, 0x51, 0x15, 0x21, 0x95,
	0x01, 0x5b, 0xa4, 0x54, 0x2c, 0x48, 0x48, 0xa8, 0x1d, 0x18, 0x8a, 0xaa, 0xa8, 0x2c, 0x2c, 0xd1,
	0x25, 0xbe, 0x3a, 0x56, 0x92, 0x3b, 0xe3, 0xbb, 0x34, 0x6b, 0xc7, 0x8e, 0x8c, 0x6c, 0x30, 0x76,
	0xac, 0xc4, 0x40, 0xc7, 0x8e, 0x19, 0x3b, 0x76, 0x42, 0xf5, 0xdd, 0xd2, 0xb1, 0x63, 0x47, 0x54,
	0xff, 0x51, 0xa8, 0x15, 0x20, 0x19, 0x98, 0xec, 0xf7, 0xee, 0x7e, 0xfe, 0xbe, 0xf7, 0xf9, 0x8e,
	0xbe, 0x0e, 0x23, 0xdd, 0x19, 0xb4, 0xbc, 0xb6, 0xec, 0xfb, 0x3b, 0x1d, 0xbe, 0xd3, 0x89, 0x44,
	0xa8, 0xde, 0x73, 0x3d, 0x94, 0x49, 0xd7, 0xd7, 0x5a, 0xf8, 0x2c, 0x8e, 0xfc, 0x3e, 0x57, 0x8a,
	0x85, 0xbc, 0x19, 0x27, 0xb2, 0xcd, 0x95, 0x92, 0x89, 0xf2, 0xe2, 0x44, 0x6a, 0xe9, 0xcc, 0x69,
	0x2d, 0xbc, 0xbd, 0x35, 0xf7, 0xe5, 0x34, 0x5f, 0xe1, 0x22, 0x68, 0x06, 0x7c, 0x2f, 0x6a, 0xf3,
	0x9c, 0x76, 0xd7, 0xa7, 0xa1, 0xa2, 0x80, 0x0b, 0x1d, 0xed, 0x46, 0xbc, 0x14, 0x75, 0x5f, 0x4c,
	0x83, 0xf5, 0x64, 0xc2, 0x86, 0x4c, 0x14, 0x48, 0x7d, 0x86, 0x29, 0xd5, 0x6c, 0x8c, 0x66, 0x01,
	0xd3, 0xac, 0x60, 0x9e, 0xff, 0xc6, 0x84, 0x32, 0x94, 0x7e, 0xd6, 0x6e, 0x0d, 0x76, 0xb3, 0x2a,
	0x2b, 0xb2, 0xb7, 0x7c, 0xfb, 0xca, 0x77, 0xa0, 0x0f, 0xb6, 0xf3, 0x4c, 0x3f, 0xc4, 0xbd, 0x48,
	0x74, 0xb7, 0x72, 0x0b, 0x0d, 0xfe, 0x69, 0xc0, 0x95, 0x76, 0xde, 0xd0, 0xbb, 0x79, 0x60, 0xcd,
	0xbe, 0x0c, 0x78, 0xaf, 0x06, 0xcb, 0xb0, 0xba, 0x58, 0x5f, 0xf2, 0xf2, 0xd4, 0xbd, 0x4d, 0x11,
	0x6c, 0x64, 0xcb, 0x5b, 0xd7, 0xab, 0x6f, 0xef, 0x8c, 0x7e, 0x3e, 0x26, 0x8d, 0xc5, 0x60, 0xdc,
	0x72, 0xd6, 0xe9, 0x7c, 0x31, 0x55, 0xed, 0x56, 0xc6, 0xde, 0x2b, 0xd9, 0x1b, 0x7a, 0x05, 0x5a,
	0xee, 0x75, 0x1e, 0xd2, 0x85, 0x98, 0x25, 0xac, 0xcf, 0x35, 0x4f, 0x6a, 0xb7, 0x97, 0x61, 0x75,
	0xa1, 0x31, 0x6e, 0xac, 0xfc, 0x00, 0xfa, 0xa8, 0x70, 0xbd, 0x21, 0x87, 0xe2, 0x7f, 0xf8, 0x7e,
	0x55, 0xf5, 0x7d, 0xbf, 0x64, 0x2b, 0x8a, 0x33, 0x39, 0xaf, 0x37, 0xe9, 0xd2, 0x8d, 0xb9, 0xb7,
	0xcb, 0xf3, 0xec, 0x6c, 0xd2, 0xf9, 0xa2, 0x70, 0x9e, 0x94, 0x52, 0x7f, 0xf9, 0x33, 0xee, 0xe4,
	0x1c, 0xeb, 0x9c, 0xd6, 0x2a, 0x06, 0xc7, 0x12, 0xef, 0xc6, 0x12, 0x4f, 0x2b, 0x12, 0x93, 0x63,
	0x74, 0xff, 0x38, 0xf4, 0x57, 0x18, 0xa5, 0x08, 0xa7, 0x29, 0xc2, 0x59, 0x8a, 0x70, 0x9e, 0x22,
	0x5c, 0xa4, 0x48, 0x2e, 0x53, 0x24, 0x57, 0x29, 0xc2, 0xbe, 0x41, 0x72, 0x60, 0x90, 0x1c, 0x1a,
	0x84, 0x23, 0x83, 0xe4, 0xd8, 0x20, 0x9c, 0x18, 0x84, 0x91, 0x41, 0x38, 0x35, 0x08, 0x67, 0x06,
	0xc9, 0xb9, 0x41, 0xb8, 0x30, 0x48, 0x2e, 0x0d, 0xc2, 0x95, 0x41, 0xb2, 0x6f, 0x91, 0x1c, 0x58,
	0x84, 0xcf, 0x16, 0xc9, 0x17, 0x8b, 0xf0, 0xcd, 0x22, 0x39, 0xb4, 0x48, 0x8e, 0x2c, 0xc2, 0xb1,
	0x45, 0x38, 0xb1, 0x08, 0x1f, 0x9f, 0xfd, 0xeb, 0x46, 0xc4, 0xdd, 0xf0, 0xfa, 0x19, 0xb7, 0x5a,
	0x73, 0xd9, 0x01, 0x5f, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x36, 0x5f, 0xac, 0x5f, 0x04,
	0x00, 0x00,
}
