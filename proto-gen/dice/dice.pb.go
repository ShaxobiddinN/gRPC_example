// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protos/dice.proto

package dice

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_dice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_dice_proto_rawDescGZIP(), []int{0}
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_dice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_protos_dice_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The request message containing coming request
type RollDiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *RollDiceRequest) Reset() {
	*x = RollDiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_dice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollDiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollDiceRequest) ProtoMessage() {}

func (x *RollDiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollDiceRequest.ProtoReflect.Descriptor instead.
func (*RollDiceRequest) Descriptor() ([]byte, []int) {
	return file_protos_dice_proto_rawDescGZIP(), []int{2}
}

func (x *RollDiceRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

// The response message containing response data
type RollDiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dice []int32 `protobuf:"varint,1,rep,packed,name=dice,proto3" json:"dice,omitempty"`
}

func (x *RollDiceResponse) Reset() {
	*x = RollDiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_dice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollDiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollDiceResponse) ProtoMessage() {}

func (x *RollDiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollDiceResponse.ProtoReflect.Descriptor instead.
func (*RollDiceResponse) Descriptor() ([]byte, []int) {
	return file_protos_dice_proto_rawDescGZIP(), []int{3}
}

func (x *RollDiceResponse) GetDice() []int32 {
	if x != nil {
		return x.Dice
	}
	return nil
}

var File_protos_dice_proto protoreflect.FileDescriptor

var file_protos_dice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x04,
	0x50, 0x6f, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x23,
	0x0a, 0x0f, 0x52, 0x6f, 0x6c, 0x6c, 0x44, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6e, 0x75, 0x6d, 0x22, 0x26, 0x0a, 0x10, 0x52, 0x6f, 0x6c, 0x6c, 0x44, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x64, 0x69, 0x63, 0x65, 0x32, 0x56, 0x0a, 0x08, 0x54,
	0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x17, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x05, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00,
	0x12, 0x31, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x6c, 0x44, 0x69, 0x63, 0x65, 0x12, 0x10, 0x2e, 0x52,
	0x6f, 0x6c, 0x6c, 0x44, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x44, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x64, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_dice_proto_rawDescOnce sync.Once
	file_protos_dice_proto_rawDescData = file_protos_dice_proto_rawDesc
)

func file_protos_dice_proto_rawDescGZIP() []byte {
	file_protos_dice_proto_rawDescOnce.Do(func() {
		file_protos_dice_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_dice_proto_rawDescData)
	})
	return file_protos_dice_proto_rawDescData
}

var file_protos_dice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_dice_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: Empty
	(*Pong)(nil),             // 1: Pong
	(*RollDiceRequest)(nil),  // 2: RollDiceRequest
	(*RollDiceResponse)(nil), // 3: RollDiceResponse
}
var file_protos_dice_proto_depIdxs = []int32{
	0, // 0: Tutorial.Ping:input_type -> Empty
	2, // 1: Tutorial.RollDice:input_type -> RollDiceRequest
	1, // 2: Tutorial.Ping:output_type -> Pong
	3, // 3: Tutorial.RollDice:output_type -> RollDiceResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_dice_proto_init() }
func file_protos_dice_proto_init() {
	if File_protos_dice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_dice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_protos_dice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
		file_protos_dice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollDiceRequest); i {
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
		file_protos_dice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollDiceResponse); i {
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
			RawDescriptor: file_protos_dice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_dice_proto_goTypes,
		DependencyIndexes: file_protos_dice_proto_depIdxs,
		MessageInfos:      file_protos_dice_proto_msgTypes,
	}.Build()
	File_protos_dice_proto = out.File
	file_protos_dice_proto_rawDesc = nil
	file_protos_dice_proto_goTypes = nil
	file_protos_dice_proto_depIdxs = nil
}
