// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wire/messages.proto

/*
	Package wire is a generated protocol buffer package.

	It is generated from these files:
		wire/messages.proto

	It has these top-level messages:
		Handshake
		HandshakeResponse
		GetAddr
		GetAddrResponse
		Ping
		Pong
*/
package wire

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Handshake represents the first message between peers
type Handshake struct {
	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	SubVersion string `protobuf:"bytes,2,opt,name=subVersion,proto3" json:"subVersion,omitempty"`
	Sig        []byte `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *Handshake) Reset()                    { *m = Handshake{} }
func (*Handshake) ProtoMessage()               {}
func (*Handshake) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{0} }

func (m *Handshake) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Handshake) GetSubVersion() string {
	if m != nil {
		return m.SubVersion
	}
	return ""
}

func (m *Handshake) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// HandshakeResponse is a response to Handshake message
type HandshakeResponse struct {
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	Sig       []byte   `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *HandshakeResponse) Reset()                    { *m = HandshakeResponse{} }
func (*HandshakeResponse) ProtoMessage()               {}
func (*HandshakeResponse) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{1} }

func (m *HandshakeResponse) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *HandshakeResponse) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// GetAddr is used to request for addresses from other peers
type GetAddr struct {
	Sig []byte `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *GetAddr) Reset()                    { *m = GetAddr{} }
func (*GetAddr) ProtoMessage()               {}
func (*GetAddr) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{2} }

func (m *GetAddr) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// GetAddrResponse is used to send addresses in response to a GetAddr
type GetAddrResponse struct {
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	Sig       []byte   `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *GetAddrResponse) Reset()                    { *m = GetAddrResponse{} }
func (*GetAddrResponse) ProtoMessage()               {}
func (*GetAddrResponse) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{3} }

func (m *GetAddrResponse) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *GetAddrResponse) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// Ping represents a ping message
type Ping struct {
	Sig []byte `protobuf:"bytes,1,opt,name=Sig,proto3" json:"Sig,omitempty"`
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{4} }

func (m *Ping) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

// Pong represents a pong message
type Pong struct {
	Sig []byte `protobuf:"bytes,1,opt,name=Sig,proto3" json:"Sig,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{5} }

func (m *Pong) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func init() {
	proto.RegisterType((*Handshake)(nil), "wire.Handshake")
	proto.RegisterType((*HandshakeResponse)(nil), "wire.HandshakeResponse")
	proto.RegisterType((*GetAddr)(nil), "wire.GetAddr")
	proto.RegisterType((*GetAddrResponse)(nil), "wire.GetAddrResponse")
	proto.RegisterType((*Ping)(nil), "wire.Ping")
	proto.RegisterType((*Pong)(nil), "wire.Pong")
}
func (this *Handshake) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Handshake)
	if !ok {
		that2, ok := that.(Handshake)
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
	if this.Address != that1.Address {
		return false
	}
	if this.SubVersion != that1.SubVersion {
		return false
	}
	if !bytes.Equal(this.Sig, that1.Sig) {
		return false
	}
	return true
}
func (this *HandshakeResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HandshakeResponse)
	if !ok {
		that2, ok := that.(HandshakeResponse)
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
	if len(this.Addresses) != len(that1.Addresses) {
		return false
	}
	for i := range this.Addresses {
		if this.Addresses[i] != that1.Addresses[i] {
			return false
		}
	}
	if !bytes.Equal(this.Sig, that1.Sig) {
		return false
	}
	return true
}
func (this *GetAddr) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetAddr)
	if !ok {
		that2, ok := that.(GetAddr)
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
	if !bytes.Equal(this.Sig, that1.Sig) {
		return false
	}
	return true
}
func (this *GetAddrResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetAddrResponse)
	if !ok {
		that2, ok := that.(GetAddrResponse)
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
	if len(this.Addresses) != len(that1.Addresses) {
		return false
	}
	for i := range this.Addresses {
		if this.Addresses[i] != that1.Addresses[i] {
			return false
		}
	}
	if !bytes.Equal(this.Sig, that1.Sig) {
		return false
	}
	return true
}
func (this *Ping) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Ping)
	if !ok {
		that2, ok := that.(Ping)
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
	if !bytes.Equal(this.Sig, that1.Sig) {
		return false
	}
	return true
}
func (this *Pong) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Pong)
	if !ok {
		that2, ok := that.(Pong)
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
	if !bytes.Equal(this.Sig, that1.Sig) {
		return false
	}
	return true
}
func (this *Handshake) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&wire.Handshake{")
	s = append(s, "Address: "+fmt.Sprintf("%#v", this.Address)+",\n")
	s = append(s, "SubVersion: "+fmt.Sprintf("%#v", this.SubVersion)+",\n")
	s = append(s, "Sig: "+fmt.Sprintf("%#v", this.Sig)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HandshakeResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&wire.HandshakeResponse{")
	s = append(s, "Addresses: "+fmt.Sprintf("%#v", this.Addresses)+",\n")
	s = append(s, "Sig: "+fmt.Sprintf("%#v", this.Sig)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetAddr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&wire.GetAddr{")
	s = append(s, "Sig: "+fmt.Sprintf("%#v", this.Sig)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetAddrResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&wire.GetAddrResponse{")
	s = append(s, "Addresses: "+fmt.Sprintf("%#v", this.Addresses)+",\n")
	s = append(s, "Sig: "+fmt.Sprintf("%#v", this.Sig)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Ping) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&wire.Ping{")
	s = append(s, "Sig: "+fmt.Sprintf("%#v", this.Sig)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Pong) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&wire.Pong{")
	s = append(s, "Sig: "+fmt.Sprintf("%#v", this.Sig)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessages(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Handshake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Handshake) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.SubVersion) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.SubVersion)))
		i += copy(dAtA[i:], m.SubVersion)
	}
	if len(m.Sig) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Sig)))
		i += copy(dAtA[i:], m.Sig)
	}
	return i, nil
}

func (m *HandshakeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandshakeResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Sig) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Sig)))
		i += copy(dAtA[i:], m.Sig)
	}
	return i, nil
}

func (m *GetAddr) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAddr) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sig) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Sig)))
		i += copy(dAtA[i:], m.Sig)
	}
	return i, nil
}

func (m *GetAddrResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAddrResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Sig) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Sig)))
		i += copy(dAtA[i:], m.Sig)
	}
	return i, nil
}

func (m *Ping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ping) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sig) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Sig)))
		i += copy(dAtA[i:], m.Sig)
	}
	return i, nil
}

func (m *Pong) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pong) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sig) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Sig)))
		i += copy(dAtA[i:], m.Sig)
	}
	return i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Handshake) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.SubVersion)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *HandshakeResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *GetAddr) Size() (n int) {
	var l int
	_ = l
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *GetAddrResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *Ping) Size() (n int) {
	var l int
	_ = l
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *Pong) Size() (n int) {
	var l int
	_ = l
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func sovMessages(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Handshake) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Handshake{`,
		`Address:` + fmt.Sprintf("%v", this.Address) + `,`,
		`SubVersion:` + fmt.Sprintf("%v", this.SubVersion) + `,`,
		`Sig:` + fmt.Sprintf("%v", this.Sig) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HandshakeResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HandshakeResponse{`,
		`Addresses:` + fmt.Sprintf("%v", this.Addresses) + `,`,
		`Sig:` + fmt.Sprintf("%v", this.Sig) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetAddr) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetAddr{`,
		`Sig:` + fmt.Sprintf("%v", this.Sig) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetAddrResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetAddrResponse{`,
		`Addresses:` + fmt.Sprintf("%v", this.Addresses) + `,`,
		`Sig:` + fmt.Sprintf("%v", this.Sig) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Ping) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Ping{`,
		`Sig:` + fmt.Sprintf("%v", this.Sig) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Pong) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Pong{`,
		`Sig:` + fmt.Sprintf("%v", this.Sig) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessages(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Handshake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: Handshake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Handshake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
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
func (m *HandshakeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: HandshakeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandshakeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
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
func (m *GetAddr) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: GetAddr: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAddr: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
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
func (m *GetAddrResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: GetAddrResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAddrResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
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
func (m *Ping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: Ping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
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
func (m *Pong) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: Pong: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pong: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
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
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
				return 0, ErrInvalidLengthMessages
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessages
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
				next, err := skipMessages(dAtA[start:])
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
	ErrInvalidLengthMessages = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("wire/messages.proto", fileDescriptorMessages) }

var fileDescriptorMessages = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xcf, 0x2c, 0x4a,
	0xd5, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x01, 0x09, 0x2a, 0x85, 0x73, 0x71, 0x7a, 0x24, 0xe6, 0xa5, 0x14, 0x67, 0x24, 0x66, 0xa7,
	0x0a, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0xc1, 0xb8, 0x42, 0x72, 0x5c, 0x5c, 0xc5, 0xa5, 0x49, 0x61, 0xa9, 0x45, 0xc5, 0x99,
	0xf9, 0x79, 0x12, 0x4c, 0x60, 0x49, 0x24, 0x11, 0x21, 0x01, 0x2e, 0xe6, 0xe2, 0xcc, 0x74, 0x09,
	0x66, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x10, 0x53, 0xc9, 0x99, 0x4b, 0x10, 0x6e, 0x70, 0x50, 0x6a,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x0c, 0x17, 0x27, 0xd4, 0xc4, 0x54, 0x90, 0x15, 0xcc,
	0x1a, 0x9c, 0x41, 0x08, 0x01, 0x98, 0x21, 0x4c, 0x08, 0x43, 0xa4, 0xb9, 0xd8, 0xdd, 0x53, 0x4b,
	0x1c, 0x53, 0x52, 0x8a, 0x60, 0x92, 0x8c, 0x08, 0x49, 0x47, 0x2e, 0x7e, 0xa8, 0x24, 0xd9, 0xe6,
	0x4b, 0x70, 0xb1, 0x04, 0x64, 0xe6, 0xa5, 0x83, 0x64, 0x82, 0x11, 0x86, 0x07, 0x43, 0x65, 0xf2,
	0xb1, 0xc9, 0x38, 0xe9, 0x5c, 0x78, 0x28, 0xc7, 0x70, 0xe3, 0xa1, 0x1c, 0xc3, 0x87, 0x87, 0x72,
	0x8c, 0x0d, 0x8f, 0xe4, 0x18, 0x57, 0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x5f, 0x3c, 0x92, 0x63, 0xf8, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x24, 0x36, 0x70, 0x60, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x28,
	0x0b, 0xc5, 0x83, 0x01, 0x00, 0x00,
}
