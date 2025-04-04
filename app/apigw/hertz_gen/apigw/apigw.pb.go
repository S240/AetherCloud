// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: apigw.proto

package apigw

import (
	_ "github.com/S240/AetherCloud/app/apigw/hertz_gen/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 文件上传请求
type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileMd5   string `protobuf:"bytes,1,opt,name=file_md5,json=fileMd5,proto3" json:"file_md5,omitempty" form:"file_md5" query:"file_md5"`            // 文件MD5
	FileName  string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty" form:"file_name" query:"file_name"`       // 文件名称
	FileSize  int64  `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty" form:"file_size" query:"file_size"`      // 文件大小
	IsChunked bool   `protobuf:"varint,4,opt,name=is_chunked,json=isChunked,proto3" json:"is_chunked,omitempty" form:"is_chunked" query:"is_chunked"` // 是否分片
	ChunkSize int64  `protobuf:"varint,5,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty" form:"chunk_size" query:"chunk_size"` // 分片大小
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apigw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apigw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_apigw_proto_rawDescGZIP(), []int{0}
}

func (x *UploadFileRequest) GetFileMd5() string {
	if x != nil {
		return x.FileMd5
	}
	return ""
}

func (x *UploadFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadFileRequest) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *UploadFileRequest) GetIsChunked() bool {
	if x != nil {
		return x.IsChunked
	}
	return false
}

func (x *UploadFileRequest) GetChunkSize() int64 {
	if x != nil {
		return x.ChunkSize
	}
	return 0
}

// 预签名信息
type PreSignMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PresignUrl string `protobuf:"bytes,1,opt,name=presign_url,json=presignUrl,proto3" json:"presign_url,omitempty" form:"presign_url" query:"presign_url"`  // 预签名 URL
	PartNumber int32  `protobuf:"varint,2,opt,name=part_number,json=partNumber,proto3" json:"part_number,omitempty" form:"part_number" query:"part_number"` // 分片序号
}

func (x *PreSignMessage) Reset() {
	*x = PreSignMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apigw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreSignMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreSignMessage) ProtoMessage() {}

func (x *PreSignMessage) ProtoReflect() protoreflect.Message {
	mi := &file_apigw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreSignMessage.ProtoReflect.Descriptor instead.
func (*PreSignMessage) Descriptor() ([]byte, []int) {
	return file_apigw_proto_rawDescGZIP(), []int{1}
}

func (x *PreSignMessage) GetPresignUrl() string {
	if x != nil {
		return x.PresignUrl
	}
	return ""
}

func (x *PreSignMessage) GetPartNumber() int32 {
	if x != nil {
		return x.PartNumber
	}
	return 0
}

// 文件上传响应
type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PresignedUrls []*PreSignMessage `protobuf:"bytes,1,rep,name=presigned_urls,json=presignedUrls,proto3" json:"presigned_urls,omitempty" form:"presigned_urls" query:"presigned_urls"` // 预签名信息数组
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apigw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apigw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_apigw_proto_rawDescGZIP(), []int{2}
}

func (x *UploadFileResponse) GetPresignedUrls() []*PreSignMessage {
	if x != nil {
		return x.PresignedUrls
	}
	return nil
}

type PartsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartNumber int32  `protobuf:"varint,1,opt,name=part_number,json=partNumber,proto3" json:"part_number,omitempty" form:"part_number" query:"part_number"` // 分片序号
	Etag       string `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty" form:"etag" query:"etag"`                                              // 分片的 Etag
}

func (x *PartsMessage) Reset() {
	*x = PartsMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apigw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartsMessage) ProtoMessage() {}

func (x *PartsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_apigw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartsMessage.ProtoReflect.Descriptor instead.
func (*PartsMessage) Descriptor() ([]byte, []int) {
	return file_apigw_proto_rawDescGZIP(), []int{3}
}

func (x *PartsMessage) GetPartNumber() int32 {
	if x != nil {
		return x.PartNumber
	}
	return 0
}

func (x *PartsMessage) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

// 上传完成和并请求
type CompletedUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parts   []*PartsMessage `protobuf:"bytes,1,rep,name=parts,proto3" json:"parts,omitempty" form:"parts" query:"parts"`                          // PartsMessages 数组
	FileMd5 string          `protobuf:"bytes,2,opt,name=file_md5,json=fileMd5,proto3" json:"file_md5,omitempty" form:"file_md5" query:"file_md5"` // 文件 MD5
}

func (x *CompletedUploadRequest) Reset() {
	*x = CompletedUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apigw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletedUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletedUploadRequest) ProtoMessage() {}

func (x *CompletedUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apigw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletedUploadRequest.ProtoReflect.Descriptor instead.
func (*CompletedUploadRequest) Descriptor() ([]byte, []int) {
	return file_apigw_proto_rawDescGZIP(), []int{4}
}

func (x *CompletedUploadRequest) GetParts() []*PartsMessage {
	if x != nil {
		return x.Parts
	}
	return nil
}

func (x *CompletedUploadRequest) GetFileMd5() string {
	if x != nil {
		return x.FileMd5
	}
	return ""
}

// 标准响应
type StandardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" form:"code" query:"code"`            // 返回码
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty" form:"message" query:"message"` // 消息描述
	Data    *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty" form:"data" query:"data"`             // 上传响应内容
}

func (x *StandardResponse) Reset() {
	*x = StandardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apigw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandardResponse) ProtoMessage() {}

func (x *StandardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apigw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandardResponse.ProtoReflect.Descriptor instead.
func (*StandardResponse) Descriptor() ([]byte, []int) {
	return file_apigw_proto_rawDescGZIP(), []int{5}
}

func (x *StandardResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StandardResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StandardResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_apigw_proto protoreflect.FileDescriptor

var file_apigw_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x70, 0x69, 0x67, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x70, 0x69, 0x67, 0x77, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x11, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x64, 0x35, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x52, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x73,
	0x69, 0x67, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x52, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x0e, 0x70, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x77, 0x2e, 0x50, 0x72,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x70, 0x72,
	0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x43, 0x0a, 0x0c, 0x50,
	0x61, 0x72, 0x74, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x61, 0x72, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x65, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67,
	0x22, 0x5e, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x67,
	0x77, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05,
	0x70, 0x61, 0x72, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x64,
	0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x64, 0x35,
	0x22, 0x6a, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xda, 0x01, 0x0a,
	0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x61, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50,
	0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x67, 0x77, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x77, 0x2e, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0xd2,
	0xc1, 0x18, 0x12, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x75, 0x72, 0x6c, 0x12, 0x62, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x77, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x77, 0x2e, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x18, 0xd2, 0xc1, 0x18, 0x14, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x32, 0x34, 0x30, 0x2f, 0x41, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x69,
	0x67, 0x77, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x67, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apigw_proto_rawDescOnce sync.Once
	file_apigw_proto_rawDescData = file_apigw_proto_rawDesc
)

func file_apigw_proto_rawDescGZIP() []byte {
	file_apigw_proto_rawDescOnce.Do(func() {
		file_apigw_proto_rawDescData = protoimpl.X.CompressGZIP(file_apigw_proto_rawDescData)
	})
	return file_apigw_proto_rawDescData
}

var file_apigw_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apigw_proto_goTypes = []interface{}{
	(*UploadFileRequest)(nil),      // 0: apigw.UploadFileRequest
	(*PreSignMessage)(nil),         // 1: apigw.PreSignMessage
	(*UploadFileResponse)(nil),     // 2: apigw.UploadFileResponse
	(*PartsMessage)(nil),           // 3: apigw.PartsMessage
	(*CompletedUploadRequest)(nil), // 4: apigw.CompletedUploadRequest
	(*StandardResponse)(nil),       // 5: apigw.StandardResponse
	(*anypb.Any)(nil),              // 6: google.protobuf.Any
}
var file_apigw_proto_depIdxs = []int32{
	1, // 0: apigw.UploadFileResponse.presigned_urls:type_name -> apigw.PreSignMessage
	3, // 1: apigw.CompletedUploadRequest.parts:type_name -> apigw.PartsMessage
	6, // 2: apigw.StandardResponse.data:type_name -> google.protobuf.Any
	0, // 3: apigw.UploadFileService.GetUploadPreSignUrls:input_type -> apigw.UploadFileRequest
	4, // 4: apigw.UploadFileService.CompleteUpload:input_type -> apigw.CompletedUploadRequest
	5, // 5: apigw.UploadFileService.GetUploadPreSignUrls:output_type -> apigw.StandardResponse
	5, // 6: apigw.UploadFileService.CompleteUpload:output_type -> apigw.StandardResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apigw_proto_init() }
func file_apigw_proto_init() {
	if File_apigw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apigw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apigw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreSignMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apigw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apigw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartsMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apigw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletedUploadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apigw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandardResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apigw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apigw_proto_goTypes,
		DependencyIndexes: file_apigw_proto_depIdxs,
		MessageInfos:      file_apigw_proto_msgTypes,
	}.Build()
	File_apigw_proto = out.File
	file_apigw_proto_rawDesc = nil
	file_apigw_proto_goTypes = nil
	file_apigw_proto_depIdxs = nil
}
