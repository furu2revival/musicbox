// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/debug/echo.proto

package debug

import (
	_ "github.com/furu2revival/musicbox/protobuf/custom_option"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EchoServiceEchoV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoServiceEchoV1Request) Reset() {
	*x = EchoServiceEchoV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_debug_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoServiceEchoV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoServiceEchoV1Request) ProtoMessage() {}

func (x *EchoServiceEchoV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_debug_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoServiceEchoV1Request.ProtoReflect.Descriptor instead.
func (*EchoServiceEchoV1Request) Descriptor() ([]byte, []int) {
	return file_api_debug_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoServiceEchoV1Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoServiceEchoV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *EchoServiceEchoV1Response) Reset() {
	*x = EchoServiceEchoV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_debug_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoServiceEchoV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoServiceEchoV1Response) ProtoMessage() {}

func (x *EchoServiceEchoV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_debug_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoServiceEchoV1Response.ProtoReflect.Descriptor instead.
func (*EchoServiceEchoV1Response) Descriptor() ([]byte, []int) {
	return file_api_debug_echo_proto_rawDescGZIP(), []int{1}
}

func (x *EchoServiceEchoV1Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EchoServiceEchoV1Response) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_api_debug_echo_proto protoreflect.FileDescriptor

var file_api_debug_echo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2f, 0x65, 0x63, 0x68, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x1a, 0x21, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x18, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6f, 0x0a, 0x19, 0x45,
	0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x64, 0x0a, 0x0b,
	0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x45,
	0x63, 0x68, 0x6f, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x63, 0x68,
	0x6f, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x94, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x42, 0x09, 0x45, 0x63, 0x68, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75,
	0x72, 0x75, 0x32, 0x72, 0x65, 0x76, 0x69, 0x76, 0x61, 0x6c, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x62, 0x6f, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0xa2, 0x02, 0x03, 0x41, 0x44, 0x58, 0xaa, 0x02, 0x09, 0x41,
	0x70, 0x69, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0xca, 0x02, 0x09, 0x41, 0x70, 0x69, 0x5c, 0x44,
	0x65, 0x62, 0x75, 0x67, 0xe2, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x44, 0x65, 0x62, 0x75, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_debug_echo_proto_rawDescOnce sync.Once
	file_api_debug_echo_proto_rawDescData = file_api_debug_echo_proto_rawDesc
)

func file_api_debug_echo_proto_rawDescGZIP() []byte {
	file_api_debug_echo_proto_rawDescOnce.Do(func() {
		file_api_debug_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_debug_echo_proto_rawDescData)
	})
	return file_api_debug_echo_proto_rawDescData
}

var file_api_debug_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_debug_echo_proto_goTypes = []any{
	(*EchoServiceEchoV1Request)(nil),  // 0: api.debug.EchoServiceEchoV1Request
	(*EchoServiceEchoV1Response)(nil), // 1: api.debug.EchoServiceEchoV1Response
	(*timestamppb.Timestamp)(nil),     // 2: google.protobuf.Timestamp
}
var file_api_debug_echo_proto_depIdxs = []int32{
	2, // 0: api.debug.EchoServiceEchoV1Response.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: api.debug.EchoService.EchoV1:input_type -> api.debug.EchoServiceEchoV1Request
	1, // 2: api.debug.EchoService.EchoV1:output_type -> api.debug.EchoServiceEchoV1Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_debug_echo_proto_init() }
func file_api_debug_echo_proto_init() {
	if File_api_debug_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_debug_echo_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EchoServiceEchoV1Request); i {
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
		file_api_debug_echo_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EchoServiceEchoV1Response); i {
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
			RawDescriptor: file_api_debug_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_debug_echo_proto_goTypes,
		DependencyIndexes: file_api_debug_echo_proto_depIdxs,
		MessageInfos:      file_api_debug_echo_proto_msgTypes,
	}.Build()
	File_api_debug_echo_proto = out.File
	file_api_debug_echo_proto_rawDesc = nil
	file_api_debug_echo_proto_goTypes = nil
	file_api_debug_echo_proto_depIdxs = nil
}
