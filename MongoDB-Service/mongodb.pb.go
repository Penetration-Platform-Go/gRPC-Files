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

type Ip struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ip) Reset()         { *m = Ip{} }
func (m *Ip) String() string { return proto.CompactTextString(m) }
func (*Ip) ProtoMessage()    {}
func (*Ip) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{0}
}

func (m *Ip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ip.Unmarshal(m, b)
}
func (m *Ip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ip.Marshal(b, m, deterministic)
}
func (m *Ip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ip.Merge(m, src)
}
func (m *Ip) XXX_Size() int {
	return xxx_messageInfo_Ip.Size(m)
}
func (m *Ip) XXX_DiscardUnknown() {
	xxx_messageInfo_Ip.DiscardUnknown(m)
}

var xxx_messageInfo_Ip proto.InternalMessageInfo

func (m *Ip) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Ip) GetValue() string {
	if m != nil {
		return m.Value
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
	return fileDescriptor_21010b45f990def2, []int{1}
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
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Column               []*Column `protobuf:"bytes,2,rep,name=column,proto3" json:"column,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Map) Reset()         { *m = Map{} }
func (m *Map) String() string { return proto.CompactTextString(m) }
func (*Map) ProtoMessage()    {}
func (*Map) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{2}
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

func (m *Map) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Map) GetColumn() []*Column {
	if m != nil {
		return m.Column
	}
	return nil
}

type ProjectInformation struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Score                int32    `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	Ip                   []*Ip    `protobuf:"bytes,4,rep,name=ip,proto3" json:"ip,omitempty"`
	Map                  *Map     `protobuf:"bytes,5,opt,name=map,proto3" json:"map,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectInformation) Reset()         { *m = ProjectInformation{} }
func (m *ProjectInformation) String() string { return proto.CompactTextString(m) }
func (*ProjectInformation) ProtoMessage()    {}
func (*ProjectInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_21010b45f990def2, []int{3}
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

func (m *ProjectInformation) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *ProjectInformation) GetIp() []*Ip {
	if m != nil {
		return m.Ip
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
	return fileDescriptor_21010b45f990def2, []int{4}
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
	return fileDescriptor_21010b45f990def2, []int{5}
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
	return fileDescriptor_21010b45f990def2, []int{6}
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

func init() {
	proto.RegisterType((*Ip)(nil), "mongodb.Ip")
	proto.RegisterType((*Column)(nil), "mongodb.Column")
	proto.RegisterType((*Map)(nil), "mongodb.Map")
	proto.RegisterType((*ProjectInformation)(nil), "mongodb.ProjectInformation")
	proto.RegisterType((*Result)(nil), "mongodb.Result")
	proto.RegisterType((*Value)(nil), "mongodb.Value")
	proto.RegisterType((*Condition)(nil), "mongodb.Condition")
}

func init() { proto.RegisterFile("mongodb.proto", fileDescriptor_21010b45f990def2) }

var fileDescriptor_21010b45f990def2 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xbf, 0x4f, 0xe3, 0x30,
	0x18, 0x55, 0x9c, 0x26, 0x6d, 0xbf, 0x5e, 0x7a, 0x27, 0xeb, 0x86, 0xa8, 0x3d, 0x9d, 0xa2, 0xe8,
	0x74, 0x17, 0xdd, 0x50, 0xa0, 0x30, 0xb0, 0x94, 0xa1, 0xed, 0x92, 0xa1, 0x12, 0x58, 0xa2, 0x7b,
	0x9a, 0x18, 0x64, 0x48, 0x6c, 0x2b, 0x3f, 0x90, 0x3a, 0x32, 0xf3, 0x4f, 0xa3, 0x38, 0x6e, 0x8a,
	0x28, 0x65, 0x60, 0xfb, 0xde, 0xf7, 0xfc, 0x9e, 0x9f, 0x5f, 0x02, 0x4e, 0x26, 0xf8, 0xbd, 0x48,
	0x36, 0x13, 0x99, 0x8b, 0x52, 0xe0, 0xae, 0x86, 0xfe, 0x7f, 0x40, 0xa1, 0xc4, 0x43, 0x40, 0x2c,
	0x71, 0x0d, 0xcf, 0x08, 0xfa, 0x04, 0xb1, 0x04, 0xff, 0x04, 0xeb, 0x29, 0x4a, 0x2b, 0xea, 0x22,
	0xb5, 0x6a, 0x80, 0xff, 0x17, 0xec, 0x85, 0x48, 0xab, 0x8c, 0xe3, 0x5f, 0xd0, 0x8f, 0x05, 0xe7,
	0x34, 0x2e, 0x69, 0x2d, 0x33, 0x83, 0x1e, 0xd9, 0x2f, 0xfc, 0x2b, 0x30, 0x57, 0xd1, 0xa1, 0xe9,
	0x3f, 0xb0, 0x63, 0x25, 0x77, 0x91, 0x67, 0x06, 0x83, 0xe9, 0xf7, 0xc9, 0x2e, 0x53, 0xe3, 0x4a,
	0x34, 0xed, 0xbf, 0x18, 0x80, 0xaf, 0x73, 0xf1, 0x40, 0xe3, 0x32, 0xe4, 0x77, 0x22, 0xcf, 0xa2,
	0x92, 0x09, 0x7e, 0xe0, 0x87, 0xa1, 0x53, 0x15, 0x34, 0xd7, 0x19, 0xd5, 0x5c, 0x07, 0x2f, 0x62,
	0x91, 0x53, 0xd7, 0xf4, 0x8c, 0xc0, 0x22, 0x0d, 0xc0, 0x63, 0x40, 0x4c, 0xba, 0x1d, 0x75, 0xeb,
	0xa0, 0xbd, 0x35, 0x94, 0x04, 0x31, 0x89, 0x7f, 0x83, 0x99, 0x45, 0xd2, 0xb5, 0x3c, 0x23, 0x18,
	0x4c, 0xbf, 0xb5, 0xec, 0x2a, 0x92, 0xa4, 0x26, 0xfc, 0x4b, 0xb0, 0x09, 0x2d, 0xaa, 0xb4, 0xc4,
	0x2e, 0x74, 0x59, 0xb1, 0x8e, 0x58, 0xda, 0xa4, 0xe8, 0x91, 0x1d, 0x3c, 0xd2, 0xd7, 0x09, 0x58,
	0xeb, 0x7a, 0xc0, 0x3f, 0xc0, 0x7c, 0xa4, 0x5b, 0x1d, 0xbd, 0x1e, 0x8f, 0x08, 0xce, 0xa0, 0xbf,
	0x10, 0x3c, 0x61, 0xea, 0xb9, 0x7f, 0x76, 0x47, 0x0c, 0x95, 0x7b, 0xd8, 0x26, 0x53, 0x9e, 0x5a,
	0x32, 0x7d, 0x46, 0xd0, 0x5d, 0xd5, 0xc4, 0x72, 0x8e, 0xe7, 0xe0, 0xdc, 0x54, 0x34, 0xdf, 0xea,
	0xee, 0x0a, 0x8c, 0xdf, 0x34, 0xac, 0x6d, 0x47, 0xe3, 0x76, 0x77, 0x58, 0xf1, 0xa9, 0x81, 0x67,
	0xe0, 0x84, 0xbc, 0xa0, 0x79, 0xa9, 0x59, 0xfc, 0xd9, 0xf9, 0xd1, 0xfe, 0x13, 0xea, 0x8a, 0x66,
	0xe0, 0xdc, 0xca, 0x24, 0x2a, 0xe9, 0xd7, 0xe4, 0x17, 0xe0, 0x2c, 0x69, 0x4a, 0xf7, 0xf2, 0x8f,
	0x5e, 0xf0, 0x5e, 0xb5, 0xb1, 0xd5, 0x3f, 0x7d, 0xfe, 0x1a, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xda,
	0x8c, 0x4a, 0xe4, 0x02, 0x00, 0x00,
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
	UpdateProject(ctx context.Context, in *ProjectInformation, opts ...grpc.CallOption) (*Result, error)
	DeleteProject(ctx context.Context, in *Condition, opts ...grpc.CallOption) (*Result, error)
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

func (c *mongoDBClient) UpdateProject(ctx context.Context, in *ProjectInformation, opts ...grpc.CallOption) (*Result, error) {
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

// MongoDBServer is the server API for MongoDB service.
type MongoDBServer interface {
	QueryProjects(*Condition, MongoDB_QueryProjectsServer) error
	InsertProject(context.Context, *ProjectInformation) (*Result, error)
	UpdateProject(context.Context, *ProjectInformation) (*Result, error)
	DeleteProject(context.Context, *Condition) (*Result, error)
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
func (*UnimplementedMongoDBServer) UpdateProject(ctx context.Context, req *ProjectInformation) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (*UnimplementedMongoDBServer) DeleteProject(ctx context.Context, req *Condition) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
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
	in := new(ProjectInformation)
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
		return srv.(MongoDBServer).UpdateProject(ctx, req.(*ProjectInformation))
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
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryProjects",
			Handler:       _MongoDB_QueryProjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mongodb.proto",
}
