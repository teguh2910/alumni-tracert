// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: certificate_message.proto

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

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AlumniId       uint64    `protobuf:"varint,2,opt,name=alumni_id,json=alumniId,proto3" json:"alumni_id,omitempty"`
	Nim            string    `protobuf:"bytes,3,opt,name=nim,proto3" json:"nim,omitempty"`
	MajorStudy     string    `protobuf:"bytes,4,opt,name=major_study,json=majorStudy,proto3" json:"major_study,omitempty"`
	EntryYear      uint32    `protobuf:"varint,5,opt,name=entry_year,json=entryYear,proto3" json:"entry_year,omitempty"`
	GraduationYear uint32    `protobuf:"varint,6,opt,name=graduation_year,json=graduationYear,proto3" json:"graduation_year,omitempty"`
	NoIjazah       string    `protobuf:"bytes,7,opt,name=no_ijazah,json=noIjazah,proto3" json:"no_ijazah,omitempty"`
	Created        string    `protobuf:"bytes,8,opt,name=created,proto3" json:"created,omitempty"`
	Updated        string    `protobuf:"bytes,9,opt,name=updated,proto3" json:"updated,omitempty"`
	Legalize       *Legalize `protobuf:"bytes,10,opt,name=legalize,proto3" json:"legalize,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_certificate_message_proto_rawDescGZIP(), []int{0}
}

func (x *Certificate) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Certificate) GetAlumniId() uint64 {
	if x != nil {
		return x.AlumniId
	}
	return 0
}

func (x *Certificate) GetNim() string {
	if x != nil {
		return x.Nim
	}
	return ""
}

func (x *Certificate) GetMajorStudy() string {
	if x != nil {
		return x.MajorStudy
	}
	return ""
}

func (x *Certificate) GetEntryYear() uint32 {
	if x != nil {
		return x.EntryYear
	}
	return 0
}

func (x *Certificate) GetGraduationYear() uint32 {
	if x != nil {
		return x.GraduationYear
	}
	return 0
}

func (x *Certificate) GetNoIjazah() string {
	if x != nil {
		return x.NoIjazah
	}
	return ""
}

func (x *Certificate) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Certificate) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *Certificate) GetLegalize() *Legalize {
	if x != nil {
		return x.Legalize
	}
	return nil
}

type Certificates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate []*Certificate `protobuf:"bytes,1,rep,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *Certificates) Reset() {
	*x = Certificates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificates) ProtoMessage() {}

func (x *Certificates) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificates.ProtoReflect.Descriptor instead.
func (*Certificates) Descriptor() ([]byte, []int) {
	return file_certificate_message_proto_rawDescGZIP(), []int{1}
}

func (x *Certificates) GetCertificate() []*Certificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

var File_certificate_message_proto protoreflect.FileDescriptor

var file_certificate_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x0b, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c,
	0x75, 0x6d, 0x6e, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61,
	0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x69, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x6a,
	0x6f, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x53, 0x74, 0x75, 0x64, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x59, 0x65, 0x61, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x67, 0x72, 0x61,
	0x64, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x67, 0x72, 0x61, 0x64, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65,
	0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x5f, 0x69, 0x6a, 0x61, 0x7a, 0x61, 0x68, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x49, 0x6a, 0x61, 0x7a, 0x61, 0x68, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x65,
	0x67, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x08, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x22, 0x44, 0x0a, 0x0c, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x34, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_certificate_message_proto_rawDescOnce sync.Once
	file_certificate_message_proto_rawDescData = file_certificate_message_proto_rawDesc
)

func file_certificate_message_proto_rawDescGZIP() []byte {
	file_certificate_message_proto_rawDescOnce.Do(func() {
		file_certificate_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_certificate_message_proto_rawDescData)
	})
	return file_certificate_message_proto_rawDescData
}

var file_certificate_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_certificate_message_proto_goTypes = []interface{}{
	(*Certificate)(nil),  // 0: proto.Certificate
	(*Certificates)(nil), // 1: proto.Certificates
	(*Legalize)(nil),     // 2: proto.Legalize
}
var file_certificate_message_proto_depIdxs = []int32{
	2, // 0: proto.Certificate.legalize:type_name -> proto.Legalize
	0, // 1: proto.Certificates.certificate:type_name -> proto.Certificate
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_certificate_message_proto_init() }
func file_certificate_message_proto_init() {
	if File_certificate_message_proto != nil {
		return
	}
	file_legalize_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_certificate_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
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
		file_certificate_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificates); i {
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
			RawDescriptor: file_certificate_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_certificate_message_proto_goTypes,
		DependencyIndexes: file_certificate_message_proto_depIdxs,
		MessageInfos:      file_certificate_message_proto_msgTypes,
	}.Build()
	File_certificate_message_proto = out.File
	file_certificate_message_proto_rawDesc = nil
	file_certificate_message_proto_goTypes = nil
	file_certificate_message_proto_depIdxs = nil
}
