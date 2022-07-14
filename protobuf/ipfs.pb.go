// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: k5_proto/proto/ipfs.proto

package protobufpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PublishNFTDataViaIPFSReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonString []byte `protobuf:"bytes,1,opt,name=JsonString,proto3" json:"JsonString,omitempty"`
}

func (x *PublishNFTDataViaIPFSReq) Reset() {
	*x = PublishNFTDataViaIPFSReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k5_proto_proto_ipfs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishNFTDataViaIPFSReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishNFTDataViaIPFSReq) ProtoMessage() {}

func (x *PublishNFTDataViaIPFSReq) ProtoReflect() protoreflect.Message {
	mi := &file_k5_proto_proto_ipfs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishNFTDataViaIPFSReq.ProtoReflect.Descriptor instead.
func (*PublishNFTDataViaIPFSReq) Descriptor() ([]byte, []int) {
	return file_k5_proto_proto_ipfs_proto_rawDescGZIP(), []int{0}
}

func (x *PublishNFTDataViaIPFSReq) GetJsonString() []byte {
	if x != nil {
		return x.JsonString
	}
	return nil
}

type PublishFileToIPFSReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileBytes []byte `protobuf:"bytes,1,opt,name=FileBytes,proto3" json:"FileBytes,omitempty"`
}

func (x *PublishFileToIPFSReq) Reset() {
	*x = PublishFileToIPFSReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k5_proto_proto_ipfs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishFileToIPFSReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishFileToIPFSReq) ProtoMessage() {}

func (x *PublishFileToIPFSReq) ProtoReflect() protoreflect.Message {
	mi := &file_k5_proto_proto_ipfs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishFileToIPFSReq.ProtoReflect.Descriptor instead.
func (*PublishFileToIPFSReq) Descriptor() ([]byte, []int) {
	return file_k5_proto_proto_ipfs_proto_rawDescGZIP(), []int{1}
}

func (x *PublishFileToIPFSReq) GetFileBytes() []byte {
	if x != nil {
		return x.FileBytes
	}
	return nil
}

type PublishToIPFSRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CID string `protobuf:"bytes,1,opt,name=CID,proto3" json:"CID,omitempty"`
}

func (x *PublishToIPFSRes) Reset() {
	*x = PublishToIPFSRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k5_proto_proto_ipfs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishToIPFSRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishToIPFSRes) ProtoMessage() {}

func (x *PublishToIPFSRes) ProtoReflect() protoreflect.Message {
	mi := &file_k5_proto_proto_ipfs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishToIPFSRes.ProtoReflect.Descriptor instead.
func (*PublishToIPFSRes) Descriptor() ([]byte, []int) {
	return file_k5_proto_proto_ipfs_proto_rawDescGZIP(), []int{2}
}

func (x *PublishToIPFSRes) GetCID() string {
	if x != nil {
		return x.CID
	}
	return ""
}

var File_k5_proto_proto_ipfs_proto protoreflect.FileDescriptor

var file_k5_proto_proto_ipfs_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6b, 0x35, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x69, 0x70, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x3a, 0x0a, 0x18, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x4e, 0x46, 0x54, 0x44, 0x61, 0x74, 0x61, 0x56, 0x69, 0x61, 0x49, 0x50, 0x46, 0x53, 0x52, 0x65,
	0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x22, 0x34, 0x0a, 0x14, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x6f, 0x49, 0x50, 0x46, 0x53, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x6c,
	0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x46, 0x69,
	0x6c, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x54, 0x6f, 0x49, 0x50, 0x46, 0x53, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x43,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x49, 0x44, 0x32, 0xb3, 0x01,
	0x0a, 0x07, 0x49, 0x70, 0x66, 0x73, 0x53, 0x76, 0x63, 0x12, 0x57, 0x0a, 0x15, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x4e, 0x46, 0x54, 0x44, 0x61, 0x74, 0x61, 0x56, 0x69, 0x61, 0x49, 0x50,
	0x46, 0x53, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x4e, 0x46, 0x54, 0x44, 0x61, 0x74, 0x61, 0x56, 0x69, 0x61, 0x49,
	0x50, 0x46, 0x53, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x6f, 0x49, 0x50, 0x46, 0x53, 0x52,
	0x65, 0x73, 0x12, 0x4f, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x46, 0x69, 0x6c,
	0x65, 0x54, 0x6f, 0x49, 0x50, 0x46, 0x53, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x6f,
	0x49, 0x50, 0x46, 0x53, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x6f, 0x49, 0x50, 0x46, 0x53,
	0x52, 0x65, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_k5_proto_proto_ipfs_proto_rawDescOnce sync.Once
	file_k5_proto_proto_ipfs_proto_rawDescData = file_k5_proto_proto_ipfs_proto_rawDesc
)

func file_k5_proto_proto_ipfs_proto_rawDescGZIP() []byte {
	file_k5_proto_proto_ipfs_proto_rawDescOnce.Do(func() {
		file_k5_proto_proto_ipfs_proto_rawDescData = protoimpl.X.CompressGZIP(file_k5_proto_proto_ipfs_proto_rawDescData)
	})
	return file_k5_proto_proto_ipfs_proto_rawDescData
}

var file_k5_proto_proto_ipfs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_k5_proto_proto_ipfs_proto_goTypes = []interface{}{
	(*PublishNFTDataViaIPFSReq)(nil), // 0: protobuf.PublishNFTDataViaIPFSReq
	(*PublishFileToIPFSReq)(nil),     // 1: protobuf.PublishFileToIPFSReq
	(*PublishToIPFSRes)(nil),         // 2: protobuf.PublishToIPFSRes
}
var file_k5_proto_proto_ipfs_proto_depIdxs = []int32{
	0, // 0: protobuf.IpfsSvc.PublishNFTDataViaIPFS:input_type -> protobuf.PublishNFTDataViaIPFSReq
	1, // 1: protobuf.IpfsSvc.PublishFileToIPFS:input_type -> protobuf.PublishFileToIPFSReq
	2, // 2: protobuf.IpfsSvc.PublishNFTDataViaIPFS:output_type -> protobuf.PublishToIPFSRes
	2, // 3: protobuf.IpfsSvc.PublishFileToIPFS:output_type -> protobuf.PublishToIPFSRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_k5_proto_proto_ipfs_proto_init() }
func file_k5_proto_proto_ipfs_proto_init() {
	if File_k5_proto_proto_ipfs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k5_proto_proto_ipfs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishNFTDataViaIPFSReq); i {
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
		file_k5_proto_proto_ipfs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishFileToIPFSReq); i {
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
		file_k5_proto_proto_ipfs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishToIPFSRes); i {
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
			RawDescriptor: file_k5_proto_proto_ipfs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_k5_proto_proto_ipfs_proto_goTypes,
		DependencyIndexes: file_k5_proto_proto_ipfs_proto_depIdxs,
		MessageInfos:      file_k5_proto_proto_ipfs_proto_msgTypes,
	}.Build()
	File_k5_proto_proto_ipfs_proto = out.File
	file_k5_proto_proto_ipfs_proto_rawDesc = nil
	file_k5_proto_proto_ipfs_proto_goTypes = nil
	file_k5_proto_proto_ipfs_proto_depIdxs = nil
}
