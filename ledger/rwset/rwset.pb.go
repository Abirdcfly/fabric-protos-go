// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ledger/rwset/rwset.proto

package rwset

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

type TxReadWriteSet_DataModel int32

const (
	TxReadWriteSet_KV TxReadWriteSet_DataModel = 0
)

var TxReadWriteSet_DataModel_name = map[int32]string{
	0: "KV",
}

var TxReadWriteSet_DataModel_value = map[string]int32{
	"KV": 0,
}

func (x TxReadWriteSet_DataModel) String() string {
	return proto.EnumName(TxReadWriteSet_DataModel_name, int32(x))
}

func (TxReadWriteSet_DataModel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_794d00b812408f20, []int{0, 0}
}

// TxReadWriteSet encapsulates a read-write set for a transaction
// DataModel specifies the enum value of the data model
// ns_rwset field specifies a list of chaincode specific read-write set (one for each chaincode)
type TxReadWriteSet struct {
	DataModel            TxReadWriteSet_DataModel `protobuf:"varint,1,opt,name=data_model,json=dataModel,proto3,enum=rwset.TxReadWriteSet_DataModel" json:"data_model,omitempty"`
	NsRwset              []*NsReadWriteSet        `protobuf:"bytes,2,rep,name=ns_rwset,json=nsRwset,proto3" json:"ns_rwset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TxReadWriteSet) Reset()         { *m = TxReadWriteSet{} }
func (m *TxReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*TxReadWriteSet) ProtoMessage()    {}
func (*TxReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_794d00b812408f20, []int{0}
}

func (m *TxReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxReadWriteSet.Unmarshal(m, b)
}
func (m *TxReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxReadWriteSet.Marshal(b, m, deterministic)
}
func (m *TxReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxReadWriteSet.Merge(m, src)
}
func (m *TxReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_TxReadWriteSet.Size(m)
}
func (m *TxReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_TxReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_TxReadWriteSet proto.InternalMessageInfo

func (m *TxReadWriteSet) GetDataModel() TxReadWriteSet_DataModel {
	if m != nil {
		return m.DataModel
	}
	return TxReadWriteSet_KV
}

func (m *TxReadWriteSet) GetNsRwset() []*NsReadWriteSet {
	if m != nil {
		return m.NsRwset
	}
	return nil
}

// NsReadWriteSet encapsulates the read-write set for a chaincode
type NsReadWriteSet struct {
	Namespace             string                          `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Rwset                 []byte                          `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"`
	CollectionHashedRwset []*CollectionHashedReadWriteSet `protobuf:"bytes,3,rep,name=collection_hashed_rwset,json=collectionHashedRwset,proto3" json:"collection_hashed_rwset,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                        `json:"-"`
	XXX_unrecognized      []byte                          `json:"-"`
	XXX_sizecache         int32                           `json:"-"`
}

func (m *NsReadWriteSet) Reset()         { *m = NsReadWriteSet{} }
func (m *NsReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*NsReadWriteSet) ProtoMessage()    {}
func (*NsReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_794d00b812408f20, []int{1}
}

func (m *NsReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NsReadWriteSet.Unmarshal(m, b)
}
func (m *NsReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NsReadWriteSet.Marshal(b, m, deterministic)
}
func (m *NsReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NsReadWriteSet.Merge(m, src)
}
func (m *NsReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_NsReadWriteSet.Size(m)
}
func (m *NsReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_NsReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_NsReadWriteSet proto.InternalMessageInfo

func (m *NsReadWriteSet) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NsReadWriteSet) GetRwset() []byte {
	if m != nil {
		return m.Rwset
	}
	return nil
}

func (m *NsReadWriteSet) GetCollectionHashedRwset() []*CollectionHashedReadWriteSet {
	if m != nil {
		return m.CollectionHashedRwset
	}
	return nil
}

// CollectionHashedReadWriteSet encapsulate the hashed representation for the private read-write set for a collection
type CollectionHashedReadWriteSet struct {
	CollectionName       string   `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	HashedRwset          []byte   `protobuf:"bytes,2,opt,name=hashed_rwset,json=hashedRwset,proto3" json:"hashed_rwset,omitempty"`
	PvtRwsetHash         []byte   `protobuf:"bytes,3,opt,name=pvt_rwset_hash,json=pvtRwsetHash,proto3" json:"pvt_rwset_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionHashedReadWriteSet) Reset()         { *m = CollectionHashedReadWriteSet{} }
func (m *CollectionHashedReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*CollectionHashedReadWriteSet) ProtoMessage()    {}
func (*CollectionHashedReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_794d00b812408f20, []int{2}
}

func (m *CollectionHashedReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionHashedReadWriteSet.Unmarshal(m, b)
}
func (m *CollectionHashedReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionHashedReadWriteSet.Marshal(b, m, deterministic)
}
func (m *CollectionHashedReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionHashedReadWriteSet.Merge(m, src)
}
func (m *CollectionHashedReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_CollectionHashedReadWriteSet.Size(m)
}
func (m *CollectionHashedReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionHashedReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionHashedReadWriteSet proto.InternalMessageInfo

func (m *CollectionHashedReadWriteSet) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *CollectionHashedReadWriteSet) GetHashedRwset() []byte {
	if m != nil {
		return m.HashedRwset
	}
	return nil
}

func (m *CollectionHashedReadWriteSet) GetPvtRwsetHash() []byte {
	if m != nil {
		return m.PvtRwsetHash
	}
	return nil
}

// TxPvtReadWriteSet encapsulate the private read-write set for a transaction
type TxPvtReadWriteSet struct {
	DataModel            TxReadWriteSet_DataModel `protobuf:"varint,1,opt,name=data_model,json=dataModel,proto3,enum=rwset.TxReadWriteSet_DataModel" json:"data_model,omitempty"`
	NsPvtRwset           []*NsPvtReadWriteSet     `protobuf:"bytes,2,rep,name=ns_pvt_rwset,json=nsPvtRwset,proto3" json:"ns_pvt_rwset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TxPvtReadWriteSet) Reset()         { *m = TxPvtReadWriteSet{} }
func (m *TxPvtReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*TxPvtReadWriteSet) ProtoMessage()    {}
func (*TxPvtReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_794d00b812408f20, []int{3}
}

func (m *TxPvtReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxPvtReadWriteSet.Unmarshal(m, b)
}
func (m *TxPvtReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxPvtReadWriteSet.Marshal(b, m, deterministic)
}
func (m *TxPvtReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxPvtReadWriteSet.Merge(m, src)
}
func (m *TxPvtReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_TxPvtReadWriteSet.Size(m)
}
func (m *TxPvtReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_TxPvtReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_TxPvtReadWriteSet proto.InternalMessageInfo

func (m *TxPvtReadWriteSet) GetDataModel() TxReadWriteSet_DataModel {
	if m != nil {
		return m.DataModel
	}
	return TxReadWriteSet_KV
}

func (m *TxPvtReadWriteSet) GetNsPvtRwset() []*NsPvtReadWriteSet {
	if m != nil {
		return m.NsPvtRwset
	}
	return nil
}

// NsPvtReadWriteSet encapsulates the private read-write set for a chaincode
type NsPvtReadWriteSet struct {
	Namespace            string                       `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CollectionPvtRwset   []*CollectionPvtReadWriteSet `protobuf:"bytes,2,rep,name=collection_pvt_rwset,json=collectionPvtRwset,proto3" json:"collection_pvt_rwset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *NsPvtReadWriteSet) Reset()         { *m = NsPvtReadWriteSet{} }
func (m *NsPvtReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*NsPvtReadWriteSet) ProtoMessage()    {}
func (*NsPvtReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_794d00b812408f20, []int{4}
}

func (m *NsPvtReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NsPvtReadWriteSet.Unmarshal(m, b)
}
func (m *NsPvtReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NsPvtReadWriteSet.Marshal(b, m, deterministic)
}
func (m *NsPvtReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NsPvtReadWriteSet.Merge(m, src)
}
func (m *NsPvtReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_NsPvtReadWriteSet.Size(m)
}
func (m *NsPvtReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_NsPvtReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_NsPvtReadWriteSet proto.InternalMessageInfo

func (m *NsPvtReadWriteSet) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NsPvtReadWriteSet) GetCollectionPvtRwset() []*CollectionPvtReadWriteSet {
	if m != nil {
		return m.CollectionPvtRwset
	}
	return nil
}

// CollectionPvtReadWriteSet encapsulates the private read-write set for a collection
type CollectionPvtReadWriteSet struct {
	CollectionName       string   `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	Rwset                []byte   `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionPvtReadWriteSet) Reset()         { *m = CollectionPvtReadWriteSet{} }
func (m *CollectionPvtReadWriteSet) String() string { return proto.CompactTextString(m) }
func (*CollectionPvtReadWriteSet) ProtoMessage()    {}
func (*CollectionPvtReadWriteSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_794d00b812408f20, []int{5}
}

func (m *CollectionPvtReadWriteSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionPvtReadWriteSet.Unmarshal(m, b)
}
func (m *CollectionPvtReadWriteSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionPvtReadWriteSet.Marshal(b, m, deterministic)
}
func (m *CollectionPvtReadWriteSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionPvtReadWriteSet.Merge(m, src)
}
func (m *CollectionPvtReadWriteSet) XXX_Size() int {
	return xxx_messageInfo_CollectionPvtReadWriteSet.Size(m)
}
func (m *CollectionPvtReadWriteSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionPvtReadWriteSet.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionPvtReadWriteSet proto.InternalMessageInfo

func (m *CollectionPvtReadWriteSet) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *CollectionPvtReadWriteSet) GetRwset() []byte {
	if m != nil {
		return m.Rwset
	}
	return nil
}

/*
	func init() {
		proto.RegisterEnum("rwset.TxReadWriteSet_DataModel", TxReadWriteSet_DataModel_name, TxReadWriteSet_DataModel_value)
		proto.RegisterType((*TxReadWriteSet)(nil), "rwset.TxReadWriteSet")
		proto.RegisterType((*NsReadWriteSet)(nil), "rwset.NsReadWriteSet")
		proto.RegisterType((*CollectionHashedReadWriteSet)(nil), "rwset.CollectionHashedReadWriteSet")
		proto.RegisterType((*TxPvtReadWriteSet)(nil), "rwset.TxPvtReadWriteSet")
		proto.RegisterType((*NsPvtReadWriteSet)(nil), "rwset.NsPvtReadWriteSet")
		proto.RegisterType((*CollectionPvtReadWriteSet)(nil), "rwset.CollectionPvtReadWriteSet")
	}

func init() { proto.RegisterFile("ledger/rwset/rwset.proto", fileDescriptor_794d00b812408f20) }
*/
var fileDescriptor_794d00b812408f20 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0xef, 0x93, 0x30,
	0x18, 0xc6, 0xed, 0x77, 0xd9, 0x94, 0x77, 0x04, 0x5d, 0xdd, 0x22, 0x26, 0x4b, 0x9c, 0xd3, 0xc4,
	0xc5, 0x64, 0x60, 0xa6, 0x27, 0x0f, 0x1e, 0xd4, 0x83, 0x89, 0x71, 0x31, 0x75, 0xd1, 0x64, 0x1e,
	0x48, 0x81, 0x0a, 0x24, 0x40, 0x09, 0xad, 0x73, 0xfe, 0x01, 0x9e, 0xf5, 0xe6, 0xd9, 0xff, 0xd4,
	0xac, 0x65, 0x0c, 0x98, 0xbf, 0x0e, 0x5e, 0x08, 0x7d, 0xfb, 0x3c, 0x3c, 0x9f, 0xf6, 0xe5, 0x05,
	0x3b, 0x65, 0x61, 0xc4, 0x4a, 0xb7, 0xfc, 0x24, 0x98, 0xd4, 0x4f, 0xa7, 0x28, 0xb9, 0xe4, 0xb8,
	0xaf, 0x16, 0xf3, 0xef, 0x08, 0xac, 0xcd, 0x9e, 0x30, 0x1a, 0xbe, 0x2b, 0x13, 0xc9, 0xde, 0x30,
	0x89, 0x9f, 0x00, 0x84, 0x54, 0x52, 0x2f, 0xe3, 0x21, 0x4b, 0x6d, 0x34, 0x43, 0x0b, 0x6b, 0x75,
	0xcb, 0xd1, 0xde, 0xb6, 0xd4, 0x79, 0x4e, 0x25, 0x7d, 0x75, 0x90, 0x11, 0x23, 0x3c, 0xbe, 0xe2,
	0x07, 0x70, 0x25, 0x17, 0x9e, 0xd2, 0xdb, 0x17, 0xb3, 0xde, 0x62, 0xb8, 0x9a, 0x54, 0xee, 0xb5,
	0x68, 0xba, 0xc9, 0xe5, 0x5c, 0x10, 0x05, 0x71, 0x1d, 0x8c, 0xfa, 0x4b, 0x78, 0x00, 0x17, 0x2f,
	0xdf, 0x5e, 0xbb, 0x34, 0xff, 0x81, 0xc0, 0x6a, 0x1b, 0xf0, 0x14, 0x8c, 0x9c, 0x66, 0x4c, 0x14,
	0x34, 0x60, 0x0a, 0xcc, 0x20, 0xa7, 0x02, 0x1e, 0x43, 0xff, 0x18, 0x8a, 0x16, 0x26, 0xd1, 0x0b,
	0xfc, 0x1e, 0x6e, 0x04, 0x3c, 0x4d, 0x59, 0x20, 0x13, 0x9e, 0x7b, 0x31, 0x15, 0x31, 0x0b, 0x2b,
	0xb8, 0x9e, 0x82, 0xbb, 0x53, 0xc1, 0x3d, 0xab, 0x55, 0x2f, 0x94, 0xa8, 0x85, 0x3a, 0x09, 0xba,
	0xbb, 0x0a, 0xfc, 0x1b, 0x82, 0xe9, 0x9f, 0x7c, 0xf8, 0x1e, 0x5c, 0x6d, 0xa4, 0x1f, 0x58, 0x2b,
	0x6e, 0xeb, 0x54, 0x5e, 0xd3, 0x8c, 0xe1, 0xdb, 0x60, 0xb6, 0xd8, 0xf4, 0x19, 0x86, 0xf1, 0x29,
	0x0c, 0xdf, 0x05, 0xab, 0xd8, 0x49, 0xbd, 0xaf, 0x0e, 0x62, 0xf7, 0x94, 0xc8, 0x2c, 0x76, 0x52,
	0x29, 0x0e, 0xf9, 0xf3, 0xaf, 0x08, 0x46, 0x9b, 0xfd, 0xeb, 0x9d, 0xfc, 0xaf, 0x3d, 0x7d, 0x0c,
	0x66, 0x2e, 0xbc, 0x3a, 0xbe, 0xea, 0xab, 0x5d, 0xf7, 0xb5, 0x93, 0x47, 0x20, 0x57, 0x25, 0x75,
	0x49, 0x5f, 0x10, 0x8c, 0xce, 0x14, 0x7f, 0xe9, 0x25, 0x81, 0x71, 0xe3, 0xde, 0xba, 0xb9, 0xb3,
	0xb3, 0x96, 0x75, 0xf3, 0x71, 0xd0, 0xda, 0x52, 0x1c, 0x5b, 0xb8, 0xf9, 0x5b, 0xc3, 0xbf, 0x37,
	0xea, 0x97, 0x7f, 0xd9, 0x53, 0x1f, 0xee, 0xf3, 0x32, 0x72, 0xe2, 0xcf, 0x05, 0x2b, 0xf5, 0xc8,
	0x39, 0x1f, 0xa8, 0x5f, 0x26, 0x81, 0x9e, 0x36, 0xe1, 0x54, 0x45, 0xa5, 0xde, 0x3e, 0x8a, 0x12,
	0x19, 0x7f, 0xf4, 0x9d, 0x80, 0x67, 0x6e, 0xc3, 0xe2, 0x6a, 0xcb, 0x52, 0x5b, 0x96, 0x11, 0x77,
	0x9b, 0xd3, 0xeb, 0x0f, 0x54, 0xfd, 0xe1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x32, 0x8f,
	0x91, 0xd4, 0x03, 0x00, 0x00,
}
