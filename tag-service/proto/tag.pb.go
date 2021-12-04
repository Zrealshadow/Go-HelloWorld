// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/tag.proto

package proto

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

type GetTagListRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State                uint32   `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTagListRequest) Reset()         { *m = GetTagListRequest{} }
func (m *GetTagListRequest) String() string { return proto.CompactTextString(m) }
func (*GetTagListRequest) ProtoMessage()    {}
func (*GetTagListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dc6f15189f1be6, []int{0}
}

func (m *GetTagListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTagListRequest.Unmarshal(m, b)
}
func (m *GetTagListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTagListRequest.Marshal(b, m, deterministic)
}
func (m *GetTagListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTagListRequest.Merge(m, src)
}
func (m *GetTagListRequest) XXX_Size() int {
	return xxx_messageInfo_GetTagListRequest.Size(m)
}
func (m *GetTagListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTagListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTagListRequest proto.InternalMessageInfo

func (m *GetTagListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetTagListRequest) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

type Tag struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State                uint32   `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dc6f15189f1be6, []int{1}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

type GetTagListReply struct {
	List                 []*Tag   `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTagListReply) Reset()         { *m = GetTagListReply{} }
func (m *GetTagListReply) String() string { return proto.CompactTextString(m) }
func (*GetTagListReply) ProtoMessage()    {}
func (*GetTagListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_48dc6f15189f1be6, []int{2}
}

func (m *GetTagListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTagListReply.Unmarshal(m, b)
}
func (m *GetTagListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTagListReply.Marshal(b, m, deterministic)
}
func (m *GetTagListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTagListReply.Merge(m, src)
}
func (m *GetTagListReply) XXX_Size() int {
	return xxx_messageInfo_GetTagListReply.Size(m)
}
func (m *GetTagListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTagListReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTagListReply proto.InternalMessageInfo

func (m *GetTagListReply) GetList() []*Tag {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *GetTagListReply) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTagListRequest)(nil), "proto.GetTagListRequest")
	proto.RegisterType((*Tag)(nil), "proto.Tag")
	proto.RegisterType((*GetTagListReply)(nil), "proto.GetTagListReply")
}

func init() { proto.RegisterFile("proto/tag.proto", fileDescriptor_48dc6f15189f1be6) }

var fileDescriptor_48dc6f15189f1be6 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x6d, 0xbb, 0x15, 0x9c, 0x55, 0x17, 0x07, 0x91, 0xb2, 0x07, 0x29, 0x39, 0xf5, 0xb4,
	0x42, 0x3d, 0x8b, 0xde, 0xbc, 0x88, 0x48, 0x8c, 0x0f, 0x30, 0xee, 0x0e, 0x21, 0xd0, 0x6e, 0x63,
	0x33, 0x0a, 0xfb, 0xf6, 0x62, 0xb2, 0xe8, 0x82, 0x3d, 0x65, 0xf2, 0x7f, 0x93, 0x6f, 0x32, 0xb0,
	0xf0, 0xe3, 0x20, 0xc3, 0x8d, 0x90, 0x5d, 0xc5, 0x0a, 0xcb, 0x78, 0x2c, 0x31, 0xe5, 0xeb, 0xa1,
	0xef, 0x69, 0x9b, 0x90, 0xba, 0x83, 0x8b, 0x47, 0x16, 0x43, 0xf6, 0xc9, 0x05, 0xd1, 0xfc, 0xf1,
	0xc9, 0x41, 0x10, 0x61, 0xb6, 0xa5, 0x9e, 0xab, 0xac, 0xce, 0x9a, 0x13, 0x1d, 0x6b, 0xbc, 0x84,
	0x32, 0x08, 0x09, 0x57, 0x79, 0x9d, 0x35, 0x67, 0x3a, 0x5d, 0xd4, 0x3d, 0x14, 0x86, 0x2c, 0x9e,
	0x43, 0xee, 0x36, 0xb1, 0xbd, 0xd0, 0xb9, 0xdb, 0xfc, 0x0a, 0xf2, 0x29, 0x41, 0x71, 0x28, 0x78,
	0x83, 0xc5, 0xe1, 0x7c, 0xdf, 0xed, 0xf0, 0x1a, 0x66, 0x9d, 0x0b, 0x52, 0x65, 0x75, 0xd1, 0xcc,
	0x5b, 0x48, 0x1f, 0x5d, 0x19, 0xb2, 0x3a, 0xe6, 0xa8, 0xa0, 0xf4, 0x64, 0x79, 0x8c, 0xf6, 0x79,
	0x7b, 0xba, 0x6f, 0x78, 0xf9, 0xc9, 0x74, 0x42, 0xed, 0x33, 0x80, 0x21, 0xfb, 0xca, 0xe3, 0x97,
	0x5b, 0x33, 0x3e, 0x00, 0xfc, 0x0d, 0xc1, 0x6a, 0xff, 0xe0, 0xdf, 0xde, 0xcb, 0xab, 0x09, 0xe2,
	0xbb, 0x9d, 0x3a, 0x7a, 0x3f, 0x8e, 0xe0, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x59, 0xfd, 0xb1,
	0x4f, 0x5b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TagServiceClient interface {
	GetTagList(ctx context.Context, in *GetTagListRequest, opts ...grpc.CallOption) (*GetTagListReply, error)
}

type tagServiceClient struct {
	cc *grpc.ClientConn
}

func NewTagServiceClient(cc *grpc.ClientConn) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) GetTagList(ctx context.Context, in *GetTagListRequest, opts ...grpc.CallOption) (*GetTagListReply, error) {
	out := new(GetTagListReply)
	err := c.cc.Invoke(ctx, "/proto.TagService/GetTagList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
type TagServiceServer interface {
	GetTagList(context.Context, *GetTagListRequest) (*GetTagListReply, error)
}

// UnimplementedTagServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (*UnimplementedTagServiceServer) GetTagList(ctx context.Context, req *GetTagListRequest) (*GetTagListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagList not implemented")
}

func RegisterTagServiceServer(s *grpc.Server, srv TagServiceServer) {
	s.RegisterService(&_TagService_serviceDesc, srv)
}

func _TagService_GetTagList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).GetTagList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TagService/GetTagList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).GetTagList(ctx, req.(*GetTagListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTagList",
			Handler:    _TagService_GetTagList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tag.proto",
}