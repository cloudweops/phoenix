// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.5
// source: pb/example/test.proto

package example

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

type FOO int32

const (
	FOO_X FOO = 0
)

// Enum value maps for FOO.
var (
	FOO_name = map[int32]string{
		0: "X",
	}
	FOO_value = map[string]int32{
		"X": 0,
	}
)

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}

func (x FOO) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FOO) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_example_test_proto_enumTypes[0].Descriptor()
}

func (FOO) Type() protoreflect.EnumType {
	return &file_pb_example_test_proto_enumTypes[0]
}

func (x FOO) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FOO.Descriptor instead.
func (FOO) EnumDescriptor() ([]byte, []int) {
	return file_pb_example_test_proto_rawDescGZIP(), []int{0}
}

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Type  int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Reps  int64  `protobuf:"varint,3,opt,name=reps,proto3" json:"reps,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_example_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_pb_example_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_pb_example_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Test) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Test) GetReps() int64 {
	if x != nil {
		return x.Reps
	}
	return 0
}

var File_pb_example_test_proto protoreflect.FileDescriptor

var file_pb_example_test_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x57, 0x65,
	0x4f, 0x70, 0x73, 0x2e, 0x70, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x22, 0x44, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x65, 0x70, 0x73, 0x2a, 0x0c, 0x0a, 0x03, 0x46, 0x4f, 0x4f,
	0x12, 0x05, 0x0a, 0x01, 0x58, 0x10, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x57, 0x65, 0x4f, 0x70, 0x73,
	0x2f, 0x70, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_example_test_proto_rawDescOnce sync.Once
	file_pb_example_test_proto_rawDescData = file_pb_example_test_proto_rawDesc
)

func file_pb_example_test_proto_rawDescGZIP() []byte {
	file_pb_example_test_proto_rawDescOnce.Do(func() {
		file_pb_example_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_example_test_proto_rawDescData)
	})
	return file_pb_example_test_proto_rawDescData
}

var file_pb_example_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_example_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_example_test_proto_goTypes = []interface{}{
	(FOO)(0),     // 0: CloudWeOps.phoenix.example.FOO
	(*Test)(nil), // 1: CloudWeOps.phoenix.example.Test
}
var file_pb_example_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_example_test_proto_init() }
func file_pb_example_test_proto_init() {
	if File_pb_example_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_example_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
			RawDescriptor: file_pb_example_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_example_test_proto_goTypes,
		DependencyIndexes: file_pb_example_test_proto_depIdxs,
		EnumInfos:         file_pb_example_test_proto_enumTypes,
		MessageInfos:      file_pb_example_test_proto_msgTypes,
	}.Build()
	File_pb_example_test_proto = out.File
	file_pb_example_test_proto_rawDesc = nil
	file_pb_example_test_proto_goTypes = nil
	file_pb_example_test_proto_depIdxs = nil
}
