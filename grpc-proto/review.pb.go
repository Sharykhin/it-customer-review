// Code generated by protoc-gen-go. DO NOT EDIT.
// source: review.proto

/*
Package review is a generated protocol buffer package.

It is generated from these files:
	review.proto

It has these top-level messages:
	ReviewRequest
	ReviewResponse
*/
package review

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for creating a new fail mail
type ReviewRequest struct {
	Name      string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=Email" json:"Email,omitempty"`
	Content   string `protobuf:"bytes,4,opt,name=Content" json:"Content,omitempty"`
	Published bool   `protobuf:"varint,5,opt,name=Published" json:"Published,omitempty"`
	Score     uint64 `protobuf:"varint,6,opt,name=Score" json:"Score,omitempty"`
	Category  string `protobuf:"bytes,7,opt,name=Category" json:"Category,omitempty"`
}

func (m *ReviewRequest) Reset()                    { *m = ReviewRequest{} }
func (m *ReviewRequest) String() string            { return proto.CompactTextString(m) }
func (*ReviewRequest) ProtoMessage()               {}
func (*ReviewRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReviewRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReviewRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ReviewRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ReviewRequest) GetPublished() bool {
	if m != nil {
		return m.Published
	}
	return false
}

func (m *ReviewRequest) GetScore() uint64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *ReviewRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

// Response of fail mail
type ReviewResponse struct {
	ID        string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=Email" json:"Email,omitempty"`
	Content   string `protobuf:"bytes,4,opt,name=Content" json:"Content,omitempty"`
	Published bool   `protobuf:"varint,5,opt,name=Published" json:"Published,omitempty"`
	Score     uint64 `protobuf:"varint,6,opt,name=Score" json:"Score,omitempty"`
	Category  string `protobuf:"bytes,7,opt,name=Category" json:"Category,omitempty"`
	CreatedAt string `protobuf:"bytes,8,opt,name=CreatedAt" json:"CreatedAt,omitempty"`
}

func (m *ReviewResponse) Reset()                    { *m = ReviewResponse{} }
func (m *ReviewResponse) String() string            { return proto.CompactTextString(m) }
func (*ReviewResponse) ProtoMessage()               {}
func (*ReviewResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReviewResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ReviewResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReviewResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ReviewResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ReviewResponse) GetPublished() bool {
	if m != nil {
		return m.Published
	}
	return false
}

func (m *ReviewResponse) GetScore() uint64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *ReviewResponse) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ReviewResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*ReviewRequest)(nil), "review.ReviewRequest")
	proto.RegisterType((*ReviewResponse)(nil), "review.ReviewResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Review service

type ReviewClient interface {
	Create(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewResponse, error)
	Update(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewResponse, error)
}

type reviewClient struct {
	cc *grpc.ClientConn
}

func NewReviewClient(cc *grpc.ClientConn) ReviewClient {
	return &reviewClient{cc}
}

func (c *reviewClient) Create(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewResponse, error) {
	out := new(ReviewResponse)
	err := grpc.Invoke(ctx, "/review.Review/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) Update(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*ReviewResponse, error) {
	out := new(ReviewResponse)
	err := grpc.Invoke(ctx, "/review.Review/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Review service

type ReviewServer interface {
	Create(context.Context, *ReviewRequest) (*ReviewResponse, error)
	Update(context.Context, *ReviewRequest) (*ReviewResponse, error)
}

func RegisterReviewServer(s *grpc.Server, srv ReviewServer) {
	s.RegisterService(&_Review_serviceDesc, srv)
}

func _Review_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.Review/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).Create(ctx, req.(*ReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.Review/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).Update(ctx, req.(*ReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Review_serviceDesc = grpc.ServiceDesc{
	ServiceName: "review.Review",
	HandlerType: (*ReviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Review_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Review_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review.proto",
}

func init() { proto.RegisterFile("review.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x2d, 0xcb,
	0x4c, 0x2d, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x16, 0x33, 0x72,
	0xf1, 0x06, 0x81, 0x99, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x7e,
	0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x08, 0x17, 0xab,
	0x6b, 0x6e, 0x62, 0x66, 0x8e, 0x04, 0x33, 0x58, 0x10, 0xc2, 0x11, 0x92, 0xe0, 0x62, 0x77, 0xce,
	0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x91, 0x60, 0x01, 0x8b, 0xc3, 0xb8, 0x42, 0x32, 0x5c, 0x9c, 0x01,
	0xa5, 0x49, 0x39, 0x99, 0xc5, 0x19, 0xa9, 0x29, 0x12, 0xac, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x08,
	0x01, 0x90, 0x69, 0xc1, 0xc9, 0xf9, 0x45, 0xa9, 0x12, 0x6c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x10,
	0x8e, 0x90, 0x14, 0x17, 0x87, 0x73, 0x62, 0x49, 0x6a, 0x7a, 0x7e, 0x51, 0xa5, 0x04, 0x3b, 0xd8,
	0x38, 0x38, 0x5f, 0xe9, 0x12, 0x23, 0x17, 0x1f, 0xcc, 0x95, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9,
	0x42, 0x7c, 0x5c, 0x4c, 0x9e, 0x2e, 0x12, 0x8c, 0x60, 0x85, 0x4c, 0x9e, 0x2e, 0x83, 0xd1, 0xd9,
	0x20, 0xf3, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x53, 0x1c, 0x4b, 0x24, 0x38, 0xc0, 0x92, 0x08,
	0x01, 0xa3, 0x3a, 0x2e, 0x36, 0x88, 0x9f, 0x84, 0x2c, 0xb9, 0xd8, 0x20, 0xc2, 0x42, 0xa2, 0x7a,
	0xd0, 0x58, 0x42, 0x89, 0x13, 0x29, 0x31, 0x74, 0x61, 0x48, 0x20, 0x28, 0x31, 0x80, 0xb4, 0x86,
	0x16, 0xa4, 0x90, 0xa3, 0x35, 0x89, 0x0d, 0x9c, 0x12, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf6, 0x8e, 0x2a, 0xb7, 0x19, 0x02, 0x00, 0x00,
}
