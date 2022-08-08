// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: config/service/github/v1/github.proto

package githubv1

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

type AppConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId          int64 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	InstallationId int64 `protobuf:"varint,2,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	// Types that are assignable to Pem:
	//	*AppConfig_KeyPem
	//	*AppConfig_Base64Pem
	Pem isAppConfig_Pem `protobuf_oneof:"pem"`
}

func (x *AppConfig) Reset() {
	*x = AppConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_github_v1_github_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfig) ProtoMessage() {}

func (x *AppConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_github_v1_github_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfig.ProtoReflect.Descriptor instead.
func (*AppConfig) Descriptor() ([]byte, []int) {
	return file_config_service_github_v1_github_proto_rawDescGZIP(), []int{0}
}

func (x *AppConfig) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *AppConfig) GetInstallationId() int64 {
	if x != nil {
		return x.InstallationId
	}
	return 0
}

func (m *AppConfig) GetPem() isAppConfig_Pem {
	if m != nil {
		return m.Pem
	}
	return nil
}

func (x *AppConfig) GetKeyPem() string {
	if x, ok := x.GetPem().(*AppConfig_KeyPem); ok {
		return x.KeyPem
	}
	return ""
}

func (x *AppConfig) GetBase64Pem() string {
	if x, ok := x.GetPem().(*AppConfig_Base64Pem); ok {
		return x.Base64Pem
	}
	return ""
}

type isAppConfig_Pem interface {
	isAppConfig_Pem()
}

type AppConfig_KeyPem struct {
	// a private encryption key from a pem file
	KeyPem string `protobuf:"bytes,3,opt,name=key_pem,json=keyPem,proto3,oneof"`
}

type AppConfig_Base64Pem struct {
	// a base64 encoded private encryption key from a pem file
	Base64Pem string `protobuf:"bytes,4,opt,name=base64_pem,json=base64Pem,proto3,oneof"`
}

func (*AppConfig_KeyPem) isAppConfig_Pem() {}

func (*AppConfig_Base64Pem) isAppConfig_Pem() {}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Auth:
	//	*Config_AccessToken
	//	*Config_AppConfig
	Auth isConfig_Auth `protobuf_oneof:"auth"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_github_v1_github_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_github_v1_github_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_service_github_v1_github_proto_rawDescGZIP(), []int{1}
}

func (m *Config) GetAuth() isConfig_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (x *Config) GetAccessToken() string {
	if x, ok := x.GetAuth().(*Config_AccessToken); ok {
		return x.AccessToken
	}
	return ""
}

func (x *Config) GetAppConfig() *AppConfig {
	if x, ok := x.GetAuth().(*Config_AppConfig); ok {
		return x.AppConfig
	}
	return nil
}

type isConfig_Auth interface {
	isConfig_Auth()
}

type Config_AccessToken struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3,oneof"`
}

type Config_AppConfig struct {
	AppConfig *AppConfig `protobuf:"bytes,2,opt,name=app_config,json=appConfig,proto3,oneof"`
}

func (*Config_AccessToken) isConfig_Auth() {}

func (*Config_AppConfig) isConfig_Auth() {}

var File_config_service_github_v1_github_proto protoreflect.FileDescriptor

var file_config_service_github_v1_github_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1e, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x01, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28,
	0x01, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x48, 0x00, 0x52, 0x06, 0x6b,
	0x65, 0x79, 0x50, 0x65, 0x6d, 0x12, 0x28, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x5f,
	0x70, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x20, 0x01, 0x48, 0x00, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x50, 0x65, 0x6d, 0x42,
	0x0a, 0x0a, 0x03, 0x70, 0x65, 0x6d, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x90, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2c, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x20, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4b, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63,
	0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x09, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x0b, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x46,
	0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66,
	0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_service_github_v1_github_proto_rawDescOnce sync.Once
	file_config_service_github_v1_github_proto_rawDescData = file_config_service_github_v1_github_proto_rawDesc
)

func file_config_service_github_v1_github_proto_rawDescGZIP() []byte {
	file_config_service_github_v1_github_proto_rawDescOnce.Do(func() {
		file_config_service_github_v1_github_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_service_github_v1_github_proto_rawDescData)
	})
	return file_config_service_github_v1_github_proto_rawDescData
}

var file_config_service_github_v1_github_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_service_github_v1_github_proto_goTypes = []interface{}{
	(*AppConfig)(nil), // 0: clutch.config.service.github.v1.AppConfig
	(*Config)(nil),    // 1: clutch.config.service.github.v1.Config
}
var file_config_service_github_v1_github_proto_depIdxs = []int32{
	0, // 0: clutch.config.service.github.v1.Config.app_config:type_name -> clutch.config.service.github.v1.AppConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config_service_github_v1_github_proto_init() }
func file_config_service_github_v1_github_proto_init() {
	if File_config_service_github_v1_github_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_service_github_v1_github_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfig); i {
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
		file_config_service_github_v1_github_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
	file_config_service_github_v1_github_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AppConfig_KeyPem)(nil),
		(*AppConfig_Base64Pem)(nil),
	}
	file_config_service_github_v1_github_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Config_AccessToken)(nil),
		(*Config_AppConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_service_github_v1_github_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_service_github_v1_github_proto_goTypes,
		DependencyIndexes: file_config_service_github_v1_github_proto_depIdxs,
		MessageInfos:      file_config_service_github_v1_github_proto_msgTypes,
	}.Build()
	File_config_service_github_v1_github_proto = out.File
	file_config_service_github_v1_github_proto_rawDesc = nil
	file_config_service_github_v1_github_proto_goTypes = nil
	file_config_service_github_v1_github_proto_depIdxs = nil
}
