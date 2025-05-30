// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: types/types.proto

package types

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

type UUIDValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUIDValue) Reset() {
	*x = UUIDValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDValue) ProtoMessage() {}

func (x *UUIDValue) ProtoReflect() protoreflect.Message {
	mi := &file_types_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDValue.ProtoReflect.Descriptor instead.
func (*UUIDValue) Descriptor() ([]byte, []int) {
	return file_types_types_proto_rawDescGZIP(), []int{0}
}

func (x *UUIDValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type JSONValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *JSONValue) Reset() {
	*x = JSONValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JSONValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSONValue) ProtoMessage() {}

func (x *JSONValue) ProtoReflect() protoreflect.Message {
	mi := &file_types_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSONValue.ProtoReflect.Descriptor instead.
func (*JSONValue) Descriptor() ([]byte, []int) {
	return file_types_types_proto_rawDescGZIP(), []int{1}
}

func (x *JSONValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_types_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_types_types_proto_rawDescGZIP(), []int{2}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type InetValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InetValue) Reset() {
	*x = InetValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InetValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InetValue) ProtoMessage() {}

func (x *InetValue) ProtoReflect() protoreflect.Message {
	mi := &file_types_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InetValue.ProtoReflect.Descriptor instead.
func (*InetValue) Descriptor() ([]byte, []int) {
	return file_types_types_proto_rawDescGZIP(), []int{3}
}

func (x *InetValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TimeOnly struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TimeOnly) Reset() {
	*x = TimeOnly{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeOnly) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeOnly) ProtoMessage() {}

func (x *TimeOnly) ProtoReflect() protoreflect.Message {
	mi := &file_types_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeOnly.ProtoReflect.Descriptor instead.
func (*TimeOnly) Descriptor() ([]byte, []int) {
	return file_types_types_proto_rawDescGZIP(), []int{4}
}

func (x *TimeOnly) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type BigInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BigInt) Reset() {
	*x = BigInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigInt) ProtoMessage() {}

func (x *BigInt) ProtoReflect() protoreflect.Message {
	mi := &file_types_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigInt.ProtoReflect.Descriptor instead.
func (*BigInt) Descriptor() ([]byte, []int) {
	return file_types_types_proto_rawDescGZIP(), []int{5}
}

func (x *BigInt) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_types_types_proto protoreflect.FileDescriptor

var file_types_types_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22,
	0x21, 0x0a, 0x09, 0x55, 0x55, 0x49, 0x44, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x4a, 0x53, 0x4f, 0x4e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1c, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x49, 0x6e, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x6e,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x42, 0x69, 0x67, 0x49,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x68, 0x6b, 0x61, 0x72, 0x53, 0x61, 0x6e,
	0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_types_proto_rawDescOnce sync.Once
	file_types_types_proto_rawDescData = file_types_types_proto_rawDesc
)

func file_types_types_proto_rawDescGZIP() []byte {
	file_types_types_proto_rawDescOnce.Do(func() {
		file_types_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_types_proto_rawDescData)
	})
	return file_types_types_proto_rawDescData
}

var file_types_types_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_types_types_proto_goTypes = []interface{}{
	(*UUIDValue)(nil), // 0: gorm.types.UUIDValue
	(*JSONValue)(nil), // 1: gorm.types.JSONValue
	(*UUID)(nil),      // 2: gorm.types.UUID
	(*InetValue)(nil), // 3: gorm.types.InetValue
	(*TimeOnly)(nil),  // 4: gorm.types.TimeOnly
	(*BigInt)(nil),    // 5: gorm.types.BigInt
}
var file_types_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_types_proto_init() }
func file_types_types_proto_init() {
	if File_types_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDValue); i {
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
		file_types_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JSONValue); i {
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
		file_types_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_types_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InetValue); i {
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
		file_types_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeOnly); i {
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
		file_types_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigInt); i {
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
			RawDescriptor: file_types_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_types_proto_goTypes,
		DependencyIndexes: file_types_types_proto_depIdxs,
		MessageInfos:      file_types_types_proto_msgTypes,
	}.Build()
	File_types_types_proto = out.File
	file_types_types_proto_rawDesc = nil
	file_types_types_proto_goTypes = nil
	file_types_types_proto_depIdxs = nil
}
