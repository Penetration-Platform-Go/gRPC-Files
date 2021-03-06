// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

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

type Username struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Username) Reset()         { *m = Username{} }
func (m *Username) String() string { return proto.CompactTextString(m) }
func (*Username) ProtoMessage()    {}
func (*Username) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Username) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Username.Unmarshal(m, b)
}
func (m *Username) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Username.Marshal(b, m, deterministic)
}
func (m *Username) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Username.Merge(m, src)
}
func (m *Username) XXX_Size() int {
	return xxx_messageInfo_Username.Size(m)
}
func (m *Username) XXX_DiscardUnknown() {
	xxx_messageInfo_Username.DiscardUnknown(m)
}

var xxx_messageInfo_Username proto.InternalMessageInfo

func (m *Username) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Token struct {
	JWT                  string   `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetJWT() string {
	if m != nil {
		return m.JWT
	}
	return ""
}

func init() {
	proto.RegisterType((*Username)(nil), "auth.Username")
	proto.RegisterType((*Token)(nil), "auth.Token")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xd4, 0xb8, 0x38, 0x42, 0x8b,
	0x53, 0x8b, 0xf2, 0x12, 0x73, 0x53, 0x85, 0xa4, 0xb8, 0x38, 0x4a, 0xa1, 0x6c, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x38, 0x5f, 0x49, 0x92, 0x8b, 0x35, 0x24, 0x3f, 0x3b, 0x35, 0x4f, 0x48,
	0x80, 0x8b, 0xd9, 0x2b, 0x3c, 0x04, 0x2a, 0x0f, 0x62, 0x1a, 0x25, 0x71, 0xb1, 0x38, 0x96, 0x96,
	0x64, 0x08, 0x19, 0x72, 0x09, 0xb9, 0xa7, 0x96, 0xc0, 0x4c, 0x73, 0xaa, 0x84, 0xa8, 0xe7, 0xd6,
	0x03, 0xdb, 0x09, 0xe6, 0x48, 0xf1, 0x41, 0x38, 0x70, 0x1b, 0xd5, 0xb9, 0x38, 0xdc, 0x53, 0x4b,
	0x20, 0x0a, 0xd1, 0xe4, 0xa4, 0x90, 0x35, 0x26, 0xb1, 0x81, 0xdd, 0x6c, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x6b, 0x14, 0xd5, 0xf4, 0xc1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	GetUsernameByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Username, error)
	GetToken(ctx context.Context, in *Username, opts ...grpc.CallOption) (*Token, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetUsernameByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Username, error) {
	out := new(Username)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetUsernameByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetToken(ctx context.Context, in *Username, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	GetUsernameByToken(context.Context, *Token) (*Username, error)
	GetToken(context.Context, *Username) (*Token, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) GetUsernameByToken(ctx context.Context, req *Token) (*Username, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsernameByToken not implemented")
}
func (*UnimplementedAuthServer) GetToken(ctx context.Context, req *Username) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_GetUsernameByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUsernameByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetUsernameByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUsernameByToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetToken(ctx, req.(*Username))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsernameByToken",
			Handler:    _Auth_GetUsernameByToken_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _Auth_GetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
