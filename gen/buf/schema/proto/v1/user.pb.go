// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: schema/proto/v1/user.proto

package psample

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

type Color int32

const (
	Color_COLOR_UNSPECIFIED Color = 0
	Color_COLOR_RED         Color = 1
	Color_COLOR_GREEN       Color = 2
	Color_COLOR_BLUE        Color = 3
)

// Enum value maps for Color.
var (
	Color_name = map[int32]string{
		0: "COLOR_UNSPECIFIED",
		1: "COLOR_RED",
		2: "COLOR_GREEN",
		3: "COLOR_BLUE",
	}
	Color_value = map[string]int32{
		"COLOR_UNSPECIFIED": 0,
		"COLOR_RED":         1,
		"COLOR_GREEN":       2,
		"COLOR_BLUE":        3,
	}
)

func (x Color) Enum() *Color {
	p := new(Color)
	*p = x
	return p
}

func (x Color) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Color) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_proto_v1_user_proto_enumTypes[0].Descriptor()
}

func (Color) Type() protoreflect.EnumType {
	return &file_schema_proto_v1_user_proto_enumTypes[0]
}

func (x Color) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Color.Descriptor instead.
func (Color) EnumDescriptor() ([]byte, []int) {
	return file_schema_proto_v1_user_proto_rawDescGZIP(), []int{0}
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_schema_proto_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Position) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pos       *Position `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
	Color     Color     `protobuf:"varint,3,opt,name=color,proto3,enum=schema.proto.v1.Color" json:"color,omitempty"`
	Inventory []*Item   `protobuf:"bytes,4,rep,name=inventory,proto3" json:"inventory,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_schema_proto_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPos() *Position {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *User) GetColor() Color {
	if x != nil {
		return x.Color
	}
	return Color_COLOR_UNSPECIFIED
}

func (x *User) GetInventory() []*Item {
	if x != nil {
		return x.Inventory
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pos  *Position `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_schema_proto_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetPos() *Position {
	if x != nil {
		return x.Pos
	}
	return nil
}

var File_schema_proto_v1_user_proto protoreflect.FileDescriptor

var file_schema_proto_v1_user_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x34, 0x0a,
	0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x01, 0x7a, 0x22, 0xaa, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2b, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x2c, 0x0a,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x09, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x22, 0x47, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x03,
	0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x2a, 0x4e, 0x0a, 0x05, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4c,
	0x4f, 0x52, 0x5f, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4c, 0x4f,
	0x52, 0x5f, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4c,
	0x4f, 0x52, 0x5f, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x03, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_proto_v1_user_proto_rawDescOnce sync.Once
	file_schema_proto_v1_user_proto_rawDescData = file_schema_proto_v1_user_proto_rawDesc
)

func file_schema_proto_v1_user_proto_rawDescGZIP() []byte {
	file_schema_proto_v1_user_proto_rawDescOnce.Do(func() {
		file_schema_proto_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_proto_v1_user_proto_rawDescData)
	})
	return file_schema_proto_v1_user_proto_rawDescData
}

var file_schema_proto_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_schema_proto_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schema_proto_v1_user_proto_goTypes = []interface{}{
	(Color)(0),       // 0: schema.proto.v1.Color
	(*Position)(nil), // 1: schema.proto.v1.Position
	(*User)(nil),     // 2: schema.proto.v1.User
	(*Item)(nil),     // 3: schema.proto.v1.Item
}
var file_schema_proto_v1_user_proto_depIdxs = []int32{
	1, // 0: schema.proto.v1.User.pos:type_name -> schema.proto.v1.Position
	0, // 1: schema.proto.v1.User.color:type_name -> schema.proto.v1.Color
	3, // 2: schema.proto.v1.User.inventory:type_name -> schema.proto.v1.Item
	1, // 3: schema.proto.v1.Item.pos:type_name -> schema.proto.v1.Position
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_schema_proto_v1_user_proto_init() }
func file_schema_proto_v1_user_proto_init() {
	if File_schema_proto_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_proto_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_schema_proto_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_schema_proto_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_schema_proto_v1_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_proto_v1_user_proto_goTypes,
		DependencyIndexes: file_schema_proto_v1_user_proto_depIdxs,
		EnumInfos:         file_schema_proto_v1_user_proto_enumTypes,
		MessageInfos:      file_schema_proto_v1_user_proto_msgTypes,
	}.Build()
	File_schema_proto_v1_user_proto = out.File
	file_schema_proto_v1_user_proto_rawDesc = nil
	file_schema_proto_v1_user_proto_goTypes = nil
	file_schema_proto_v1_user_proto_depIdxs = nil
}
