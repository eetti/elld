// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wire/messages.proto

/*
	Package wire is a generated protocol buffer package.

	It is generated from these files:
		wire/messages.proto

	It has these top-level messages:
		Handshake
		HandshakeResponse
		ContractInvokeRequest
		ContractInvokeResponse
*/
package wire

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

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

// Handshake represents the first message between peers
type Handshake struct {
	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	SubVersion string `protobuf:"bytes,2,opt,name=subVersion,proto3" json:"subVersion,omitempty"`
	Sig        []byte `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *Handshake) Reset()                    { *m = Handshake{} }
func (m *Handshake) String() string            { return proto.CompactTextString(m) }
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
func (m *HandshakeResponse) String() string            { return proto.CompactTextString(m) }
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

// ContractInvokeRequest represents a request to invoke a smart contract
type ContractInvokeRequest struct {
	Function   string `protobuf:"bytes,1,opt,name=function,proto3" json:"function,omitempty"`
	ContractID string `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Data       []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ContractInvokeRequest) Reset()                    { *m = ContractInvokeRequest{} }
func (m *ContractInvokeRequest) String() string            { return proto.CompactTextString(m) }
func (*ContractInvokeRequest) ProtoMessage()               {}
func (*ContractInvokeRequest) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{2} }

func (m *ContractInvokeRequest) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *ContractInvokeRequest) GetContractID() string {
	if m != nil {
		return m.ContractID
	}
	return ""
}

func (m *ContractInvokeRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// ContractInvokeResponse is a response to the ContractService message
type ContractInvokeResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ContractInvokeResponse) Reset()                    { *m = ContractInvokeResponse{} }
func (m *ContractInvokeResponse) String() string            { return proto.CompactTextString(m) }
func (*ContractInvokeResponse) ProtoMessage()               {}
func (*ContractInvokeResponse) Descriptor() ([]byte, []int) { return fileDescriptorMessages, []int{3} }

func (m *ContractInvokeResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ContractInvokeResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Handshake)(nil), "wire.Handshake")
	proto.RegisterType((*HandshakeResponse)(nil), "wire.HandshakeResponse")
	proto.RegisterType((*ContractInvokeRequest)(nil), "wire.ContractInvokeRequest")
	proto.RegisterType((*ContractInvokeResponse)(nil), "wire.ContractInvokeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContractService service

type ContractServiceClient interface {
	ContractInvoke(ctx context.Context, in *ContractInvokeRequest, opts ...grpc.CallOption) (*ContractInvokeResponse, error)
}

type contractServiceClient struct {
	cc *grpc.ClientConn
}

func NewContractServiceClient(cc *grpc.ClientConn) ContractServiceClient {
	return &contractServiceClient{cc}
}

func (c *contractServiceClient) ContractInvoke(ctx context.Context, in *ContractInvokeRequest, opts ...grpc.CallOption) (*ContractInvokeResponse, error) {
	out := new(ContractInvokeResponse)
	err := grpc.Invoke(ctx, "/wire.ContractService/ContractInvoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContractService service

type ContractServiceServer interface {
	ContractInvoke(context.Context, *ContractInvokeRequest) (*ContractInvokeResponse, error)
}

func RegisterContractServiceServer(s *grpc.Server, srv ContractServiceServer) {
	s.RegisterService(&_ContractService_serviceDesc, srv)
}

func _ContractService_ContractInvoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContractInvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractServiceServer).ContractInvoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wire.ContractService/ContractInvoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractServiceServer).ContractInvoke(ctx, req.(*ContractInvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContractService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wire.ContractService",
	HandlerType: (*ContractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContractInvoke",
			Handler:    _ContractService_ContractInvoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wire/messages.proto",
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

func (m *ContractInvokeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractInvokeRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Function) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Function)))
		i += copy(dAtA[i:], m.Function)
	}
	if len(m.ContractID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.ContractID)))
		i += copy(dAtA[i:], m.ContractID)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *ContractInvokeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractInvokeResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Status)))
		i += copy(dAtA[i:], m.Status)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
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

func (m *ContractInvokeRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Function)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.ContractID)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *ContractInvokeResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Data)
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
func (m *ContractInvokeRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContractInvokeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractInvokeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Function", wireType)
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
			m.Function = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractID", wireType)
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
			m.ContractID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
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
func (m *ContractInvokeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContractInvokeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractInvokeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
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
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4e, 0xeb, 0x30,
	0x10, 0x86, 0xeb, 0xb6, 0xea, 0x7b, 0x19, 0x21, 0x28, 0x46, 0x54, 0x51, 0xa9, 0xa2, 0x2a, 0xab,
	0xac, 0x82, 0x04, 0x37, 0xa0, 0x5d, 0xc0, 0x82, 0x4d, 0x90, 0x60, 0x8b, 0x9b, 0x0c, 0x21, 0x42,
	0xd8, 0xc5, 0xe3, 0x94, 0xab, 0x70, 0x24, 0x96, 0x1c, 0x01, 0x85, 0x8b, 0x20, 0xa7, 0x71, 0x02,
	0xa8, 0xbb, 0x99, 0x7f, 0x94, 0x6f, 0xbe, 0x8c, 0xe1, 0xe8, 0xb5, 0xd0, 0x78, 0xfa, 0x8c, 0x44,
	0x22, 0x47, 0x8a, 0xd7, 0x5a, 0x19, 0xc5, 0x87, 0x36, 0x0c, 0xef, 0xc0, 0xbb, 0x14, 0x32, 0xa3,
	0x47, 0xf1, 0x84, 0xdc, 0x87, 0x7f, 0x22, 0xcb, 0x34, 0x12, 0xf9, 0x6c, 0xce, 0x22, 0x2f, 0x71,
	0x2d, 0x0f, 0x00, 0xa8, 0x5c, 0xdd, 0xa2, 0xa6, 0x42, 0x49, 0xbf, 0x5f, 0x0f, 0x7f, 0x24, 0x7c,
	0x0c, 0x03, 0x2a, 0x72, 0x7f, 0x30, 0x67, 0xd1, 0x5e, 0x62, 0xcb, 0x70, 0x01, 0x87, 0x2d, 0x38,
	0x41, 0x5a, 0x2b, 0x49, 0xc8, 0x67, 0xe0, 0x35, 0x44, 0xb4, 0x2b, 0x06, 0x91, 0x97, 0x74, 0x81,
	0x83, 0xf4, 0x3b, 0x48, 0x0e, 0xc7, 0x0b, 0x25, 0x8d, 0x16, 0xa9, 0xb9, 0x92, 0x1b, 0x65, 0x49,
	0x2f, 0x25, 0x92, 0xe1, 0x53, 0xf8, 0xff, 0x50, 0xca, 0xd4, 0x58, 0x9b, 0xad, 0x6a, 0xdb, 0x5b,
	0xd7, 0xd4, 0x7d, 0xb4, 0x74, 0xae, 0x5d, 0xc2, 0x39, 0x0c, 0x33, 0x61, 0x44, 0x23, 0x5b, 0xd7,
	0xe1, 0x12, 0x26, 0x7f, 0x17, 0x35, 0xca, 0x13, 0x18, 0x91, 0x11, 0xa6, 0x74, 0x27, 0x69, 0xba,
	0x96, 0xd2, 0xef, 0x28, 0x67, 0xf7, 0x70, 0xe0, 0x28, 0x37, 0xa8, 0x37, 0x45, 0x8a, 0xfc, 0x1a,
	0xf6, 0x7f, 0x83, 0xf9, 0x49, 0x6c, 0x0f, 0x1f, 0xef, 0xfc, 0xaf, 0xe9, 0x6c, 0xf7, 0x70, 0xeb,
	0x12, 0xf6, 0x2e, 0xc6, 0xef, 0x55, 0xc0, 0x3e, 0xaa, 0x80, 0x7d, 0x56, 0x01, 0x7b, 0xfb, 0x0a,
	0x7a, 0xab, 0x51, 0xfd, 0x9a, 0xe7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0xb7, 0x34, 0xa3,
	0xe4, 0x01, 0x00, 0x00,
}
