// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: sourcegraph/v1/sourcegraph.proto

package sourcegraphv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CompareCommitsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Base       string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Head       string `protobuf:"bytes,3,opt,name=head,proto3" json:"head,omitempty"`
}

func (x *CompareCommitsRequest) Reset() {
	*x = CompareCommitsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourcegraph_v1_sourcegraph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareCommitsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareCommitsRequest) ProtoMessage() {}

func (x *CompareCommitsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sourcegraph_v1_sourcegraph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareCommitsRequest.ProtoReflect.Descriptor instead.
func (*CompareCommitsRequest) Descriptor() ([]byte, []int) {
	return file_sourcegraph_v1_sourcegraph_proto_rawDescGZIP(), []int{0}
}

func (x *CompareCommitsRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *CompareCommitsRequest) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *CompareCommitsRequest) GetHead() string {
	if x != nil {
		return x.Head
	}
	return ""
}

type CompareCommitsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commits []*Commit `protobuf:"bytes,1,rep,name=commits,proto3" json:"commits,omitempty"`
}

func (x *CompareCommitsResponse) Reset() {
	*x = CompareCommitsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourcegraph_v1_sourcegraph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareCommitsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareCommitsResponse) ProtoMessage() {}

func (x *CompareCommitsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sourcegraph_v1_sourcegraph_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareCommitsResponse.ProtoReflect.Descriptor instead.
func (*CompareCommitsResponse) Descriptor() ([]byte, []int) {
	return file_sourcegraph_v1_sourcegraph_proto_rawDescGZIP(), []int{1}
}

func (x *CompareCommitsResponse) GetCommits() []*Commit {
	if x != nil {
		return x.Commits
	}
	return nil
}

type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid         string `protobuf:"bytes,1,opt,name=oid,proto3" json:"oid,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Message     string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourcegraph_v1_sourcegraph_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

func (x *Commit) ProtoReflect() protoreflect.Message {
	mi := &file_sourcegraph_v1_sourcegraph_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_sourcegraph_v1_sourcegraph_proto_rawDescGZIP(), []int{2}
}

func (x *Commit) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

func (x *Commit) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Commit) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Commit) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_sourcegraph_v1_sourcegraph_proto protoreflect.FileDescriptor

var file_sourcegraph_v1_sourcegraph_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x22, 0x51,
	0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x73, 0x22, 0x6d, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sourcegraph_v1_sourcegraph_proto_rawDescOnce sync.Once
	file_sourcegraph_v1_sourcegraph_proto_rawDescData = file_sourcegraph_v1_sourcegraph_proto_rawDesc
)

func file_sourcegraph_v1_sourcegraph_proto_rawDescGZIP() []byte {
	file_sourcegraph_v1_sourcegraph_proto_rawDescOnce.Do(func() {
		file_sourcegraph_v1_sourcegraph_proto_rawDescData = protoimpl.X.CompressGZIP(file_sourcegraph_v1_sourcegraph_proto_rawDescData)
	})
	return file_sourcegraph_v1_sourcegraph_proto_rawDescData
}

var file_sourcegraph_v1_sourcegraph_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sourcegraph_v1_sourcegraph_proto_goTypes = []interface{}{
	(*CompareCommitsRequest)(nil),  // 0: clutch.sourcegraph.v1.CompareCommitsRequest
	(*CompareCommitsResponse)(nil), // 1: clutch.sourcegraph.v1.CompareCommitsResponse
	(*Commit)(nil),                 // 2: clutch.sourcegraph.v1.Commit
}
var file_sourcegraph_v1_sourcegraph_proto_depIdxs = []int32{
	2, // 0: clutch.sourcegraph.v1.CompareCommitsResponse.commits:type_name -> clutch.sourcegraph.v1.Commit
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sourcegraph_v1_sourcegraph_proto_init() }
func file_sourcegraph_v1_sourcegraph_proto_init() {
	if File_sourcegraph_v1_sourcegraph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sourcegraph_v1_sourcegraph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareCommitsRequest); i {
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
		file_sourcegraph_v1_sourcegraph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareCommitsResponse); i {
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
		file_sourcegraph_v1_sourcegraph_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commit); i {
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
			RawDescriptor: file_sourcegraph_v1_sourcegraph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sourcegraph_v1_sourcegraph_proto_goTypes,
		DependencyIndexes: file_sourcegraph_v1_sourcegraph_proto_depIdxs,
		MessageInfos:      file_sourcegraph_v1_sourcegraph_proto_msgTypes,
	}.Build()
	File_sourcegraph_v1_sourcegraph_proto = out.File
	file_sourcegraph_v1_sourcegraph_proto_rawDesc = nil
	file_sourcegraph_v1_sourcegraph_proto_goTypes = nil
	file_sourcegraph_v1_sourcegraph_proto_depIdxs = nil
}
