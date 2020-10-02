// Code generated by protoc-gen-go. DO NOT EDIT.
// source: memories.proto

package memoriespb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Memories struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Video                string   `protobuf:"bytes,4,opt,name=video,proto3" json:"video,omitempty"`
	Thumbnail            string   `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Memories) Reset()         { *m = Memories{} }
func (m *Memories) String() string { return proto.CompactTextString(m) }
func (*Memories) ProtoMessage()    {}
func (*Memories) Descriptor() ([]byte, []int) {
	return fileDescriptor_f269ed14db9bbae1, []int{0}
}

func (m *Memories) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Memories.Unmarshal(m, b)
}
func (m *Memories) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Memories.Marshal(b, m, deterministic)
}
func (m *Memories) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memories.Merge(m, src)
}
func (m *Memories) XXX_Size() int {
	return xxx_messageInfo_Memories.Size(m)
}
func (m *Memories) XXX_DiscardUnknown() {
	xxx_messageInfo_Memories.DiscardUnknown(m)
}

var xxx_messageInfo_Memories proto.InternalMessageInfo

func (m *Memories) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Memories) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Memories) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Memories) GetVideo() string {
	if m != nil {
		return m.Video
	}
	return ""
}

func (m *Memories) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *Memories) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Memories) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListVideosRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListVideosRequest) Reset()         { *m = ListVideosRequest{} }
func (m *ListVideosRequest) String() string { return proto.CompactTextString(m) }
func (*ListVideosRequest) ProtoMessage()    {}
func (*ListVideosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f269ed14db9bbae1, []int{1}
}

func (m *ListVideosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVideosRequest.Unmarshal(m, b)
}
func (m *ListVideosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVideosRequest.Marshal(b, m, deterministic)
}
func (m *ListVideosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVideosRequest.Merge(m, src)
}
func (m *ListVideosRequest) XXX_Size() int {
	return xxx_messageInfo_ListVideosRequest.Size(m)
}
func (m *ListVideosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVideosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListVideosRequest proto.InternalMessageInfo

type ListVideosResponse struct {
	Memories             *Memories `protobuf:"bytes,1,opt,name=memories,proto3" json:"memories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListVideosResponse) Reset()         { *m = ListVideosResponse{} }
func (m *ListVideosResponse) String() string { return proto.CompactTextString(m) }
func (*ListVideosResponse) ProtoMessage()    {}
func (*ListVideosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f269ed14db9bbae1, []int{2}
}

func (m *ListVideosResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVideosResponse.Unmarshal(m, b)
}
func (m *ListVideosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVideosResponse.Marshal(b, m, deterministic)
}
func (m *ListVideosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVideosResponse.Merge(m, src)
}
func (m *ListVideosResponse) XXX_Size() int {
	return xxx_messageInfo_ListVideosResponse.Size(m)
}
func (m *ListVideosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVideosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListVideosResponse proto.InternalMessageInfo

func (m *ListVideosResponse) GetMemories() *Memories {
	if m != nil {
		return m.Memories
	}
	return nil
}

type UploadVideoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Video                string   `protobuf:"bytes,4,opt,name=video,proto3" json:"video,omitempty"`
	Thumbnail            string   `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadVideoRequest) Reset()         { *m = UploadVideoRequest{} }
func (m *UploadVideoRequest) String() string { return proto.CompactTextString(m) }
func (*UploadVideoRequest) ProtoMessage()    {}
func (*UploadVideoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f269ed14db9bbae1, []int{3}
}

func (m *UploadVideoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadVideoRequest.Unmarshal(m, b)
}
func (m *UploadVideoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadVideoRequest.Marshal(b, m, deterministic)
}
func (m *UploadVideoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadVideoRequest.Merge(m, src)
}
func (m *UploadVideoRequest) XXX_Size() int {
	return xxx_messageInfo_UploadVideoRequest.Size(m)
}
func (m *UploadVideoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadVideoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadVideoRequest proto.InternalMessageInfo

func (m *UploadVideoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UploadVideoRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UploadVideoRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UploadVideoRequest) GetVideo() string {
	if m != nil {
		return m.Video
	}
	return ""
}

func (m *UploadVideoRequest) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *UploadVideoRequest) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UploadVideoRequest) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type UploadVideoResponse struct {
	Memories             *Memories `protobuf:"bytes,1,opt,name=memories,proto3" json:"memories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UploadVideoResponse) Reset()         { *m = UploadVideoResponse{} }
func (m *UploadVideoResponse) String() string { return proto.CompactTextString(m) }
func (*UploadVideoResponse) ProtoMessage()    {}
func (*UploadVideoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f269ed14db9bbae1, []int{4}
}

func (m *UploadVideoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadVideoResponse.Unmarshal(m, b)
}
func (m *UploadVideoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadVideoResponse.Marshal(b, m, deterministic)
}
func (m *UploadVideoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadVideoResponse.Merge(m, src)
}
func (m *UploadVideoResponse) XXX_Size() int {
	return xxx_messageInfo_UploadVideoResponse.Size(m)
}
func (m *UploadVideoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadVideoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadVideoResponse proto.InternalMessageInfo

func (m *UploadVideoResponse) GetMemories() *Memories {
	if m != nil {
		return m.Memories
	}
	return nil
}

func init() {
	proto.RegisterType((*Memories)(nil), "memories.Memories")
	proto.RegisterType((*ListVideosRequest)(nil), "memories.ListVideosRequest")
	proto.RegisterType((*ListVideosResponse)(nil), "memories.ListVideosResponse")
	proto.RegisterType((*UploadVideoRequest)(nil), "memories.UploadVideoRequest")
	proto.RegisterType((*UploadVideoResponse)(nil), "memories.UploadVideoResponse")
}

func init() {
	proto.RegisterFile("memories.proto", fileDescriptor_f269ed14db9bbae1)
}

var fileDescriptor_f269ed14db9bbae1 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0xb1, 0x4e, 0x3a, 0x41,
	0x10, 0xc6, 0xb3, 0xf7, 0xff, 0x83, 0x30, 0x18, 0xd4, 0xc1, 0x98, 0x13, 0x21, 0x21, 0x57, 0x19,
	0x0b, 0x50, 0xec, 0xec, 0x30, 0xda, 0x69, 0x83, 0xd1, 0xc2, 0xc6, 0x2c, 0xec, 0x04, 0x37, 0x39,
	0x6e, 0xd7, 0xdb, 0x85, 0x07, 0xf0, 0x15, 0x7c, 0x25, 0x6b, 0x1b, 0x4b, 0x5b, 0x1f, 0xc4, 0xdc,
	0xde, 0x1d, 0x9c, 0xa2, 0x8d, 0x9d, 0xe5, 0x7c, 0xbf, 0x99, 0xf9, 0x76, 0xbf, 0xbd, 0x83, 0xfa,
	0x94, 0xa6, 0x2a, 0x96, 0x64, 0xba, 0x3a, 0x56, 0x56, 0x61, 0x25, 0xaf, 0x9b, 0xad, 0x89, 0x52,
	0x93, 0x90, 0x7a, 0x5c, 0xcb, 0x1e, 0x8f, 0x22, 0x65, 0xb9, 0x95, 0x2a, 0xca, 0xfa, 0x82, 0x67,
	0x06, 0x95, 0xcb, 0xac, 0x15, 0xeb, 0xe0, 0x49, 0xe1, 0xb3, 0x0e, 0xdb, 0xaf, 0x0e, 0x3d, 0x29,
	0x70, 0x1b, 0x4a, 0x56, 0xda, 0x90, 0x7c, 0xcf, 0x49, 0x69, 0x81, 0x1d, 0xa8, 0x09, 0x32, 0xe3,
	0x58, 0xea, 0x64, 0x91, 0xff, 0xcf, 0xb1, 0xa2, 0x94, 0xcc, 0xcd, 0xa5, 0x20, 0xe5, 0xff, 0x4f,
	0xe7, 0x5c, 0x81, 0x2d, 0xa8, 0xda, 0xfb, 0xd9, 0x74, 0x14, 0x71, 0x19, 0xfa, 0x25, 0x47, 0x96,
	0x02, 0xb6, 0x01, 0xc6, 0x31, 0x71, 0x4b, 0xe2, 0x8e, 0x5b, 0xbf, 0x9c, 0xe2, 0x4c, 0x19, 0xd8,
	0x04, 0xcf, 0xb4, 0xc8, 0xf1, 0x5a, 0x8a, 0x33, 0x65, 0x60, 0x83, 0x06, 0x6c, 0x5d, 0x48, 0x63,
	0x6f, 0x12, 0x23, 0x33, 0xa4, 0x87, 0x19, 0x19, 0x1b, 0x9c, 0x01, 0x16, 0x45, 0xa3, 0x55, 0x64,
	0x08, 0xbb, 0xb0, 0xc8, 0xc6, 0x5d, 0xb5, 0xd6, 0xc7, 0xee, 0x22, 0xbc, 0x3c, 0x8a, 0xe1, 0xa2,
	0x27, 0x78, 0x61, 0x80, 0xd7, 0x3a, 0x54, 0x5c, 0xb8, 0x45, 0xd9, 0xf2, 0x3f, 0x9c, 0xd5, 0x39,
	0x34, 0x3e, 0xdd, 0xe7, 0x77, 0xb9, 0xf4, 0xdf, 0x18, 0x6c, 0xe4, 0xf2, 0x15, 0xc5, 0x73, 0x39,
	0x26, 0xe4, 0x00, 0xcb, 0xc4, 0x71, 0x6f, 0x39, 0xbf, 0xf2, 0x38, 0xcd, 0xd6, 0xf7, 0x30, 0x3d,
	0x4c, 0xb0, 0xf3, 0xf8, 0xfa, 0xfe, 0xe4, 0x6d, 0x62, 0xdd, 0x7d, 0xb6, 0xf3, 0xa3, 0x9e, 0x8b,
	0xc5, 0x1c, 0x32, 0x24, 0xa8, 0x15, 0x4e, 0x8f, 0x85, 0x35, 0xab, 0x8f, 0xd4, 0x6c, 0xff, 0x40,
	0x33, 0x97, 0x5d, 0xe7, 0xd2, 0x08, 0xbe, 0xb8, 0x9c, 0xb0, 0x83, 0xd3, 0xf5, 0x5b, 0xc8, 0x47,
	0xf5, 0x68, 0x54, 0x76, 0x3f, 0xcb, 0xf1, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x8f, 0x16,
	0x29, 0x66, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MemoriesServiceClient is the client API for MemoriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MemoriesServiceClient interface {
	ListVideos(ctx context.Context, in *ListVideosRequest, opts ...grpc.CallOption) (MemoriesService_ListVideosClient, error)
	UploadVideo(ctx context.Context, in *UploadVideoRequest, opts ...grpc.CallOption) (*UploadVideoResponse, error)
}

type memoriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemoriesServiceClient(cc grpc.ClientConnInterface) MemoriesServiceClient {
	return &memoriesServiceClient{cc}
}

func (c *memoriesServiceClient) ListVideos(ctx context.Context, in *ListVideosRequest, opts ...grpc.CallOption) (MemoriesService_ListVideosClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MemoriesService_serviceDesc.Streams[0], "/memories.MemoriesService/ListVideos", opts...)
	if err != nil {
		return nil, err
	}
	x := &memoriesServiceListVideosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MemoriesService_ListVideosClient interface {
	Recv() (*ListVideosResponse, error)
	grpc.ClientStream
}

type memoriesServiceListVideosClient struct {
	grpc.ClientStream
}

func (x *memoriesServiceListVideosClient) Recv() (*ListVideosResponse, error) {
	m := new(ListVideosResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *memoriesServiceClient) UploadVideo(ctx context.Context, in *UploadVideoRequest, opts ...grpc.CallOption) (*UploadVideoResponse, error) {
	out := new(UploadVideoResponse)
	err := c.cc.Invoke(ctx, "/memories.MemoriesService/UploadVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemoriesServiceServer is the server API for MemoriesService service.
type MemoriesServiceServer interface {
	ListVideos(*ListVideosRequest, MemoriesService_ListVideosServer) error
	UploadVideo(context.Context, *UploadVideoRequest) (*UploadVideoResponse, error)
}

// UnimplementedMemoriesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMemoriesServiceServer struct {
}

func (*UnimplementedMemoriesServiceServer) ListVideos(req *ListVideosRequest, srv MemoriesService_ListVideosServer) error {
	return status.Errorf(codes.Unimplemented, "method ListVideos not implemented")
}
func (*UnimplementedMemoriesServiceServer) UploadVideo(ctx context.Context, req *UploadVideoRequest) (*UploadVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}

func RegisterMemoriesServiceServer(s *grpc.Server, srv MemoriesServiceServer) {
	s.RegisterService(&_MemoriesService_serviceDesc, srv)
}

func _MemoriesService_ListVideos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListVideosRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MemoriesServiceServer).ListVideos(m, &memoriesServiceListVideosServer{stream})
}

type MemoriesService_ListVideosServer interface {
	Send(*ListVideosResponse) error
	grpc.ServerStream
}

type memoriesServiceListVideosServer struct {
	grpc.ServerStream
}

func (x *memoriesServiceListVideosServer) Send(m *ListVideosResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MemoriesService_UploadVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoriesServiceServer).UploadVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memories.MemoriesService/UploadVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoriesServiceServer).UploadVideo(ctx, req.(*UploadVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MemoriesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "memories.MemoriesService",
	HandlerType: (*MemoriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadVideo",
			Handler:    _MemoriesService_UploadVideo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListVideos",
			Handler:       _MemoriesService_ListVideos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "memories.proto",
}
