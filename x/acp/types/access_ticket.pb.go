// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sourcehub/acp/access_ticket.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Represents a Capability token containing an opaque proof and a set of Operations
// the Actor is allowed to perform.
// Tickets should be verified by a Reference Monitor before granting access to the requested operations.
type AccessTicket struct {
	// identified the ticket version
	VersionDenominator string          `protobuf:"bytes,1,opt,name=version_denominator,json=versionDenominator,proto3" json:"version_denominator,omitempty"`
	DecisionId         string          `protobuf:"bytes,2,opt,name=decision_id,json=decisionId,proto3" json:"decision_id,omitempty"`
	Decision           *AccessDecision `protobuf:"bytes,3,opt,name=decision,proto3" json:"decision,omitempty"`
	// proof of existance that the given decision exists in the chain
	// validation strategy is dependent on ticket version
	DecisionProof []byte `protobuf:"bytes,4,opt,name=decision_proof,json=decisionProof,proto3" json:"decision_proof,omitempty"`
	// signature of ticket which must match actor pkey in the access decision
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *AccessTicket) Reset()         { *m = AccessTicket{} }
func (m *AccessTicket) String() string { return proto.CompactTextString(m) }
func (*AccessTicket) ProtoMessage()    {}
func (*AccessTicket) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd5443967ce9352d, []int{0}
}
func (m *AccessTicket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessTicket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessTicket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessTicket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTicket.Merge(m, src)
}
func (m *AccessTicket) XXX_Size() int {
	return m.Size()
}
func (m *AccessTicket) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTicket.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTicket proto.InternalMessageInfo

func (m *AccessTicket) GetVersionDenominator() string {
	if m != nil {
		return m.VersionDenominator
	}
	return ""
}

func (m *AccessTicket) GetDecisionId() string {
	if m != nil {
		return m.DecisionId
	}
	return ""
}

func (m *AccessTicket) GetDecision() *AccessDecision {
	if m != nil {
		return m.Decision
	}
	return nil
}

func (m *AccessTicket) GetDecisionProof() []byte {
	if m != nil {
		return m.DecisionProof
	}
	return nil
}

func (m *AccessTicket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*AccessTicket)(nil), "sourcehub.acp.AccessTicket")
}

func init() { proto.RegisterFile("sourcehub/acp/access_ticket.proto", fileDescriptor_bd5443967ce9352d) }

var fileDescriptor_bd5443967ce9352d = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xeb, 0xef, 0x03, 0x44, 0xdd, 0x96, 0xc1, 0x30, 0x44, 0x15, 0xb8, 0x05, 0x84, 0xd4,
	0x29, 0x96, 0x60, 0x62, 0x04, 0x75, 0x80, 0x0d, 0x55, 0x4c, 0x2c, 0x95, 0xeb, 0xb8, 0xc6, 0x2a,
	0xc9, 0xb5, 0x6c, 0x87, 0x3f, 0x6f, 0xc1, 0x63, 0x31, 0x76, 0x64, 0xac, 0xda, 0x17, 0x41, 0x71,
	0x9a, 0x44, 0x48, 0x6c, 0xbe, 0xe7, 0xfc, 0x74, 0x7d, 0xee, 0xc1, 0xa7, 0x0e, 0x72, 0x2b, 0xe4,
	0x73, 0x3e, 0x63, 0x5c, 0x18, 0xc6, 0x85, 0x90, 0xce, 0x4d, 0xbd, 0x16, 0x0b, 0xe9, 0x63, 0x63,
	0xc1, 0x03, 0xe9, 0xd5, 0x48, 0xcc, 0x85, 0xe9, 0x1f, 0x29, 0x50, 0x10, 0x1c, 0x56, 0xbc, 0x4a,
	0xa8, 0x3f, 0x50, 0x00, 0xea, 0x45, 0xb2, 0x30, 0xcd, 0xf2, 0x39, 0xf3, 0x3a, 0x95, 0xce, 0xf3,
	0xd4, 0x6c, 0x81, 0xf3, 0x3f, 0x3f, 0x4a, 0xa4, 0xd0, 0x4e, 0x43, 0x56, 0x42, 0x67, 0x2b, 0x84,
	0xbb, 0x37, 0xc1, 0x79, 0x0c, 0x09, 0x08, 0xc3, 0x87, 0xaf, 0xd2, 0x16, 0xc4, 0x34, 0x91, 0x19,
	0xa4, 0x3a, 0xe3, 0x1e, 0x6c, 0x84, 0x86, 0x68, 0xd4, 0x9e, 0x90, 0xad, 0x35, 0x6e, 0x1c, 0x32,
	0xc0, 0x9d, 0x6a, 0xe7, 0x54, 0x27, 0xd1, 0xbf, 0x00, 0xe2, 0x4a, 0xba, 0x4f, 0xc8, 0x35, 0xde,
	0xaf, 0xa6, 0xe8, 0xff, 0x10, 0x8d, 0x3a, 0x97, 0x27, 0xf1, 0xaf, 0x03, 0xe3, 0x32, 0xc0, 0x78,
	0x0b, 0x4d, 0x6a, 0x9c, 0x5c, 0xe0, 0x83, 0x7a, 0xb7, 0xb1, 0x00, 0xf3, 0x68, 0x67, 0x88, 0x46,
	0xdd, 0x49, 0xaf, 0x52, 0x1f, 0x0a, 0x91, 0x1c, 0xe3, 0xb6, 0xd3, 0x2a, 0xe3, 0x3e, 0xb7, 0x32,
	0xda, 0x0d, 0x44, 0x23, 0xdc, 0xde, 0x7d, 0xad, 0x29, 0x5a, 0xae, 0x29, 0x5a, 0xad, 0x29, 0xfa,
	0xdc, 0xd0, 0xd6, 0x72, 0x43, 0x5b, 0xdf, 0x1b, 0xda, 0x7a, 0x8a, 0x95, 0xf6, 0x45, 0x06, 0x01,
	0x29, 0x2b, 0x13, 0x65, 0xd2, 0xbf, 0x81, 0x5d, 0xb0, 0xa6, 0xba, 0xf7, 0x50, 0x9e, 0xff, 0x30,
	0xd2, 0xcd, 0xf6, 0x42, 0x67, 0x57, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xf0, 0x38, 0xfc,
	0xc3, 0x01, 0x00, 0x00,
}

func (m *AccessTicket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessTicket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessTicket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintAccessTicket(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DecisionProof) > 0 {
		i -= len(m.DecisionProof)
		copy(dAtA[i:], m.DecisionProof)
		i = encodeVarintAccessTicket(dAtA, i, uint64(len(m.DecisionProof)))
		i--
		dAtA[i] = 0x22
	}
	if m.Decision != nil {
		{
			size, err := m.Decision.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAccessTicket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DecisionId) > 0 {
		i -= len(m.DecisionId)
		copy(dAtA[i:], m.DecisionId)
		i = encodeVarintAccessTicket(dAtA, i, uint64(len(m.DecisionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.VersionDenominator) > 0 {
		i -= len(m.VersionDenominator)
		copy(dAtA[i:], m.VersionDenominator)
		i = encodeVarintAccessTicket(dAtA, i, uint64(len(m.VersionDenominator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccessTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccessTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccessTicket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VersionDenominator)
	if l > 0 {
		n += 1 + l + sovAccessTicket(uint64(l))
	}
	l = len(m.DecisionId)
	if l > 0 {
		n += 1 + l + sovAccessTicket(uint64(l))
	}
	if m.Decision != nil {
		l = m.Decision.Size()
		n += 1 + l + sovAccessTicket(uint64(l))
	}
	l = len(m.DecisionProof)
	if l > 0 {
		n += 1 + l + sovAccessTicket(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovAccessTicket(uint64(l))
	}
	return n
}

func sovAccessTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccessTicket(x uint64) (n int) {
	return sovAccessTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessTicket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccessTicket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AccessTicket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessTicket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionDenominator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccessTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VersionDenominator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DecisionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccessTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DecisionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decision", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAccessTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Decision == nil {
				m.Decision = &AccessDecision{}
			}
			if err := m.Decision.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DecisionProof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAccessTicket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DecisionProof = append(m.DecisionProof[:0], dAtA[iNdEx:postIndex]...)
			if m.DecisionProof == nil {
				m.DecisionProof = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccessTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAccessTicket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAccessTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccessTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccessTicket
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
func skipAccessTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccessTicket
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
					return 0, ErrIntOverflowAccessTicket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccessTicket
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
			if length < 0 {
				return 0, ErrInvalidLengthAccessTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccessTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccessTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccessTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccessTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccessTicket = fmt.Errorf("proto: unexpected end of group")
)