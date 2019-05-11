// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/service.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type CreateArtifactTypeRequest struct {
	ArtifactType         *ArtifactType `protobuf:"bytes,1,opt,name=artifact_type,json=artifactType,proto3" json:"artifact_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateArtifactTypeRequest) Reset()         { *m = CreateArtifactTypeRequest{} }
func (m *CreateArtifactTypeRequest) String() string { return proto.CompactTextString(m) }
func (*CreateArtifactTypeRequest) ProtoMessage()    {}
func (*CreateArtifactTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42c32aec9010f89c, []int{0}
}

func (m *CreateArtifactTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArtifactTypeRequest.Unmarshal(m, b)
}
func (m *CreateArtifactTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArtifactTypeRequest.Marshal(b, m, deterministic)
}
func (m *CreateArtifactTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArtifactTypeRequest.Merge(m, src)
}
func (m *CreateArtifactTypeRequest) XXX_Size() int {
	return xxx_messageInfo_CreateArtifactTypeRequest.Size(m)
}
func (m *CreateArtifactTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArtifactTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArtifactTypeRequest proto.InternalMessageInfo

func (m *CreateArtifactTypeRequest) GetArtifactType() *ArtifactType {
	if m != nil {
		return m.ArtifactType
	}
	return nil
}

type CreateArtifactTypeResponse struct {
	ArtifactType         *ArtifactType `protobuf:"bytes,1,opt,name=artifact_type,json=artifactType,proto3" json:"artifact_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateArtifactTypeResponse) Reset()         { *m = CreateArtifactTypeResponse{} }
func (m *CreateArtifactTypeResponse) String() string { return proto.CompactTextString(m) }
func (*CreateArtifactTypeResponse) ProtoMessage()    {}
func (*CreateArtifactTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42c32aec9010f89c, []int{1}
}

func (m *CreateArtifactTypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArtifactTypeResponse.Unmarshal(m, b)
}
func (m *CreateArtifactTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArtifactTypeResponse.Marshal(b, m, deterministic)
}
func (m *CreateArtifactTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArtifactTypeResponse.Merge(m, src)
}
func (m *CreateArtifactTypeResponse) XXX_Size() int {
	return xxx_messageInfo_CreateArtifactTypeResponse.Size(m)
}
func (m *CreateArtifactTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArtifactTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArtifactTypeResponse proto.InternalMessageInfo

func (m *CreateArtifactTypeResponse) GetArtifactType() *ArtifactType {
	if m != nil {
		return m.ArtifactType
	}
	return nil
}

type GetArtifactTypeRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArtifactTypeRequest) Reset()         { *m = GetArtifactTypeRequest{} }
func (m *GetArtifactTypeRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtifactTypeRequest) ProtoMessage()    {}
func (*GetArtifactTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42c32aec9010f89c, []int{2}
}

func (m *GetArtifactTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactTypeRequest.Unmarshal(m, b)
}
func (m *GetArtifactTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactTypeRequest.Marshal(b, m, deterministic)
}
func (m *GetArtifactTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactTypeRequest.Merge(m, src)
}
func (m *GetArtifactTypeRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtifactTypeRequest.Size(m)
}
func (m *GetArtifactTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactTypeRequest proto.InternalMessageInfo

func (m *GetArtifactTypeRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetArtifactTypeResponse struct {
	ArtifactType         *ArtifactType `protobuf:"bytes,1,opt,name=artifact_type,json=artifactType,proto3" json:"artifact_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetArtifactTypeResponse) Reset()         { *m = GetArtifactTypeResponse{} }
func (m *GetArtifactTypeResponse) String() string { return proto.CompactTextString(m) }
func (*GetArtifactTypeResponse) ProtoMessage()    {}
func (*GetArtifactTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42c32aec9010f89c, []int{3}
}

func (m *GetArtifactTypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactTypeResponse.Unmarshal(m, b)
}
func (m *GetArtifactTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactTypeResponse.Marshal(b, m, deterministic)
}
func (m *GetArtifactTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactTypeResponse.Merge(m, src)
}
func (m *GetArtifactTypeResponse) XXX_Size() int {
	return xxx_messageInfo_GetArtifactTypeResponse.Size(m)
}
func (m *GetArtifactTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactTypeResponse proto.InternalMessageInfo

func (m *GetArtifactTypeResponse) GetArtifactType() *ArtifactType {
	if m != nil {
		return m.ArtifactType
	}
	return nil
}

type DeleteArtifactTypeRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteArtifactTypeRequest) Reset()         { *m = DeleteArtifactTypeRequest{} }
func (m *DeleteArtifactTypeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteArtifactTypeRequest) ProtoMessage()    {}
func (*DeleteArtifactTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42c32aec9010f89c, []int{4}
}

func (m *DeleteArtifactTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArtifactTypeRequest.Unmarshal(m, b)
}
func (m *DeleteArtifactTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArtifactTypeRequest.Marshal(b, m, deterministic)
}
func (m *DeleteArtifactTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArtifactTypeRequest.Merge(m, src)
}
func (m *DeleteArtifactTypeRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteArtifactTypeRequest.Size(m)
}
func (m *DeleteArtifactTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArtifactTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArtifactTypeRequest proto.InternalMessageInfo

func (m *DeleteArtifactTypeRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListArtifactTypesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListArtifactTypesRequest) Reset()         { *m = ListArtifactTypesRequest{} }
func (m *ListArtifactTypesRequest) String() string { return proto.CompactTextString(m) }
func (*ListArtifactTypesRequest) ProtoMessage()    {}
func (*ListArtifactTypesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42c32aec9010f89c, []int{5}
}

func (m *ListArtifactTypesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArtifactTypesRequest.Unmarshal(m, b)
}
func (m *ListArtifactTypesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArtifactTypesRequest.Marshal(b, m, deterministic)
}
func (m *ListArtifactTypesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArtifactTypesRequest.Merge(m, src)
}
func (m *ListArtifactTypesRequest) XXX_Size() int {
	return xxx_messageInfo_ListArtifactTypesRequest.Size(m)
}
func (m *ListArtifactTypesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArtifactTypesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListArtifactTypesRequest proto.InternalMessageInfo

type ListArtifactTypesResponse struct {
	ArtifactTypes        []*ArtifactType `protobuf:"bytes,1,rep,name=artifact_types,json=artifactTypes,proto3" json:"artifact_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListArtifactTypesResponse) Reset()         { *m = ListArtifactTypesResponse{} }
func (m *ListArtifactTypesResponse) String() string { return proto.CompactTextString(m) }
func (*ListArtifactTypesResponse) ProtoMessage()    {}
func (*ListArtifactTypesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42c32aec9010f89c, []int{6}
}

func (m *ListArtifactTypesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArtifactTypesResponse.Unmarshal(m, b)
}
func (m *ListArtifactTypesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArtifactTypesResponse.Marshal(b, m, deterministic)
}
func (m *ListArtifactTypesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArtifactTypesResponse.Merge(m, src)
}
func (m *ListArtifactTypesResponse) XXX_Size() int {
	return xxx_messageInfo_ListArtifactTypesResponse.Size(m)
}
func (m *ListArtifactTypesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArtifactTypesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListArtifactTypesResponse proto.InternalMessageInfo

func (m *ListArtifactTypesResponse) GetArtifactTypes() []*ArtifactType {
	if m != nil {
		return m.ArtifactTypes
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateArtifactTypeRequest)(nil), "api.CreateArtifactTypeRequest")
	proto.RegisterType((*CreateArtifactTypeResponse)(nil), "api.CreateArtifactTypeResponse")
	proto.RegisterType((*GetArtifactTypeRequest)(nil), "api.GetArtifactTypeRequest")
	proto.RegisterType((*GetArtifactTypeResponse)(nil), "api.GetArtifactTypeResponse")
	proto.RegisterType((*DeleteArtifactTypeRequest)(nil), "api.DeleteArtifactTypeRequest")
	proto.RegisterType((*ListArtifactTypesRequest)(nil), "api.ListArtifactTypesRequest")
	proto.RegisterType((*ListArtifactTypesResponse)(nil), "api.ListArtifactTypesResponse")
}

func init() { proto.RegisterFile("api/service.proto", fileDescriptor_42c32aec9010f89c) }

var fileDescriptor_42c32aec9010f89c = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x95, 0x46, 0x42, 0x68, 0xa0, 0x8b, 0x6a, 0xa1, 0xa5, 0xf5, 0x96, 0x6e, 0x31, 0x1c,
	0x16, 0x10, 0x89, 0x76, 0x57, 0x42, 0x88, 0x1b, 0xff, 0xc4, 0x05, 0x0e, 0x2c, 0xcb, 0x85, 0x0b,
	0x72, 0x9a, 0x69, 0x6b, 0x91, 0xc4, 0x26, 0x76, 0xb6, 0x54, 0x88, 0x0b, 0x47, 0xae, 0x3c, 0x1a,
	0xaf, 0xc0, 0x23, 0xf0, 0x00, 0x28, 0x8e, 0x91, 0x9a, 0x26, 0xe9, 0x01, 0x8e, 0x99, 0xcf, 0xf9,
	0xbe, 0x9f, 0x67, 0xc6, 0x30, 0xe0, 0x4a, 0x84, 0x1a, 0xf3, 0x0b, 0x31, 0xc3, 0x40, 0xe5, 0xd2,
	0x48, 0xe2, 0x73, 0x25, 0x68, 0xbf, 0xac, 0x73, 0x25, 0xaa, 0x1a, 0x1d, 0x2f, 0xa4, 0x5c, 0x24,
	0x18, 0xda, 0x6a, 0x96, 0x49, 0xc3, 0x8d, 0x90, 0x99, 0x76, 0xea, 0x81, 0x53, 0xed, 0x57, 0x54,
	0xcc, 0x43, 0x4c, 0x95, 0x59, 0x3b, 0xf1, 0x70, 0x5b, 0x34, 0x22, 0x45, 0x6d, 0x78, 0xaa, 0xaa,
	0x03, 0xec, 0x2d, 0x8c, 0x9e, 0xe5, 0xc8, 0x0d, 0x3e, 0xc9, 0x8d, 0x98, 0xf3, 0x99, 0x39, 0x5f,
	0x2b, 0x3c, 0xc3, 0x4f, 0x05, 0x6a, 0x43, 0x1e, 0x42, 0x9f, 0xbb, 0xf2, 0x07, 0xb3, 0x56, 0x38,
	0xf4, 0xa6, 0xde, 0xd1, 0x95, 0x93, 0x41, 0x50, 0xb2, 0xd5, 0x7e, 0xb8, 0xca, 0x37, 0xbe, 0xd8,
	0x39, 0xd0, 0x36, 0x53, 0xad, 0x64, 0xa6, 0xf1, 0x9f, 0x5d, 0x8f, 0x60, 0xff, 0x25, 0x9a, 0x36,
	0xce, 0x3d, 0xe8, 0x89, 0xd8, 0xda, 0xf8, 0x67, 0x3d, 0x11, 0xb3, 0x37, 0x70, 0xa3, 0x71, 0xf2,
	0x3f, 0xc3, 0xef, 0xc3, 0xe8, 0x39, 0x26, 0xd8, 0xde, 0xa7, 0xed, 0x7c, 0x0a, 0xc3, 0x57, 0x42,
	0xd7, 0x00, 0xb4, 0x3b, 0xcb, 0xde, 0xc1, 0xa8, 0x45, 0x73, 0x74, 0x8f, 0x60, 0xaf, 0x46, 0xa7,
	0x87, 0xde, 0xd4, 0x6f, 0xc7, 0xeb, 0x6f, 0xe2, 0xe9, 0x93, 0xdf, 0x3e, 0x5c, 0x7e, 0x8d, 0x86,
	0xc7, 0xdc, 0x70, 0xf2, 0xdd, 0x03, 0xd2, 0x1c, 0x00, 0x99, 0x58, 0x97, 0xce, 0x71, 0xd3, 0xc3,
	0x4e, 0xbd, 0xc2, 0x63, 0xa7, 0xdf, 0x7e, 0xfe, 0xfa, 0xd1, 0x7b, 0xc0, 0xc6, 0x76, 0x15, 0x2f,
	0x8e, 0x79, 0xa2, 0x96, 0xfc, 0x38, 0xac, 0x23, 0x3f, 0xae, 0x37, 0x98, 0x2c, 0xe1, 0xda, 0xd6,
	0x30, 0xc8, 0x81, 0x0d, 0x6a, 0x1f, 0x26, 0x1d, 0xb7, 0x8b, 0x0e, 0x81, 0x5a, 0x84, 0xeb, 0x84,
	0xd4, 0x11, 0xbe, 0x88, 0xf8, 0x2b, 0xf9, 0x0c, 0x83, 0x46, 0x6b, 0xc9, 0x4d, 0x6b, 0xd7, 0x35,
	0x0e, 0x3a, 0xe9, 0x92, 0x5d, 0xde, 0x1d, 0x9b, 0x37, 0x21, 0x3b, 0xaf, 0x4c, 0x56, 0x40, 0x9a,
	0xdb, 0xe1, 0xfa, 0xdd, 0xb9, 0x36, 0x74, 0x3f, 0xa8, 0x5e, 0x67, 0xf0, 0xf7, 0x75, 0x06, 0x2f,
	0xca, 0xa7, 0xcb, 0xee, 0xda, 0xcc, 0xdb, 0xf7, 0x6e, 0xed, 0xca, 0xb4, 0x57, 0x7e, 0xca, 0xde,
	0x4f, 0x17, 0xc2, 0x2c, 0x8b, 0x28, 0x98, 0xc9, 0x34, 0xfc, 0x58, 0x44, 0x38, 0x4f, 0xe4, 0x2a,
	0x4c, 0xdd, 0x26, 0x94, 0x06, 0xd1, 0x25, 0x6b, 0x7f, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x94,
	0xf2, 0x44, 0x0a, 0x6e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetadataClient is the client API for Metadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetadataClient interface {
	CreateArtifactType(ctx context.Context, in *CreateArtifactTypeRequest, opts ...grpc.CallOption) (*CreateArtifactTypeResponse, error)
	GetArtifactType(ctx context.Context, in *GetArtifactTypeRequest, opts ...grpc.CallOption) (*GetArtifactTypeResponse, error)
	ListArtifactTypes(ctx context.Context, in *ListArtifactTypesRequest, opts ...grpc.CallOption) (*ListArtifactTypesResponse, error)
	DeleteArtifactType(ctx context.Context, in *DeleteArtifactTypeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type metadataClient struct {
	cc *grpc.ClientConn
}

func NewMetadataClient(cc *grpc.ClientConn) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) CreateArtifactType(ctx context.Context, in *CreateArtifactTypeRequest, opts ...grpc.CallOption) (*CreateArtifactTypeResponse, error) {
	out := new(CreateArtifactTypeResponse)
	err := c.cc.Invoke(ctx, "/api.Metadata/CreateArtifactType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) GetArtifactType(ctx context.Context, in *GetArtifactTypeRequest, opts ...grpc.CallOption) (*GetArtifactTypeResponse, error) {
	out := new(GetArtifactTypeResponse)
	err := c.cc.Invoke(ctx, "/api.Metadata/GetArtifactType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) ListArtifactTypes(ctx context.Context, in *ListArtifactTypesRequest, opts ...grpc.CallOption) (*ListArtifactTypesResponse, error) {
	out := new(ListArtifactTypesResponse)
	err := c.cc.Invoke(ctx, "/api.Metadata/ListArtifactTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) DeleteArtifactType(ctx context.Context, in *DeleteArtifactTypeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.Metadata/DeleteArtifactType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServer is the server API for Metadata service.
type MetadataServer interface {
	CreateArtifactType(context.Context, *CreateArtifactTypeRequest) (*CreateArtifactTypeResponse, error)
	GetArtifactType(context.Context, *GetArtifactTypeRequest) (*GetArtifactTypeResponse, error)
	ListArtifactTypes(context.Context, *ListArtifactTypesRequest) (*ListArtifactTypesResponse, error)
	DeleteArtifactType(context.Context, *DeleteArtifactTypeRequest) (*empty.Empty, error)
}

func RegisterMetadataServer(s *grpc.Server, srv MetadataServer) {
	s.RegisterService(&_Metadata_serviceDesc, srv)
}

func _Metadata_CreateArtifactType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArtifactTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).CreateArtifactType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Metadata/CreateArtifactType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).CreateArtifactType(ctx, req.(*CreateArtifactTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_GetArtifactType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtifactTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).GetArtifactType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Metadata/GetArtifactType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).GetArtifactType(ctx, req.(*GetArtifactTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_ListArtifactTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArtifactTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ListArtifactTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Metadata/ListArtifactTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ListArtifactTypes(ctx, req.(*ListArtifactTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_DeleteArtifactType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtifactTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).DeleteArtifactType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Metadata/DeleteArtifactType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).DeleteArtifactType(ctx, req.(*DeleteArtifactTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArtifactType",
			Handler:    _Metadata_CreateArtifactType_Handler,
		},
		{
			MethodName: "GetArtifactType",
			Handler:    _Metadata_GetArtifactType_Handler,
		},
		{
			MethodName: "ListArtifactTypes",
			Handler:    _Metadata_ListArtifactTypes_Handler,
		},
		{
			MethodName: "DeleteArtifactType",
			Handler:    _Metadata_DeleteArtifactType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service.proto",
}
