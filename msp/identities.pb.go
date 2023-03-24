// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msp/identities.proto

package msp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// This struct represents an Identity
// (with its MSP identifier) to be used
// to serialize it and deserialize it
type SerializedIdentity struct {
	// The identifier of the associated membership service provider
	Mspid string `protobuf:"bytes,1,opt,name=mspid,proto3" json:"mspid,omitempty"`
	// the Identity, serialized according to the rules of its MPS
	IdBytes              []byte   `protobuf:"bytes,2,opt,name=id_bytes,json=idBytes,proto3" json:"id_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SerializedIdentity) Reset()         { *m = SerializedIdentity{} }
func (m *SerializedIdentity) String() string { return proto.CompactTextString(m) }
func (*SerializedIdentity) ProtoMessage()    {}
func (*SerializedIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_872f7fc14bf2c238, []int{0}
}

func (m *SerializedIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SerializedIdentity.Unmarshal(m, b)
}
func (m *SerializedIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SerializedIdentity.Marshal(b, m, deterministic)
}
func (m *SerializedIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SerializedIdentity.Merge(m, src)
}
func (m *SerializedIdentity) XXX_Size() int {
	return xxx_messageInfo_SerializedIdentity.Size(m)
}
func (m *SerializedIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_SerializedIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_SerializedIdentity proto.InternalMessageInfo

func (m *SerializedIdentity) GetMspid() string {
	if m != nil {
		return m.Mspid
	}
	return ""
}

func (m *SerializedIdentity) GetIdBytes() []byte {
	if m != nil {
		return m.IdBytes
	}
	return nil
}

// This struct represents an Idemix Identity
// to be used to serialize it and deserialize it.
// The IdemixMSP will first serialize an idemix identity to bytes using
// this proto, and then uses these bytes as id_bytes in SerializedIdentity
type SerializedIdemixIdentity struct {
	// nym_x is the X-component of the pseudonym elliptic curve point.
	// It is a []byte representation of an amcl.BIG
	// The pseudonym can be seen as a public key of the identity, it is used to verify signatures.
	NymX []byte `protobuf:"bytes,1,opt,name=nym_x,json=nymX,proto3" json:"nym_x,omitempty"`
	// nym_y is the Y-component of the pseudonym elliptic curve point.
	// It is a []byte representation of an amcl.BIG
	// The pseudonym can be seen as a public key of the identity, it is used to verify signatures.
	NymY []byte `protobuf:"bytes,2,opt,name=nym_y,json=nymY,proto3" json:"nym_y,omitempty"`
	// ou contains the organizational unit of the idemix identity
	Ou []byte `protobuf:"bytes,3,opt,name=ou,proto3" json:"ou,omitempty"`
	// role contains the role of this identity (e.g., ADMIN or MEMBER)
	Role []byte `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	// proof contains the cryptographic evidence that this identity is valid
	Proof                []byte   `protobuf:"bytes,5,opt,name=proof,proto3" json:"proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SerializedIdemixIdentity) Reset()         { *m = SerializedIdemixIdentity{} }
func (m *SerializedIdemixIdentity) String() string { return proto.CompactTextString(m) }
func (*SerializedIdemixIdentity) ProtoMessage()    {}
func (*SerializedIdemixIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_872f7fc14bf2c238, []int{1}
}

func (m *SerializedIdemixIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SerializedIdemixIdentity.Unmarshal(m, b)
}
func (m *SerializedIdemixIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SerializedIdemixIdentity.Marshal(b, m, deterministic)
}
func (m *SerializedIdemixIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SerializedIdemixIdentity.Merge(m, src)
}
func (m *SerializedIdemixIdentity) XXX_Size() int {
	return xxx_messageInfo_SerializedIdemixIdentity.Size(m)
}
func (m *SerializedIdemixIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_SerializedIdemixIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_SerializedIdemixIdentity proto.InternalMessageInfo

func (m *SerializedIdemixIdentity) GetNymX() []byte {
	if m != nil {
		return m.NymX
	}
	return nil
}

func (m *SerializedIdemixIdentity) GetNymY() []byte {
	if m != nil {
		return m.NymY
	}
	return nil
}

func (m *SerializedIdemixIdentity) GetOu() []byte {
	if m != nil {
		return m.Ou
	}
	return nil
}

func (m *SerializedIdemixIdentity) GetRole() []byte {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *SerializedIdemixIdentity) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

/*
	func init() {
		proto.RegisterType((*SerializedIdentity)(nil), "msp.SerializedIdentity")
		proto.RegisterType((*SerializedIdemixIdentity)(nil), "msp.SerializedIdemixIdentity")
	}

func init() { proto.RegisterFile("msp/identities.proto", fileDescriptor_872f7fc14bf2c238) }
*/
var fileDescriptor_872f7fc14bf2c238 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0x47, 0x95, 0x34, 0xf9, 0xff, 0xc1, 0xaa, 0x18, 0x4c, 0x07, 0xb3, 0x95, 0x4e, 0x95, 0x50,
	0x9d, 0x81, 0x6f, 0x50, 0x89, 0x81, 0x0d, 0x85, 0x05, 0x58, 0xaa, 0xa6, 0xbe, 0xa6, 0x27, 0xe5,
	0x72, 0x96, 0xed, 0x48, 0x35, 0x03, 0x9f, 0x1d, 0x25, 0x06, 0x04, 0x9b, 0x7f, 0x4f, 0x4f, 0x4f,
	0x3e, 0xb1, 0x20, 0x6f, 0x2b, 0x34, 0xd0, 0x07, 0x0c, 0x08, 0x5e, 0x5b, 0xc7, 0x81, 0xe5, 0x8c,
	0xbc, 0x5d, 0x3d, 0x08, 0xf9, 0x0c, 0x0e, 0xf7, 0x1d, 0xbe, 0x83, 0x79, 0x4c, 0x4a, 0x94, 0x0b,
	0x51, 0x92, 0xb7, 0x68, 0x54, 0xb6, 0xcc, 0xd6, 0x97, 0x75, 0x1a, 0xf2, 0x46, 0x5c, 0xa0, 0xd9,
	0x35, 0x31, 0x80, 0x57, 0xf9, 0x32, 0x5b, 0xcf, 0xeb, 0xff, 0x68, 0xb6, 0xe3, 0x5c, 0x7d, 0x08,
	0xf5, 0x27, 0x43, 0x78, 0xfe, 0x89, 0x5d, 0x8b, 0xb2, 0x8f, 0xb4, 0x3b, 0x4f, 0xb1, 0x79, 0x5d,
	0xf4, 0x91, 0x5e, 0xbe, 0x61, 0xfc, 0x0a, 0x8d, 0xf0, 0x55, 0x5e, 0x89, 0x9c, 0x07, 0x35, 0x9b,
	0x48, 0xce, 0x83, 0x94, 0xa2, 0x70, 0xdc, 0x81, 0x2a, 0x92, 0x33, 0xbe, 0xc7, 0xaf, 0x59, 0xc7,
	0x7c, 0x54, 0xe5, 0x04, 0xd3, 0xd8, 0x3e, 0x89, 0x5b, 0x76, 0xad, 0x3e, 0x45, 0x0b, 0xae, 0x03,
	0xd3, 0x82, 0xd3, 0xc7, 0x7d, 0xe3, 0xf0, 0x90, 0x6e, 0xf5, 0x9a, 0xbc, 0x7d, 0xbb, 0x6b, 0x31,
	0x9c, 0x86, 0x46, 0x1f, 0x98, 0xaa, 0x5f, 0x66, 0x95, 0xcc, 0x4d, 0x32, 0x37, 0x2d, 0x57, 0xe4,
	0x6d, 0xf3, 0x6f, 0x9a, 0xf7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xf7, 0x41, 0xf7, 0x3c,
	0x01, 0x00, 0x00,
}
