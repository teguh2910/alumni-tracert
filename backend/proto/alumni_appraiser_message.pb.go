// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: alumni_appraiser_message.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AlumniAppraiser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId         uint64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Alumni         *Alumni `protobuf:"bytes,3,opt,name=alumni,proto3" json:"alumni,omitempty"`
	Name           string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Instansi       string  `protobuf:"bytes,5,opt,name=instansi,proto3" json:"instansi,omitempty"`
	Position       string  `protobuf:"bytes,6,opt,name=position,proto3" json:"position,omitempty"`
	AlumniPosition string  `protobuf:"bytes,7,opt,name=alumni_position,json=alumniPosition,proto3" json:"alumni_position,omitempty"`
	Created        string  `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty"`
	Updated        string  `protobuf:"bytes,9,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *AlumniAppraiser) Reset() {
	*x = AlumniAppraiser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alumni_appraiser_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlumniAppraiser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlumniAppraiser) ProtoMessage() {}

func (x *AlumniAppraiser) ProtoReflect() protoreflect.Message {
	mi := &file_alumni_appraiser_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlumniAppraiser.ProtoReflect.Descriptor instead.
func (*AlumniAppraiser) Descriptor() ([]byte, []int) {
	return file_alumni_appraiser_message_proto_rawDescGZIP(), []int{0}
}

func (x *AlumniAppraiser) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AlumniAppraiser) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AlumniAppraiser) GetAlumni() *Alumni {
	if x != nil {
		return x.Alumni
	}
	return nil
}

func (x *AlumniAppraiser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlumniAppraiser) GetInstansi() string {
	if x != nil {
		return x.Instansi
	}
	return ""
}

func (x *AlumniAppraiser) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *AlumniAppraiser) GetAlumniPosition() string {
	if x != nil {
		return x.AlumniPosition
	}
	return ""
}

func (x *AlumniAppraiser) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *AlumniAppraiser) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

type AlumniAppraiserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlumniAppraiser *AlumniAppraiser `protobuf:"bytes,1,opt,name=alumni_appraiser,json=alumniAppraiser,proto3" json:"alumni_appraiser,omitempty"`
	ListInput       *ListInput       `protobuf:"bytes,2,opt,name=list_input,json=listInput,proto3" json:"list_input,omitempty"`
}

func (x *AlumniAppraiserListResponse) Reset() {
	*x = AlumniAppraiserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alumni_appraiser_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlumniAppraiserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlumniAppraiserListResponse) ProtoMessage() {}

func (x *AlumniAppraiserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alumni_appraiser_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlumniAppraiserListResponse.ProtoReflect.Descriptor instead.
func (*AlumniAppraiserListResponse) Descriptor() ([]byte, []int) {
	return file_alumni_appraiser_message_proto_rawDescGZIP(), []int{1}
}

func (x *AlumniAppraiserListResponse) GetAlumniAppraiser() *AlumniAppraiser {
	if x != nil {
		return x.AlumniAppraiser
	}
	return nil
}

func (x *AlumniAppraiserListResponse) GetListInput() *ListInput {
	if x != nil {
		return x.ListInput
	}
	return nil
}

var File_alumni_appraiser_message_proto protoreflect.FileDescriptor

var file_alumni_appraiser_message_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x61, 0x69, 0x73,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x61, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x0f, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x41,
	0x70, 0x70, 0x72, 0x61, 0x69, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69,
	0x52, 0x06, 0x61, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x73, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x73, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x5f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61,
	0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x22, 0x91, 0x01, 0x0a, 0x1b, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x41, 0x70, 0x70, 0x72,
	0x61, 0x69, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x10, 0x61, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x5f, 0x61, 0x70, 0x70, 0x72,
	0x61, 0x69, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x41, 0x70, 0x70, 0x72, 0x61, 0x69,
	0x73, 0x65, 0x72, 0x52, 0x0f, 0x61, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x41, 0x70, 0x70, 0x72, 0x61,
	0x69, 0x73, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alumni_appraiser_message_proto_rawDescOnce sync.Once
	file_alumni_appraiser_message_proto_rawDescData = file_alumni_appraiser_message_proto_rawDesc
)

func file_alumni_appraiser_message_proto_rawDescGZIP() []byte {
	file_alumni_appraiser_message_proto_rawDescOnce.Do(func() {
		file_alumni_appraiser_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_alumni_appraiser_message_proto_rawDescData)
	})
	return file_alumni_appraiser_message_proto_rawDescData
}

var file_alumni_appraiser_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_alumni_appraiser_message_proto_goTypes = []interface{}{
	(*AlumniAppraiser)(nil),             // 0: proto.AlumniAppraiser
	(*AlumniAppraiserListResponse)(nil), // 1: proto.AlumniAppraiserListResponse
	(*Alumni)(nil),                      // 2: proto.Alumni
	(*ListInput)(nil),                   // 3: proto.ListInput
}
var file_alumni_appraiser_message_proto_depIdxs = []int32{
	2, // 0: proto.AlumniAppraiser.alumni:type_name -> proto.Alumni
	0, // 1: proto.AlumniAppraiserListResponse.alumni_appraiser:type_name -> proto.AlumniAppraiser
	3, // 2: proto.AlumniAppraiserListResponse.list_input:type_name -> proto.ListInput
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_alumni_appraiser_message_proto_init() }
func file_alumni_appraiser_message_proto_init() {
	if File_alumni_appraiser_message_proto != nil {
		return
	}
	file_alumni_message_proto_init()
	file_generic_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_alumni_appraiser_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlumniAppraiser); i {
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
		file_alumni_appraiser_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlumniAppraiserListResponse); i {
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
			RawDescriptor: file_alumni_appraiser_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alumni_appraiser_message_proto_goTypes,
		DependencyIndexes: file_alumni_appraiser_message_proto_depIdxs,
		MessageInfos:      file_alumni_appraiser_message_proto_msgTypes,
	}.Build()
	File_alumni_appraiser_message_proto = out.File
	file_alumni_appraiser_message_proto_rawDesc = nil
	file_alumni_appraiser_message_proto_goTypes = nil
	file_alumni_appraiser_message_proto_depIdxs = nil
}
