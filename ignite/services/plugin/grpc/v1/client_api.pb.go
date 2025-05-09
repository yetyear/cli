// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: ignite/services/plugin/grpc/v1/client_api.proto

package v1

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChainInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChainId       string                 `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	AppPath       string                 `protobuf:"bytes,2,opt,name=app_path,json=appPath,proto3" json:"app_path,omitempty"`
	ConfigPath    string                 `protobuf:"bytes,3,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
	RpcAddress    string                 `protobuf:"bytes,4,opt,name=rpc_address,json=rpcAddress,proto3" json:"rpc_address,omitempty"`
	Home          string                 `protobuf:"bytes,5,opt,name=home,proto3" json:"home,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChainInfo) Reset() {
	*x = ChainInfo{}
	mi := &file_ignite_services_plugin_grpc_v1_client_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChainInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainInfo) ProtoMessage() {}

func (x *ChainInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ignite_services_plugin_grpc_v1_client_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainInfo.ProtoReflect.Descriptor instead.
func (*ChainInfo) Descriptor() ([]byte, []int) {
	return file_ignite_services_plugin_grpc_v1_client_api_proto_rawDescGZIP(), []int{0}
}

func (x *ChainInfo) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *ChainInfo) GetAppPath() string {
	if x != nil {
		return x.AppPath
	}
	return ""
}

func (x *ChainInfo) GetConfigPath() string {
	if x != nil {
		return x.ConfigPath
	}
	return ""
}

func (x *ChainInfo) GetRpcAddress() string {
	if x != nil {
		return x.RpcAddress
	}
	return ""
}

func (x *ChainInfo) GetHome() string {
	if x != nil {
		return x.Home
	}
	return ""
}

type IgniteInfo struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	CliVersion      string                 `protobuf:"bytes,1,opt,name=cli_version,json=cliVersion,proto3" json:"cli_version,omitempty"`
	GoVersion       string                 `protobuf:"bytes,2,opt,name=go_version,json=goVersion,proto3" json:"go_version,omitempty"`
	SdkVersion      string                 `protobuf:"bytes,3,opt,name=sdk_version,json=sdkVersion,proto3" json:"sdk_version,omitempty"`
	BufVersion      string                 `protobuf:"bytes,4,opt,name=buf_version,json=bufVersion,proto3" json:"buf_version,omitempty"`
	BuildDate       string                 `protobuf:"bytes,5,opt,name=build_date,json=buildDate,proto3" json:"build_date,omitempty"`
	SourceHash      string                 `protobuf:"bytes,6,opt,name=source_hash,json=sourceHash,proto3" json:"source_hash,omitempty"`
	ConfigVersion   string                 `protobuf:"bytes,7,opt,name=config_version,json=configVersion,proto3" json:"config_version,omitempty"`
	Os              string                 `protobuf:"bytes,8,opt,name=os,proto3" json:"os,omitempty"`
	Arch            string                 `protobuf:"bytes,9,opt,name=arch,proto3" json:"arch,omitempty"`
	BuildFromSource bool                   `protobuf:"varint,10,opt,name=build_from_source,json=buildFromSource,proto3" json:"build_from_source,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *IgniteInfo) Reset() {
	*x = IgniteInfo{}
	mi := &file_ignite_services_plugin_grpc_v1_client_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IgniteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IgniteInfo) ProtoMessage() {}

func (x *IgniteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ignite_services_plugin_grpc_v1_client_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IgniteInfo.ProtoReflect.Descriptor instead.
func (*IgniteInfo) Descriptor() ([]byte, []int) {
	return file_ignite_services_plugin_grpc_v1_client_api_proto_rawDescGZIP(), []int{1}
}

func (x *IgniteInfo) GetCliVersion() string {
	if x != nil {
		return x.CliVersion
	}
	return ""
}

func (x *IgniteInfo) GetGoVersion() string {
	if x != nil {
		return x.GoVersion
	}
	return ""
}

func (x *IgniteInfo) GetSdkVersion() string {
	if x != nil {
		return x.SdkVersion
	}
	return ""
}

func (x *IgniteInfo) GetBufVersion() string {
	if x != nil {
		return x.BufVersion
	}
	return ""
}

func (x *IgniteInfo) GetBuildDate() string {
	if x != nil {
		return x.BuildDate
	}
	return ""
}

func (x *IgniteInfo) GetSourceHash() string {
	if x != nil {
		return x.SourceHash
	}
	return ""
}

func (x *IgniteInfo) GetConfigVersion() string {
	if x != nil {
		return x.ConfigVersion
	}
	return ""
}

func (x *IgniteInfo) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *IgniteInfo) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *IgniteInfo) GetBuildFromSource() bool {
	if x != nil {
		return x.BuildFromSource
	}
	return false
}

var File_ignite_services_plugin_grpc_v1_client_api_proto protoreflect.FileDescriptor

const file_ignite_services_plugin_grpc_v1_client_api_proto_rawDesc = "" +
	"\n" +
	"/ignite/services/plugin/grpc/v1/client_api.proto\x12\x1eignite.services.plugin.grpc.v1\"\x97\x01\n" +
	"\tChainInfo\x12\x19\n" +
	"\bchain_id\x18\x01 \x01(\tR\achainId\x12\x19\n" +
	"\bapp_path\x18\x02 \x01(\tR\aappPath\x12\x1f\n" +
	"\vconfig_path\x18\x03 \x01(\tR\n" +
	"configPath\x12\x1f\n" +
	"\vrpc_address\x18\x04 \x01(\tR\n" +
	"rpcAddress\x12\x12\n" +
	"\x04home\x18\x05 \x01(\tR\x04home\"\xc5\x02\n" +
	"\n" +
	"IgniteInfo\x12\x1f\n" +
	"\vcli_version\x18\x01 \x01(\tR\n" +
	"cliVersion\x12\x1d\n" +
	"\n" +
	"go_version\x18\x02 \x01(\tR\tgoVersion\x12\x1f\n" +
	"\vsdk_version\x18\x03 \x01(\tR\n" +
	"sdkVersion\x12\x1f\n" +
	"\vbuf_version\x18\x04 \x01(\tR\n" +
	"bufVersion\x12\x1d\n" +
	"\n" +
	"build_date\x18\x05 \x01(\tR\tbuildDate\x12\x1f\n" +
	"\vsource_hash\x18\x06 \x01(\tR\n" +
	"sourceHash\x12%\n" +
	"\x0econfig_version\x18\a \x01(\tR\rconfigVersion\x12\x0e\n" +
	"\x02os\x18\b \x01(\tR\x02os\x12\x12\n" +
	"\x04arch\x18\t \x01(\tR\x04arch\x12*\n" +
	"\x11build_from_source\x18\n" +
	" \x01(\bR\x0fbuildFromSourceB:Z8github.com/ignite/cli/v29/ignite/services/plugin/grpc/v1b\x06proto3"

var (
	file_ignite_services_plugin_grpc_v1_client_api_proto_rawDescOnce sync.Once
	file_ignite_services_plugin_grpc_v1_client_api_proto_rawDescData []byte
)

func file_ignite_services_plugin_grpc_v1_client_api_proto_rawDescGZIP() []byte {
	file_ignite_services_plugin_grpc_v1_client_api_proto_rawDescOnce.Do(func() {
		file_ignite_services_plugin_grpc_v1_client_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ignite_services_plugin_grpc_v1_client_api_proto_rawDesc), len(file_ignite_services_plugin_grpc_v1_client_api_proto_rawDesc)))
	})
	return file_ignite_services_plugin_grpc_v1_client_api_proto_rawDescData
}

var file_ignite_services_plugin_grpc_v1_client_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ignite_services_plugin_grpc_v1_client_api_proto_goTypes = []any{
	(*ChainInfo)(nil),  // 0: ignite.services.plugin.grpc.v1.ChainInfo
	(*IgniteInfo)(nil), // 1: ignite.services.plugin.grpc.v1.IgniteInfo
}
var file_ignite_services_plugin_grpc_v1_client_api_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ignite_services_plugin_grpc_v1_client_api_proto_init() }
func file_ignite_services_plugin_grpc_v1_client_api_proto_init() {
	if File_ignite_services_plugin_grpc_v1_client_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ignite_services_plugin_grpc_v1_client_api_proto_rawDesc), len(file_ignite_services_plugin_grpc_v1_client_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ignite_services_plugin_grpc_v1_client_api_proto_goTypes,
		DependencyIndexes: file_ignite_services_plugin_grpc_v1_client_api_proto_depIdxs,
		MessageInfos:      file_ignite_services_plugin_grpc_v1_client_api_proto_msgTypes,
	}.Build()
	File_ignite_services_plugin_grpc_v1_client_api_proto = out.File
	file_ignite_services_plugin_grpc_v1_client_api_proto_goTypes = nil
	file_ignite_services_plugin_grpc_v1_client_api_proto_depIdxs = nil
}
