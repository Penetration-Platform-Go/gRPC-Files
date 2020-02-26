// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mongodb.proto

package mongodb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Views struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Views) Reset()         { *m = Views{} }
func (m *Views) String() string { return proto.CompactTextString(m) }
func (*Views) ProtoMessage()    {}
func (*Views) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{1}
}

func (m *Views) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Views.Unmarshal(m, b)
}
func (m *Views) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Views.Marshal(b, m, deterministic)
}
func (m *Views) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Views.Merge(m, src)
}
func (m *Views) XXX_Size() int {
	return xxx_messageInfo_Views.Size(m)
}
func (m *Views) XXX_DiscardUnknown() {
	xxx_messageInfo_Views.DiscardUnknown(m)
}

var xxx_messageInfo_Views proto.InternalMessageInfo

func (m *Views) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Views) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Equipment struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ip                   []string `protobuf:"bytes,2,rep,name=ip,proto3" json:"ip,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Equipment) Reset()         { *m = Equipment{} }
func (m *Equipment) String() string { return proto.CompactTextString(m) }
func (*Equipment) ProtoMessage()    {}
func (*Equipment) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{2}
}

func (m *Equipment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Equipment.Unmarshal(m, b)
}
func (m *Equipment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Equipment.Marshal(b, m, deterministic)
}
func (m *Equipment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Equipment.Merge(m, src)
}
func (m *Equipment) XXX_Size() int {
	return xxx_messageInfo_Equipment.Size(m)
}
func (m *Equipment) XXX_DiscardUnknown() {
	xxx_messageInfo_Equipment.DiscardUnknown(m)
}

var xxx_messageInfo_Equipment proto.InternalMessageInfo

func (m *Equipment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Equipment) GetIp() []string {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *Equipment) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Column struct {
	Connected            []bool   `protobuf:"varint,1,rep,packed,name=connected,proto3" json:"connected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{3}
}

func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (m *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(m, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetConnected() []bool {
	if m != nil {
		return m.Connected
	}
	return nil
}

type Map struct {
	Column               []*Column `protobuf:"bytes,1,rep,name=column,proto3" json:"column,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Map) Reset()         { *m = Map{} }
func (m *Map) String() string { return proto.CompactTextString(m) }
func (*Map) ProtoMessage()    {}
func (*Map) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{4}
}

func (m *Map) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Map.Unmarshal(m, b)
}
func (m *Map) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Map.Marshal(b, m, deterministic)
}
func (m *Map) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Map.Merge(m, src)
}
func (m *Map) XXX_Size() int {
	return xxx_messageInfo_Map.Size(m)
}
func (m *Map) XXX_DiscardUnknown() {
	xxx_messageInfo_Map.DiscardUnknown(m)
}

var xxx_messageInfo_Map proto.InternalMessageInfo

func (m *Map) GetColumn() []*Column {
	if m != nil {
		return m.Column
	}
	return nil
}

type ProjectInformation struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 string       `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Title                string       `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Score                int32        `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	Equipment            []*Equipment `protobuf:"bytes,5,rep,name=equipment,proto3" json:"equipment,omitempty"`
	Map                  *Map         `protobuf:"bytes,6,opt,name=map,proto3" json:"map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProjectInformation) Reset()         { *m = ProjectInformation{} }
func (m *ProjectInformation) String() string { return proto.CompactTextString(m) }
func (*ProjectInformation) ProtoMessage()    {}
func (*ProjectInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{5}
}

func (m *ProjectInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectInformation.Unmarshal(m, b)
}
func (m *ProjectInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectInformation.Marshal(b, m, deterministic)
}
func (m *ProjectInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectInformation.Merge(m, src)
}
func (m *ProjectInformation) XXX_Size() int {
	return xxx_messageInfo_ProjectInformation.Size(m)
}
func (m *ProjectInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectInformation.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectInformation proto.InternalMessageInfo

func (m *ProjectInformation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProjectInformation) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ProjectInformation) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ProjectInformation) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *ProjectInformation) GetEquipment() []*Equipment {
	if m != nil {
		return m.Equipment
	}
	return nil
}

func (m *ProjectInformation) GetMap() *Map {
	if m != nil {
		return m.Map
	}
	return nil
}

type Result struct {
	IsVaild              bool     `protobuf:"varint,1,opt,name=isVaild,proto3" json:"isVaild,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{6}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetIsVaild() bool {
	if m != nil {
		return m.IsVaild
	}
	return false
}

func (m *Result) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Value struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{7}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Value) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Condition struct {
	Value                []*Value `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{8}
}

func (m *Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Condition.Unmarshal(m, b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
}
func (m *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(m, src)
}
func (m *Condition) XXX_Size() int {
	return xxx_messageInfo_Condition.Size(m)
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetValue() []*Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type UpdateMessage struct {
	Condition            *Condition          `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	Key                  []string            `protobuf:"bytes,2,rep,name=key,proto3" json:"key,omitempty"`
	Value                *ProjectInformation `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpdateMessage) Reset()         { *m = UpdateMessage{} }
func (m *UpdateMessage) String() string { return proto.CompactTextString(m) }
func (*UpdateMessage) ProtoMessage()    {}
func (*UpdateMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{9}
}

func (m *UpdateMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMessage.Unmarshal(m, b)
}
func (m *UpdateMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMessage.Marshal(b, m, deterministic)
}
func (m *UpdateMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMessage.Merge(m, src)
}
func (m *UpdateMessage) XXX_Size() int {
	return xxx_messageInfo_UpdateMessage.Size(m)
}
func (m *UpdateMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMessage proto.InternalMessageInfo

func (m *UpdateMessage) GetCondition() *Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *UpdateMessage) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *UpdateMessage) GetValue() *ProjectInformation {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "mongodb.Empty")
	proto.RegisterType((*Views)(nil), "mongodb.Views")
	proto.RegisterType((*Equipment)(nil), "mongodb.Equipment")
	proto.RegisterType((*Column)(nil), "mongodb.Column")
	proto.RegisterType((*Map)(nil), "mongodb.Map")
	proto.RegisterType((*ProjectInformation)(nil), "mongodb.ProjectInformation")
	proto.RegisterType((*Result)(nil), "mongodb.Result")
	proto.RegisterType((*Value)(nil), "mongodb.Value")
	proto.RegisterType((*Condition)(nil), "mongodb.Condition")
	proto.RegisterType((*UpdateMessage)(nil), "mongodb.UpdateMessage")
}

func init() { proto.RegisterFile("mongodb.proto", fileDescriptor_21010b45f990def2) }

var fileDescriptor_21010b45f990def2 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x55, 0x92, 0xee, 0x6e, 0x33, 0x25, 0x05, 0x59, 0x08, 0x45, 0x0b, 0x42, 0x51, 0x84, 0x20,
	0xe2, 0xb0, 0x6c, 0x17, 0x0e, 0xbd, 0x70, 0xa0, 0xdb, 0x1e, 0x7a, 0x58, 0x09, 0x2c, 0xd1, 0x7b,
	0x9a, 0x0c, 0x95, 0x21, 0xb1, 0x4d, 0xe2, 0x80, 0xf6, 0x07, 0xf8, 0x24, 0xbe, 0x86, 0x8f, 0x41,
	0xb6, 0xe3, 0xa4, 0x65, 0x57, 0xe5, 0x36, 0x93, 0x79, 0x6f, 0xfc, 0xe6, 0x8d, 0x1d, 0x88, 0x6a,
	0xc1, 0x6f, 0x44, 0x79, 0xbd, 0x90, 0x8d, 0x50, 0x82, 0xcc, 0xfa, 0x34, 0x9d, 0xc1, 0xe4, 0xa2,
	0x96, 0x6a, 0x9b, 0xbe, 0x81, 0xc9, 0x15, 0xc3, 0x9f, 0x2d, 0x79, 0x04, 0xc1, 0x37, 0xdc, 0xc6,
	0x5e, 0xe2, 0x65, 0x21, 0xd5, 0x21, 0x79, 0x0c, 0x93, 0x1f, 0x79, 0xd5, 0x61, 0xec, 0x27, 0x5e,
	0x36, 0xa1, 0x36, 0x49, 0xd7, 0x10, 0x5e, 0x7c, 0xef, 0x98, 0xac, 0x91, 0x2b, 0x42, 0xe0, 0x80,
	0xe7, 0x35, 0xf6, 0x2c, 0x13, 0x93, 0x63, 0xf0, 0x99, 0x8c, 0xfd, 0x24, 0xc8, 0x42, 0xea, 0x33,
	0xa9, 0x31, 0x6a, 0x2b, 0x31, 0x0e, 0x2c, 0x46, 0xc7, 0xe9, 0x4b, 0x98, 0xae, 0x45, 0xd5, 0xd5,
	0x9c, 0x3c, 0x83, 0xb0, 0x10, 0x9c, 0x63, 0xa1, 0xb0, 0x8c, 0xbd, 0x24, 0xc8, 0x0e, 0xe9, 0xf8,
	0x21, 0x5d, 0x40, 0xb0, 0xc9, 0x25, 0x79, 0x05, 0xd3, 0xc2, 0xc0, 0x0d, 0xe2, 0x68, 0xf5, 0x70,
	0xe1, 0xc6, 0xb2, 0x5d, 0x68, 0x5f, 0x4e, 0x7f, 0x7b, 0x40, 0x3e, 0x36, 0xe2, 0x2b, 0x16, 0xea,
	0x92, 0x7f, 0x11, 0x4d, 0x9d, 0x2b, 0x26, 0xb8, 0x91, 0x54, 0xf6, 0x22, 0x7d, 0x56, 0x6a, 0x49,
	0x5d, 0x8b, 0x8d, 0x19, 0x2c, 0xa4, 0x26, 0xd6, 0xd3, 0x2a, 0xa6, 0x2a, 0xa7, 0xd3, 0x26, 0xfa,
	0x6b, 0x5b, 0x88, 0x06, 0xe3, 0x03, 0xeb, 0x81, 0x49, 0xc8, 0x12, 0x42, 0x74, 0x1e, 0xc4, 0x13,
	0x23, 0x89, 0x0c, 0x92, 0x06, 0x77, 0xe8, 0x08, 0x22, 0xcf, 0x21, 0xa8, 0x73, 0x19, 0x4f, 0x13,
	0x2f, 0x3b, 0x5a, 0x3d, 0x18, 0xb0, 0x9b, 0x5c, 0x52, 0x5d, 0x48, 0x4f, 0x61, 0x4a, 0xb1, 0xed,
	0x2a, 0x45, 0x62, 0x98, 0xb1, 0xf6, 0x2a, 0x67, 0x95, 0x15, 0x7c, 0x48, 0x5d, 0x7a, 0x77, 0x1f,
	0xa1, 0xdb, 0x87, 0x5e, 0xa0, 0x0e, 0xfe, 0xb7, 0xc0, 0x81, 0x70, 0x02, 0xe1, 0x5a, 0xf0, 0x92,
	0x19, 0x67, 0x5e, 0x38, 0x88, 0x35, 0xf6, 0x78, 0x50, 0x66, 0x7a, 0x3a, 0xca, 0x2f, 0x0f, 0xa2,
	0xcf, 0xb2, 0xcc, 0x15, 0x6e, 0xb0, 0x6d, 0xf3, 0x1b, 0xe3, 0x40, 0xe1, 0x9a, 0x98, 0x23, 0x6f,
	0x3b, 0x30, 0xb4, 0xa7, 0x23, 0xc8, 0xc9, 0xb3, 0xf7, 0xc2, 0xc8, 0x3b, 0x71, 0x67, 0x07, 0x86,
	0xff, 0x74, 0xe0, 0xef, 0x6e, 0xb0, 0x17, 0xb2, 0xfa, 0xe3, 0xc3, 0x6c, 0xa3, 0x51, 0xe7, 0x67,
	0xe4, 0x0c, 0xa2, 0x4f, 0x1d, 0x36, 0xdb, 0x1e, 0xdd, 0x92, 0x3d, 0x02, 0xe6, 0xf7, 0x35, 0x5d,
	0x7a, 0xe4, 0x3d, 0x44, 0x97, 0xbc, 0xc5, 0x46, 0xf5, 0x55, 0x72, 0x1f, 0x7e, 0x3e, 0x5e, 0xbb,
	0x7e, 0x57, 0xa7, 0xce, 0x16, 0x47, 0x7f, 0x32, 0x20, 0xee, 0xd8, 0xb5, 0xcb, 0x7c, 0x07, 0xd1,
	0x39, 0x56, 0x38, 0x32, 0xf7, 0x89, 0xdf, 0x61, 0xad, 0x00, 0xcc, 0xc8, 0xf6, 0xc5, 0xee, 0xa3,
	0xdc, 0x5a, 0xa0, 0xc6, 0x2c, 0x3d, 0xf2, 0x1a, 0x0e, 0x3f, 0x94, 0xa5, 0x65, 0x8c, 0x55, 0xf3,
	0xf8, 0xe7, 0xff, 0xe4, 0xd7, 0x53, 0xf3, 0x97, 0x78, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xeb,
	0x1d, 0x0a, 0xe2, 0x36, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MongoDBClient is the client API for MongoDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MongoDBClient interface {
	QueryProjects(ctx context.Context, in *Condition, opts ...grpc.CallOption) (MongoDB_QueryProjectsClient, error)
	InsertProject(ctx context.Context, in *ProjectInformation, opts ...grpc.CallOption) (*Result, error)
	UpdateProject(ctx context.Context, in *UpdateMessage, opts ...grpc.CallOption) (*Result, error)
	DeleteProject(ctx context.Context, in *Condition, opts ...grpc.CallOption) (*Result, error)
	QueryViews(ctx context.Context, in *Condition, opts ...grpc.CallOption) (MongoDB_QueryViewsClient, error)
	AddViews(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type mongoDBClient struct {
	cc *grpc.ClientConn
}

func NewMongoDBClient(cc *grpc.ClientConn) MongoDBClient {
	return &mongoDBClient{cc}
}

func (c *mongoDBClient) QueryProjects(ctx context.Context, in *Condition, opts ...grpc.CallOption) (MongoDB_QueryProjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MongoDB_serviceDesc.Streams[0], "/mongodb.MongoDB/QueryProjects", opts...)
	if err != nil {
		return nil, err
	}
	x := &mongoDBQueryProjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MongoDB_QueryProjectsClient interface {
	Recv() (*ProjectInformation, error)
	grpc.ClientStream
}

type mongoDBQueryProjectsClient struct {
	grpc.ClientStream
}

func (x *mongoDBQueryProjectsClient) Recv() (*ProjectInformation, error) {
	m := new(ProjectInformation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mongoDBClient) InsertProject(ctx context.Context, in *ProjectInformation, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/mongodb.MongoDB/InsertProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoDBClient) UpdateProject(ctx context.Context, in *UpdateMessage, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/mongodb.MongoDB/UpdateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoDBClient) DeleteProject(ctx context.Context, in *Condition, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/mongodb.MongoDB/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoDBClient) QueryViews(ctx context.Context, in *Condition, opts ...grpc.CallOption) (MongoDB_QueryViewsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MongoDB_serviceDesc.Streams[1], "/mongodb.MongoDB/QueryViews", opts...)
	if err != nil {
		return nil, err
	}
	x := &mongoDBQueryViewsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MongoDB_QueryViewsClient interface {
	Recv() (*Views, error)
	grpc.ClientStream
}

type mongoDBQueryViewsClient struct {
	grpc.ClientStream
}

func (x *mongoDBQueryViewsClient) Recv() (*Views, error) {
	m := new(Views)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mongoDBClient) AddViews(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/mongodb.MongoDB/AddViews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongoDBServer is the server API for MongoDB service.
type MongoDBServer interface {
	QueryProjects(*Condition, MongoDB_QueryProjectsServer) error
	InsertProject(context.Context, *ProjectInformation) (*Result, error)
	UpdateProject(context.Context, *UpdateMessage) (*Result, error)
	DeleteProject(context.Context, *Condition) (*Result, error)
	QueryViews(*Condition, MongoDB_QueryViewsServer) error
	AddViews(context.Context, *Empty) (*Empty, error)
}

// UnimplementedMongoDBServer can be embedded to have forward compatible implementations.
type UnimplementedMongoDBServer struct {
}

func (*UnimplementedMongoDBServer) QueryProjects(req *Condition, srv MongoDB_QueryProjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryProjects not implemented")
}
func (*UnimplementedMongoDBServer) InsertProject(ctx context.Context, req *ProjectInformation) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertProject not implemented")
}
func (*UnimplementedMongoDBServer) UpdateProject(ctx context.Context, req *UpdateMessage) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (*UnimplementedMongoDBServer) DeleteProject(ctx context.Context, req *Condition) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (*UnimplementedMongoDBServer) QueryViews(req *Condition, srv MongoDB_QueryViewsServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryViews not implemented")
}
func (*UnimplementedMongoDBServer) AddViews(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddViews not implemented")
}

func RegisterMongoDBServer(s *grpc.Server, srv MongoDBServer) {
	s.RegisterService(&_MongoDB_serviceDesc, srv)
}

func _MongoDB_QueryProjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Condition)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MongoDBServer).QueryProjects(m, &mongoDBQueryProjectsServer{stream})
}

type MongoDB_QueryProjectsServer interface {
	Send(*ProjectInformation) error
	grpc.ServerStream
}

type mongoDBQueryProjectsServer struct {
	grpc.ServerStream
}

func (x *mongoDBQueryProjectsServer) Send(m *ProjectInformation) error {
	return x.ServerStream.SendMsg(m)
}

func _MongoDB_InsertProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoDBServer).InsertProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongodb.MongoDB/InsertProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoDBServer).InsertProject(ctx, req.(*ProjectInformation))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoDB_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoDBServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongodb.MongoDB/UpdateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoDBServer).UpdateProject(ctx, req.(*UpdateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoDB_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoDBServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongodb.MongoDB/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoDBServer).DeleteProject(ctx, req.(*Condition))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoDB_QueryViews_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Condition)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MongoDBServer).QueryViews(m, &mongoDBQueryViewsServer{stream})
}

type MongoDB_QueryViewsServer interface {
	Send(*Views) error
	grpc.ServerStream
}

type mongoDBQueryViewsServer struct {
	grpc.ServerStream
}

func (x *mongoDBQueryViewsServer) Send(m *Views) error {
	return x.ServerStream.SendMsg(m)
}

func _MongoDB_AddViews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoDBServer).AddViews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongodb.MongoDB/AddViews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoDBServer).AddViews(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MongoDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mongodb.MongoDB",
	HandlerType: (*MongoDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertProject",
			Handler:    _MongoDB_InsertProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _MongoDB_UpdateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _MongoDB_DeleteProject_Handler,
		},
		{
			MethodName: "AddViews",
			Handler:    _MongoDB_AddViews_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryProjects",
			Handler:       _MongoDB_QueryProjects_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "QueryViews",
			Handler:       _MongoDB_QueryViews_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mongodb.proto",
}
