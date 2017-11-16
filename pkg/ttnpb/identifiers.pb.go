// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/identifiers.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_TheThingsNetwork_ttn_pkg_types "github.com/TheThingsNetwork/ttn/pkg/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// UserIdentifier is the message that is used to identify an user by ID.
type UserIdentifier struct {
	// user_id is the ID of the user.
	UserID string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (m *UserIdentifier) Reset()                    { *m = UserIdentifier{} }
func (m *UserIdentifier) String() string            { return proto.CompactTextString(m) }
func (*UserIdentifier) ProtoMessage()               {}
func (*UserIdentifier) Descriptor() ([]byte, []int) { return fileDescriptorIdentifiers, []int{0} }

func (m *UserIdentifier) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

// ApplicationIdentifier is the message that is used to identify an application.
type ApplicationIdentifier struct {
	// TTN Application ID.
	ApplicationID string `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
}

func (m *ApplicationIdentifier) Reset()                    { *m = ApplicationIdentifier{} }
func (m *ApplicationIdentifier) String() string            { return proto.CompactTextString(m) }
func (*ApplicationIdentifier) ProtoMessage()               {}
func (*ApplicationIdentifier) Descriptor() ([]byte, []int) { return fileDescriptorIdentifiers, []int{1} }

func (m *ApplicationIdentifier) GetApplicationID() string {
	if m != nil {
		return m.ApplicationID
	}
	return ""
}

// GatewayIdentifier is the message that is used to identify a gateway.
type GatewayIdentifier struct {
	// TTN Gateway ID.
	GatewayID string `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
}

func (m *GatewayIdentifier) Reset()                    { *m = GatewayIdentifier{} }
func (m *GatewayIdentifier) String() string            { return proto.CompactTextString(m) }
func (*GatewayIdentifier) ProtoMessage()               {}
func (*GatewayIdentifier) Descriptor() ([]byte, []int) { return fileDescriptorIdentifiers, []int{2} }

func (m *GatewayIdentifier) GetGatewayID() string {
	if m != nil {
		return m.GatewayID
	}
	return ""
}

// End device identifiers are carried with uplink and downlink messages.
// Unknown fields are left empty.
type EndDeviceIdentifiers struct {
	// TTN Device ID.
	DeviceID string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// TTN Application ID.
	ApplicationID string `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	// LoRaWAN DevEUI.
	DevEUI *github_com_TheThingsNetwork_ttn_pkg_types.EUI64 `protobuf:"bytes,4,opt,name=dev_eui,json=devEui,proto3,customtype=github.com/TheThingsNetwork/ttn/pkg/types.EUI64" json:"dev_eui,omitempty"`
	// LoRaWAN JoinEUI (or AppEUI for LoRaWAN 1.0 end devices).
	JoinEUI *github_com_TheThingsNetwork_ttn_pkg_types.EUI64 `protobuf:"bytes,5,opt,name=join_eui,json=joinEui,proto3,customtype=github.com/TheThingsNetwork/ttn/pkg/types.EUI64" json:"join_eui,omitempty"`
	// LoRaWAN DevAddr.
	DevAddr *github_com_TheThingsNetwork_ttn_pkg_types.DevAddr `protobuf:"bytes,6,opt,name=dev_addr,json=devAddr,proto3,customtype=github.com/TheThingsNetwork/ttn/pkg/types.DevAddr" json:"dev_addr,omitempty"`
}

func (m *EndDeviceIdentifiers) Reset()                    { *m = EndDeviceIdentifiers{} }
func (m *EndDeviceIdentifiers) String() string            { return proto.CompactTextString(m) }
func (*EndDeviceIdentifiers) ProtoMessage()               {}
func (*EndDeviceIdentifiers) Descriptor() ([]byte, []int) { return fileDescriptorIdentifiers, []int{3} }

func (m *EndDeviceIdentifiers) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *EndDeviceIdentifiers) GetApplicationID() string {
	if m != nil {
		return m.ApplicationID
	}
	return ""
}

// ClientIdentifier is the message that is used to identify a client.
type ClientIdentifier struct {
	// TTN Client ID.
	ClientID string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *ClientIdentifier) Reset()                    { *m = ClientIdentifier{} }
func (m *ClientIdentifier) String() string            { return proto.CompactTextString(m) }
func (*ClientIdentifier) ProtoMessage()               {}
func (*ClientIdentifier) Descriptor() ([]byte, []int) { return fileDescriptorIdentifiers, []int{4} }

func (m *ClientIdentifier) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func init() {
	proto.RegisterType((*UserIdentifier)(nil), "ttn.v3.UserIdentifier")
	golang_proto.RegisterType((*UserIdentifier)(nil), "ttn.v3.UserIdentifier")
	proto.RegisterType((*ApplicationIdentifier)(nil), "ttn.v3.ApplicationIdentifier")
	golang_proto.RegisterType((*ApplicationIdentifier)(nil), "ttn.v3.ApplicationIdentifier")
	proto.RegisterType((*GatewayIdentifier)(nil), "ttn.v3.GatewayIdentifier")
	golang_proto.RegisterType((*GatewayIdentifier)(nil), "ttn.v3.GatewayIdentifier")
	proto.RegisterType((*EndDeviceIdentifiers)(nil), "ttn.v3.EndDeviceIdentifiers")
	golang_proto.RegisterType((*EndDeviceIdentifiers)(nil), "ttn.v3.EndDeviceIdentifiers")
	proto.RegisterType((*ClientIdentifier)(nil), "ttn.v3.ClientIdentifier")
	golang_proto.RegisterType((*ClientIdentifier)(nil), "ttn.v3.ClientIdentifier")
}
func (this *UserIdentifier) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*UserIdentifier)
	if !ok {
		that2, ok := that.(UserIdentifier)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *UserIdentifier")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *UserIdentifier but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *UserIdentifier but is not nil && this == nil")
	}
	if this.UserID != that1.UserID {
		return fmt.Errorf("UserID this(%v) Not Equal that(%v)", this.UserID, that1.UserID)
	}
	return nil
}
func (this *UserIdentifier) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*UserIdentifier)
	if !ok {
		that2, ok := that.(UserIdentifier)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.UserID != that1.UserID {
		return false
	}
	return true
}
func (this *ApplicationIdentifier) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ApplicationIdentifier)
	if !ok {
		that2, ok := that.(ApplicationIdentifier)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ApplicationIdentifier")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ApplicationIdentifier but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ApplicationIdentifier but is not nil && this == nil")
	}
	if this.ApplicationID != that1.ApplicationID {
		return fmt.Errorf("ApplicationID this(%v) Not Equal that(%v)", this.ApplicationID, that1.ApplicationID)
	}
	return nil
}
func (this *ApplicationIdentifier) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ApplicationIdentifier)
	if !ok {
		that2, ok := that.(ApplicationIdentifier)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ApplicationID != that1.ApplicationID {
		return false
	}
	return true
}
func (this *GatewayIdentifier) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*GatewayIdentifier)
	if !ok {
		that2, ok := that.(GatewayIdentifier)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *GatewayIdentifier")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *GatewayIdentifier but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *GatewayIdentifier but is not nil && this == nil")
	}
	if this.GatewayID != that1.GatewayID {
		return fmt.Errorf("GatewayID this(%v) Not Equal that(%v)", this.GatewayID, that1.GatewayID)
	}
	return nil
}
func (this *GatewayIdentifier) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GatewayIdentifier)
	if !ok {
		that2, ok := that.(GatewayIdentifier)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.GatewayID != that1.GatewayID {
		return false
	}
	return true
}
func (this *EndDeviceIdentifiers) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*EndDeviceIdentifiers)
	if !ok {
		that2, ok := that.(EndDeviceIdentifiers)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *EndDeviceIdentifiers")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *EndDeviceIdentifiers but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *EndDeviceIdentifiers but is not nil && this == nil")
	}
	if this.DeviceID != that1.DeviceID {
		return fmt.Errorf("DeviceID this(%v) Not Equal that(%v)", this.DeviceID, that1.DeviceID)
	}
	if this.ApplicationID != that1.ApplicationID {
		return fmt.Errorf("ApplicationID this(%v) Not Equal that(%v)", this.ApplicationID, that1.ApplicationID)
	}
	if that1.DevEUI == nil {
		if this.DevEUI != nil {
			return fmt.Errorf("this.DevEUI != nil && that1.DevEUI == nil")
		}
	} else if !this.DevEUI.Equal(*that1.DevEUI) {
		return fmt.Errorf("DevEUI this(%v) Not Equal that(%v)", this.DevEUI, that1.DevEUI)
	}
	if that1.JoinEUI == nil {
		if this.JoinEUI != nil {
			return fmt.Errorf("this.JoinEUI != nil && that1.JoinEUI == nil")
		}
	} else if !this.JoinEUI.Equal(*that1.JoinEUI) {
		return fmt.Errorf("JoinEUI this(%v) Not Equal that(%v)", this.JoinEUI, that1.JoinEUI)
	}
	if that1.DevAddr == nil {
		if this.DevAddr != nil {
			return fmt.Errorf("this.DevAddr != nil && that1.DevAddr == nil")
		}
	} else if !this.DevAddr.Equal(*that1.DevAddr) {
		return fmt.Errorf("DevAddr this(%v) Not Equal that(%v)", this.DevAddr, that1.DevAddr)
	}
	return nil
}
func (this *EndDeviceIdentifiers) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*EndDeviceIdentifiers)
	if !ok {
		that2, ok := that.(EndDeviceIdentifiers)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.DeviceID != that1.DeviceID {
		return false
	}
	if this.ApplicationID != that1.ApplicationID {
		return false
	}
	if that1.DevEUI == nil {
		if this.DevEUI != nil {
			return false
		}
	} else if !this.DevEUI.Equal(*that1.DevEUI) {
		return false
	}
	if that1.JoinEUI == nil {
		if this.JoinEUI != nil {
			return false
		}
	} else if !this.JoinEUI.Equal(*that1.JoinEUI) {
		return false
	}
	if that1.DevAddr == nil {
		if this.DevAddr != nil {
			return false
		}
	} else if !this.DevAddr.Equal(*that1.DevAddr) {
		return false
	}
	return true
}
func (this *ClientIdentifier) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ClientIdentifier)
	if !ok {
		that2, ok := that.(ClientIdentifier)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ClientIdentifier")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ClientIdentifier but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ClientIdentifier but is not nil && this == nil")
	}
	if this.ClientID != that1.ClientID {
		return fmt.Errorf("ClientID this(%v) Not Equal that(%v)", this.ClientID, that1.ClientID)
	}
	return nil
}
func (this *ClientIdentifier) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ClientIdentifier)
	if !ok {
		that2, ok := that.(ClientIdentifier)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ClientID != that1.ClientID {
		return false
	}
	return true
}
func (m *UserIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserIdentifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UserID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIdentifiers(dAtA, i, uint64(len(m.UserID)))
		i += copy(dAtA[i:], m.UserID)
	}
	return i, nil
}

func (m *ApplicationIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplicationIdentifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ApplicationID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIdentifiers(dAtA, i, uint64(len(m.ApplicationID)))
		i += copy(dAtA[i:], m.ApplicationID)
	}
	return i, nil
}

func (m *GatewayIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GatewayIdentifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.GatewayID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIdentifiers(dAtA, i, uint64(len(m.GatewayID)))
		i += copy(dAtA[i:], m.GatewayID)
	}
	return i, nil
}

func (m *EndDeviceIdentifiers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EndDeviceIdentifiers) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DeviceID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIdentifiers(dAtA, i, uint64(len(m.DeviceID)))
		i += copy(dAtA[i:], m.DeviceID)
	}
	if len(m.ApplicationID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintIdentifiers(dAtA, i, uint64(len(m.ApplicationID)))
		i += copy(dAtA[i:], m.ApplicationID)
	}
	if m.DevEUI != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintIdentifiers(dAtA, i, uint64(m.DevEUI.Size()))
		n1, err := m.DevEUI.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.JoinEUI != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintIdentifiers(dAtA, i, uint64(m.JoinEUI.Size()))
		n2, err := m.JoinEUI.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.DevAddr != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintIdentifiers(dAtA, i, uint64(m.DevAddr.Size()))
		n3, err := m.DevAddr.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *ClientIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientIdentifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ClientID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIdentifiers(dAtA, i, uint64(len(m.ClientID)))
		i += copy(dAtA[i:], m.ClientID)
	}
	return i, nil
}

func encodeVarintIdentifiers(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedUserIdentifier(r randyIdentifiers, easy bool) *UserIdentifier {
	this := &UserIdentifier{}
	this.UserID = string(randStringIdentifiers(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedApplicationIdentifier(r randyIdentifiers, easy bool) *ApplicationIdentifier {
	this := &ApplicationIdentifier{}
	this.ApplicationID = string(randStringIdentifiers(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedGatewayIdentifier(r randyIdentifiers, easy bool) *GatewayIdentifier {
	this := &GatewayIdentifier{}
	this.GatewayID = string(randStringIdentifiers(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedClientIdentifier(r randyIdentifiers, easy bool) *ClientIdentifier {
	this := &ClientIdentifier{}
	this.ClientID = string(randStringIdentifiers(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyIdentifiers interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneIdentifiers(r randyIdentifiers) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringIdentifiers(r randyIdentifiers) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneIdentifiers(r)
	}
	return string(tmps)
}
func randUnrecognizedIdentifiers(r randyIdentifiers, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldIdentifiers(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldIdentifiers(dAtA []byte, r randyIdentifiers, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateIdentifiers(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateIdentifiers(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateIdentifiers(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateIdentifiers(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateIdentifiers(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateIdentifiers(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateIdentifiers(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *UserIdentifier) Size() (n int) {
	var l int
	_ = l
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovIdentifiers(uint64(l))
	}
	return n
}

func (m *ApplicationIdentifier) Size() (n int) {
	var l int
	_ = l
	l = len(m.ApplicationID)
	if l > 0 {
		n += 1 + l + sovIdentifiers(uint64(l))
	}
	return n
}

func (m *GatewayIdentifier) Size() (n int) {
	var l int
	_ = l
	l = len(m.GatewayID)
	if l > 0 {
		n += 1 + l + sovIdentifiers(uint64(l))
	}
	return n
}

func (m *EndDeviceIdentifiers) Size() (n int) {
	var l int
	_ = l
	l = len(m.DeviceID)
	if l > 0 {
		n += 1 + l + sovIdentifiers(uint64(l))
	}
	l = len(m.ApplicationID)
	if l > 0 {
		n += 1 + l + sovIdentifiers(uint64(l))
	}
	if m.DevEUI != nil {
		l = m.DevEUI.Size()
		n += 1 + l + sovIdentifiers(uint64(l))
	}
	if m.JoinEUI != nil {
		l = m.JoinEUI.Size()
		n += 1 + l + sovIdentifiers(uint64(l))
	}
	if m.DevAddr != nil {
		l = m.DevAddr.Size()
		n += 1 + l + sovIdentifiers(uint64(l))
	}
	return n
}

func (m *ClientIdentifier) Size() (n int) {
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovIdentifiers(uint64(l))
	}
	return n
}

func sovIdentifiers(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIdentifiers(x uint64) (n int) {
	return sovIdentifiers(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserIdentifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentifiers
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
			return fmt.Errorf("proto: UserIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifiers
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
				return ErrInvalidLengthIdentifiers
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentifiers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentifiers
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
func (m *ApplicationIdentifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentifiers
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
			return fmt.Errorf("proto: ApplicationIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplicationIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifiers
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
				return ErrInvalidLengthIdentifiers
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentifiers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentifiers
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
func (m *GatewayIdentifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentifiers
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
			return fmt.Errorf("proto: GatewayIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GatewayIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifiers
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
				return ErrInvalidLengthIdentifiers
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GatewayID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentifiers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentifiers
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
func (m *EndDeviceIdentifiers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentifiers
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
			return fmt.Errorf("proto: EndDeviceIdentifiers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EndDeviceIdentifiers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifiers
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
				return ErrInvalidLengthIdentifiers
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifiers
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
				return ErrInvalidLengthIdentifiers
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifiers
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
				return ErrInvalidLengthIdentifiers
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_pkg_types.EUI64
			m.DevEUI = &v
			if err := m.DevEUI.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JoinEUI", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifiers
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
				return ErrInvalidLengthIdentifiers
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_pkg_types.EUI64
			m.JoinEUI = &v
			if err := m.JoinEUI.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifiers
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
				return ErrInvalidLengthIdentifiers
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_pkg_types.DevAddr
			m.DevAddr = &v
			if err := m.DevAddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentifiers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentifiers
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
func (m *ClientIdentifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentifiers
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
			return fmt.Errorf("proto: ClientIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentifiers
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
				return ErrInvalidLengthIdentifiers
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentifiers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIdentifiers
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
func skipIdentifiers(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIdentifiers
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
					return 0, ErrIntOverflowIdentifiers
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
					return 0, ErrIntOverflowIdentifiers
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
				return 0, ErrInvalidLengthIdentifiers
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIdentifiers
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
				next, err := skipIdentifiers(dAtA[start:])
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
	ErrInvalidLengthIdentifiers = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIdentifiers   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/identifiers.proto", fileDescriptorIdentifiers)
}
func init() {
	golang_proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/identifiers.proto", fileDescriptorIdentifiers)
}

var fileDescriptorIdentifiers = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x3f, 0x68, 0xdb, 0x4c,
	0x18, 0xc6, 0xef, 0xbe, 0x2f, 0x95, 0xed, 0x23, 0x09, 0x8d, 0x69, 0x21, 0x74, 0x78, 0x15, 0xdc,
	0x25, 0x86, 0xd6, 0xa2, 0xa4, 0x2e, 0xa5, 0x90, 0x82, 0x5d, 0x99, 0xe2, 0x0e, 0xa5, 0x15, 0x71,
	0xa1, 0x5d, 0x8c, 0xec, 0xbb, 0xc8, 0xd7, 0xa4, 0x92, 0x90, 0x4e, 0x0a, 0xd9, 0x32, 0x95, 0x8c,
	0x1d, 0xbb, 0x35, 0x63, 0xc6, 0x8c, 0x1e, 0x33, 0x66, 0xcc, 0x18, 0x32, 0x88, 0xe8, 0xb4, 0x64,
	0xcc, 0x98, 0xb1, 0xe8, 0xe2, 0xd4, 0xa2, 0x14, 0x6a, 0xba, 0xbd, 0x7f, 0x9e, 0xe7, 0x07, 0xcf,
	0xbd, 0x1c, 0x69, 0x3a, 0x5c, 0x8c, 0xa2, 0x41, 0x63, 0xe8, 0x7d, 0x31, 0x36, 0x46, 0x6c, 0x63,
	0xc4, 0x5d, 0x27, 0x7c, 0xcb, 0xc4, 0x8e, 0x17, 0x6c, 0x19, 0x42, 0xb8, 0x86, 0xed, 0x73, 0x83,
	0x53, 0xe6, 0x0a, 0xbe, 0xc9, 0x59, 0x10, 0x36, 0xfc, 0xc0, 0x13, 0x5e, 0x55, 0x13, 0xc2, 0x6d,
	0xc4, 0x6b, 0x0f, 0x1e, 0x17, 0xec, 0x8e, 0xe7, 0x78, 0x86, 0x5a, 0x0f, 0xa2, 0x4d, 0xd5, 0xa9,
	0x46, 0x55, 0x37, 0xb6, 0x5a, 0x93, 0x2c, 0xf6, 0x42, 0x16, 0x74, 0x7f, 0xf1, 0xaa, 0x0f, 0x49,
	0x29, 0x0a, 0x59, 0xd0, 0xe7, 0x74, 0x19, 0xaf, 0xe0, 0xd5, 0x4a, 0x9b, 0xc8, 0x44, 0xd7, 0x94,
	0xc8, 0xb4, 0xb4, 0x48, 0x89, 0x6b, 0xef, 0xc9, 0xfd, 0x96, 0xef, 0x6f, 0xf3, 0xa1, 0x2d, 0xb8,
	0xe7, 0x16, 0xdc, 0xcf, 0xc9, 0xa2, 0x3d, 0x5d, 0x4c, 0x21, 0x4b, 0x32, 0xd1, 0x17, 0x8a, 0x16,
	0xd3, 0x5a, 0xb0, 0x8b, 0x84, 0x5a, 0x8b, 0x2c, 0xbd, 0xb6, 0x05, 0xdb, 0xb1, 0x77, 0x0b, 0xb8,
	0x47, 0x84, 0x38, 0x37, 0xc3, 0x29, 0x6a, 0x41, 0x26, 0x7a, 0xe5, 0x56, 0x6a, 0x5a, 0x15, 0xe7,
	0xd6, 0x55, 0xfb, 0xfa, 0x3f, 0xb9, 0xd7, 0x71, 0xa9, 0xc9, 0x62, 0x3e, 0x64, 0x53, 0x4a, 0x58,
	0xad, 0x93, 0x0a, 0x55, 0xc3, 0x29, 0x65, 0x5e, 0x26, 0x7a, 0x79, 0xa2, 0x34, 0xad, 0x32, 0x9d,
	0x78, 0xfe, 0x10, 0xe0, 0xbf, 0xd9, 0x02, 0x54, 0x3f, 0x90, 0x12, 0x65, 0x71, 0x9f, 0x45, 0x7c,
	0x79, 0x6e, 0x05, 0xaf, 0xce, 0xb7, 0xd7, 0xcf, 0x13, 0xdd, 0xf8, 0xdb, 0x35, 0xfd, 0x2d, 0xc7,
	0x10, 0xbb, 0x3e, 0x0b, 0x1b, 0x9d, 0x5e, 0xf7, 0xd9, 0xd3, 0xfc, 0xad, 0x4d, 0x16, 0x77, 0x7a,
	0x5d, 0x4b, 0xa3, 0x2c, 0xee, 0x44, 0xbc, 0xfa, 0x91, 0x94, 0x3f, 0x7b, 0xdc, 0x55, 0xe0, 0x3b,
	0x0a, 0xfc, 0xf2, 0xdf, 0xc0, 0xa5, 0x37, 0x1e, 0x77, 0x73, 0x72, 0x29, 0xe7, 0xe5, 0xe8, 0x77,
	0x24, 0x0f, 0xde, 0xb7, 0x29, 0x0d, 0x96, 0x35, 0x85, 0x6e, 0x9e, 0x27, 0xfa, 0x93, 0xd9, 0xd1,
	0x26, 0x8b, 0x5b, 0x94, 0x06, 0x56, 0x9e, 0x3c, 0x2f, 0x5e, 0xcc, 0x8d, 0x0f, 0x74, 0x54, 0x5b,
	0x27, 0x77, 0x5f, 0x6d, 0x73, 0xe6, 0x8a, 0xc2, 0x29, 0xeb, 0xa4, 0x32, 0x54, 0xb3, 0xdf, 0x6e,
	0x30, 0x11, 0x9a, 0x56, 0x79, 0x38, 0xb1, 0xb4, 0x7f, 0xe0, 0x93, 0x14, 0xf0, 0x69, 0x0a, 0xf8,
	0x2c, 0x05, 0x7c, 0x91, 0x02, 0xbe, 0x4c, 0x01, 0x5d, 0xa5, 0x80, 0xae, 0x53, 0xc0, 0x7b, 0x12,
	0xd0, 0xbe, 0x04, 0x74, 0x28, 0x01, 0x1f, 0x49, 0x40, 0x63, 0x09, 0xf8, 0x58, 0x02, 0x3e, 0x91,
	0x80, 0x4f, 0x25, 0xe0, 0x33, 0x09, 0xe8, 0x42, 0x02, 0xbe, 0x94, 0x80, 0xae, 0x24, 0xe0, 0x6b,
	0x09, 0x68, 0x2f, 0x03, 0xb4, 0x9f, 0x01, 0xfe, 0x96, 0x01, 0xfa, 0x9e, 0x01, 0x3e, 0xc8, 0x00,
	0x1d, 0x66, 0x80, 0x8e, 0x32, 0xc0, 0xe3, 0x0c, 0xf0, 0x71, 0x06, 0xf8, 0x53, 0x7d, 0xa6, 0xdc,
	0xc2, 0xf5, 0x07, 0x03, 0x4d, 0xfd, 0x9e, 0xb5, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x00, 0xfd,
	0x77, 0x0e, 0xad, 0x03, 0x00, 0x00,
}
