// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: control/error.proto

package controlv1

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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_control_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_control_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_control_error_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_control_error_proto protoreflect.FileDescriptor

var file_control_error_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6c, 0x6f, 0x6e, 0x67, 0x62, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x22,
	0x2d, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0xf1,
	0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x62, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42,
	0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x62, 0x72,
	0x69, 0x64, 0x67, 0x65, 0x61, 0x70, 0x70, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x4c, 0x43, 0x58, 0xaa, 0x02, 0x18, 0x4c, 0x6f, 0x6e, 0x67, 0x62,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x4c, 0x6f, 0x6e, 0x67, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x61, 0x70, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x24, 0x4c, 0x6f, 0x6e, 0x67, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x61, 0x70, 0x70, 0x5c, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x4c, 0x6f, 0x6e, 0x67, 0x62, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x61, 0x70, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_control_error_proto_rawDescOnce sync.Once
	file_control_error_proto_rawDescData = file_control_error_proto_rawDesc
)

func file_control_error_proto_rawDescGZIP() []byte {
	file_control_error_proto_rawDescOnce.Do(func() {
		file_control_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_control_error_proto_rawDescData)
	})
	return file_control_error_proto_rawDescData
}

var file_control_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_control_error_proto_goTypes = []interface{}{
	(*Error)(nil), // 0: longbridgeapp.control.v1.Error
}
var file_control_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_control_error_proto_init() }
func file_control_error_proto_init() {
	if File_control_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_control_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_control_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_control_error_proto_goTypes,
		DependencyIndexes: file_control_error_proto_depIdxs,
		MessageInfos:      file_control_error_proto_msgTypes,
	}.Build()
	File_control_error_proto = out.File
	file_control_error_proto_rawDesc = nil
	file_control_error_proto_goTypes = nil
	file_control_error_proto_depIdxs = nil
}
