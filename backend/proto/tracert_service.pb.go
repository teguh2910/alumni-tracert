// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: tracert_service.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_tracert_service_proto protoreflect.FileDescriptor

var file_tracert_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x61, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x61, 0x69, 0x73, 0x65, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x61, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6c, 0x65, 0x67,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xce, 0x05, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x2e, 0x0a, 0x0c, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x1a,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0a, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x2b, 0x0a, 0x09, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x47, 0x65, 0x74, 0x12, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x1a, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x15,
	0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x41, 0x70, 0x70, 0x72, 0x61, 0x69, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c,
	0x75, 0x6d, 0x6e, 0x69, 0x41, 0x70, 0x70, 0x72, 0x61, 0x69, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x41, 0x70, 0x70, 0x72,
	0x61, 0x69, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x13, 0x41, 0x6c, 0x75, 0x6d, 0x6e,
	0x69, 0x41, 0x70, 0x70, 0x72, 0x61, 0x69, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x41,
	0x70, 0x70, 0x72, 0x61, 0x69, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x12, 0x41, 0x6c, 0x75, 0x6d,
	0x6e, 0x69, 0x41, 0x70, 0x70, 0x72, 0x61, 0x69, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x12, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x41, 0x70, 0x70,
	0x72, 0x61, 0x69, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x6c, 0x75, 0x6d, 0x6e, 0x69, 0x41, 0x70, 0x70, 0x72, 0x61, 0x69, 0x73, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x28, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x25, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74,
	0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x10,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tracert_service_proto_goTypes = []interface{}{
	(*LoginInput)(nil),                  // 0: proto.LoginInput
	(*QuestionGroupListInput)(nil),      // 1: proto.QuestionGroupListInput
	(*Alumni)(nil),                      // 2: proto.Alumni
	(*ListInput)(nil),                   // 3: proto.ListInput
	(*AlumniAppraiser)(nil),             // 4: proto.AlumniAppraiser
	(*User)(nil),                        // 5: proto.User
	(*UserAnswer)(nil),                  // 6: proto.UserAnswer
	(*QuestionGroupList)(nil),           // 7: proto.QuestionGroupList
	(*AlumniListResponse)(nil),          // 8: proto.AlumniListResponse
	(*AlumniAppraiserListResponse)(nil), // 9: proto.AlumniAppraiserListResponse
	(*UserListResponse)(nil),            // 10: proto.UserListResponse
}
var file_tracert_service_proto_depIdxs = []int32{
	0,  // 0: proto.TracertService.Login:input_type -> proto.LoginInput
	1,  // 1: proto.TracertService.QuestionList:input_type -> proto.QuestionGroupListInput
	2,  // 2: proto.TracertService.AlumniCreate:input_type -> proto.Alumni
	3,  // 3: proto.TracertService.AlumniList:input_type -> proto.ListInput
	2,  // 4: proto.TracertService.AlumniGet:input_type -> proto.Alumni
	4,  // 5: proto.TracertService.AlumniAppraiserCreate:input_type -> proto.AlumniAppraiser
	3,  // 6: proto.TracertService.AlumniAppraiserList:input_type -> proto.ListInput
	4,  // 7: proto.TracertService.AlumniAppraiserGet:input_type -> proto.AlumniAppraiser
	5,  // 8: proto.TracertService.UserCreate:input_type -> proto.User
	3,  // 9: proto.TracertService.UserList:input_type -> proto.ListInput
	5,  // 10: proto.TracertService.UserGet:input_type -> proto.User
	6,  // 11: proto.TracertService.UserAnswerCreate:input_type -> proto.UserAnswer
	5,  // 12: proto.TracertService.Login:output_type -> proto.User
	7,  // 13: proto.TracertService.QuestionList:output_type -> proto.QuestionGroupList
	2,  // 14: proto.TracertService.AlumniCreate:output_type -> proto.Alumni
	8,  // 15: proto.TracertService.AlumniList:output_type -> proto.AlumniListResponse
	2,  // 16: proto.TracertService.AlumniGet:output_type -> proto.Alumni
	4,  // 17: proto.TracertService.AlumniAppraiserCreate:output_type -> proto.AlumniAppraiser
	9,  // 18: proto.TracertService.AlumniAppraiserList:output_type -> proto.AlumniAppraiserListResponse
	4,  // 19: proto.TracertService.AlumniAppraiserGet:output_type -> proto.AlumniAppraiser
	5,  // 20: proto.TracertService.UserCreate:output_type -> proto.User
	10, // 21: proto.TracertService.UserList:output_type -> proto.UserListResponse
	5,  // 22: proto.TracertService.UserGet:output_type -> proto.User
	6,  // 23: proto.TracertService.UserAnswerCreate:output_type -> proto.UserAnswer
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_tracert_service_proto_init() }
func file_tracert_service_proto_init() {
	if File_tracert_service_proto != nil {
		return
	}
	file_alumni_appraiser_message_proto_init()
	file_alumni_message_proto_init()
	file_generic_message_proto_init()
	file_legalize_message_proto_init()
	file_question_group_message_proto_init()
	file_question_message_proto_init()
	file_question_option_message_proto_init()
	file_user_answer_message_proto_init()
	file_user_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tracert_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tracert_service_proto_goTypes,
		DependencyIndexes: file_tracert_service_proto_depIdxs,
	}.Build()
	File_tracert_service_proto = out.File
	file_tracert_service_proto_rawDesc = nil
	file_tracert_service_proto_goTypes = nil
	file_tracert_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TracertServiceClient is the client API for TracertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TracertServiceClient interface {
	Login(ctx context.Context, in *LoginInput, opts ...grpc.CallOption) (*User, error)
	QuestionList(ctx context.Context, in *QuestionGroupListInput, opts ...grpc.CallOption) (*QuestionGroupList, error)
	AlumniCreate(ctx context.Context, in *Alumni, opts ...grpc.CallOption) (*Alumni, error)
	AlumniList(ctx context.Context, in *ListInput, opts ...grpc.CallOption) (TracertService_AlumniListClient, error)
	AlumniGet(ctx context.Context, in *Alumni, opts ...grpc.CallOption) (*Alumni, error)
	AlumniAppraiserCreate(ctx context.Context, in *AlumniAppraiser, opts ...grpc.CallOption) (*AlumniAppraiser, error)
	AlumniAppraiserList(ctx context.Context, in *ListInput, opts ...grpc.CallOption) (TracertService_AlumniAppraiserListClient, error)
	AlumniAppraiserGet(ctx context.Context, in *AlumniAppraiser, opts ...grpc.CallOption) (*AlumniAppraiser, error)
	// rpc AlumniAppraiserUpdate(AlumniAppraiser) returns (AlumniAppraiser) {}
	//
	UserCreate(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UserList(ctx context.Context, in *ListInput, opts ...grpc.CallOption) (TracertService_UserListClient, error)
	UserGet(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UserAnswerCreate(ctx context.Context, in *UserAnswer, opts ...grpc.CallOption) (*UserAnswer, error)
}

type tracertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTracertServiceClient(cc grpc.ClientConnInterface) TracertServiceClient {
	return &tracertServiceClient{cc}
}

func (c *tracertServiceClient) Login(ctx context.Context, in *LoginInput, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.TracertService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracertServiceClient) QuestionList(ctx context.Context, in *QuestionGroupListInput, opts ...grpc.CallOption) (*QuestionGroupList, error) {
	out := new(QuestionGroupList)
	err := c.cc.Invoke(ctx, "/proto.TracertService/QuestionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracertServiceClient) AlumniCreate(ctx context.Context, in *Alumni, opts ...grpc.CallOption) (*Alumni, error) {
	out := new(Alumni)
	err := c.cc.Invoke(ctx, "/proto.TracertService/AlumniCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracertServiceClient) AlumniList(ctx context.Context, in *ListInput, opts ...grpc.CallOption) (TracertService_AlumniListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TracertService_serviceDesc.Streams[0], "/proto.TracertService/AlumniList", opts...)
	if err != nil {
		return nil, err
	}
	x := &tracertServiceAlumniListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TracertService_AlumniListClient interface {
	Recv() (*AlumniListResponse, error)
	grpc.ClientStream
}

type tracertServiceAlumniListClient struct {
	grpc.ClientStream
}

func (x *tracertServiceAlumniListClient) Recv() (*AlumniListResponse, error) {
	m := new(AlumniListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tracertServiceClient) AlumniGet(ctx context.Context, in *Alumni, opts ...grpc.CallOption) (*Alumni, error) {
	out := new(Alumni)
	err := c.cc.Invoke(ctx, "/proto.TracertService/AlumniGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracertServiceClient) AlumniAppraiserCreate(ctx context.Context, in *AlumniAppraiser, opts ...grpc.CallOption) (*AlumniAppraiser, error) {
	out := new(AlumniAppraiser)
	err := c.cc.Invoke(ctx, "/proto.TracertService/AlumniAppraiserCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracertServiceClient) AlumniAppraiserList(ctx context.Context, in *ListInput, opts ...grpc.CallOption) (TracertService_AlumniAppraiserListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TracertService_serviceDesc.Streams[1], "/proto.TracertService/AlumniAppraiserList", opts...)
	if err != nil {
		return nil, err
	}
	x := &tracertServiceAlumniAppraiserListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TracertService_AlumniAppraiserListClient interface {
	Recv() (*AlumniAppraiserListResponse, error)
	grpc.ClientStream
}

type tracertServiceAlumniAppraiserListClient struct {
	grpc.ClientStream
}

func (x *tracertServiceAlumniAppraiserListClient) Recv() (*AlumniAppraiserListResponse, error) {
	m := new(AlumniAppraiserListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tracertServiceClient) AlumniAppraiserGet(ctx context.Context, in *AlumniAppraiser, opts ...grpc.CallOption) (*AlumniAppraiser, error) {
	out := new(AlumniAppraiser)
	err := c.cc.Invoke(ctx, "/proto.TracertService/AlumniAppraiserGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracertServiceClient) UserCreate(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.TracertService/UserCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracertServiceClient) UserList(ctx context.Context, in *ListInput, opts ...grpc.CallOption) (TracertService_UserListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TracertService_serviceDesc.Streams[2], "/proto.TracertService/UserList", opts...)
	if err != nil {
		return nil, err
	}
	x := &tracertServiceUserListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TracertService_UserListClient interface {
	Recv() (*UserListResponse, error)
	grpc.ClientStream
}

type tracertServiceUserListClient struct {
	grpc.ClientStream
}

func (x *tracertServiceUserListClient) Recv() (*UserListResponse, error) {
	m := new(UserListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tracertServiceClient) UserGet(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.TracertService/UserGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracertServiceClient) UserAnswerCreate(ctx context.Context, in *UserAnswer, opts ...grpc.CallOption) (*UserAnswer, error) {
	out := new(UserAnswer)
	err := c.cc.Invoke(ctx, "/proto.TracertService/UserAnswerCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TracertServiceServer is the server API for TracertService service.
type TracertServiceServer interface {
	Login(context.Context, *LoginInput) (*User, error)
	QuestionList(context.Context, *QuestionGroupListInput) (*QuestionGroupList, error)
	AlumniCreate(context.Context, *Alumni) (*Alumni, error)
	AlumniList(*ListInput, TracertService_AlumniListServer) error
	AlumniGet(context.Context, *Alumni) (*Alumni, error)
	AlumniAppraiserCreate(context.Context, *AlumniAppraiser) (*AlumniAppraiser, error)
	AlumniAppraiserList(*ListInput, TracertService_AlumniAppraiserListServer) error
	AlumniAppraiserGet(context.Context, *AlumniAppraiser) (*AlumniAppraiser, error)
	// rpc AlumniAppraiserUpdate(AlumniAppraiser) returns (AlumniAppraiser) {}
	//
	UserCreate(context.Context, *User) (*User, error)
	UserList(*ListInput, TracertService_UserListServer) error
	UserGet(context.Context, *User) (*User, error)
	UserAnswerCreate(context.Context, *UserAnswer) (*UserAnswer, error)
}

// UnimplementedTracertServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTracertServiceServer struct {
}

func (*UnimplementedTracertServiceServer) Login(context.Context, *LoginInput) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedTracertServiceServer) QuestionList(context.Context, *QuestionGroupListInput) (*QuestionGroupList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuestionList not implemented")
}
func (*UnimplementedTracertServiceServer) AlumniCreate(context.Context, *Alumni) (*Alumni, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlumniCreate not implemented")
}
func (*UnimplementedTracertServiceServer) AlumniList(*ListInput, TracertService_AlumniListServer) error {
	return status.Errorf(codes.Unimplemented, "method AlumniList not implemented")
}
func (*UnimplementedTracertServiceServer) AlumniGet(context.Context, *Alumni) (*Alumni, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlumniGet not implemented")
}
func (*UnimplementedTracertServiceServer) AlumniAppraiserCreate(context.Context, *AlumniAppraiser) (*AlumniAppraiser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlumniAppraiserCreate not implemented")
}
func (*UnimplementedTracertServiceServer) AlumniAppraiserList(*ListInput, TracertService_AlumniAppraiserListServer) error {
	return status.Errorf(codes.Unimplemented, "method AlumniAppraiserList not implemented")
}
func (*UnimplementedTracertServiceServer) AlumniAppraiserGet(context.Context, *AlumniAppraiser) (*AlumniAppraiser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlumniAppraiserGet not implemented")
}
func (*UnimplementedTracertServiceServer) UserCreate(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (*UnimplementedTracertServiceServer) UserList(*ListInput, TracertService_UserListServer) error {
	return status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (*UnimplementedTracertServiceServer) UserGet(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGet not implemented")
}
func (*UnimplementedTracertServiceServer) UserAnswerCreate(context.Context, *UserAnswer) (*UserAnswer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAnswerCreate not implemented")
}

func RegisterTracertServiceServer(s *grpc.Server, srv TracertServiceServer) {
	s.RegisterService(&_TracertService_serviceDesc, srv)
}

func _TracertService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracertServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TracertService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracertServiceServer).Login(ctx, req.(*LoginInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TracertService_QuestionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionGroupListInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracertServiceServer).QuestionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TracertService/QuestionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracertServiceServer).QuestionList(ctx, req.(*QuestionGroupListInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _TracertService_AlumniCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Alumni)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracertServiceServer).AlumniCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TracertService/AlumniCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracertServiceServer).AlumniCreate(ctx, req.(*Alumni))
	}
	return interceptor(ctx, in, info, handler)
}

func _TracertService_AlumniList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TracertServiceServer).AlumniList(m, &tracertServiceAlumniListServer{stream})
}

type TracertService_AlumniListServer interface {
	Send(*AlumniListResponse) error
	grpc.ServerStream
}

type tracertServiceAlumniListServer struct {
	grpc.ServerStream
}

func (x *tracertServiceAlumniListServer) Send(m *AlumniListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TracertService_AlumniGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Alumni)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracertServiceServer).AlumniGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TracertService/AlumniGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracertServiceServer).AlumniGet(ctx, req.(*Alumni))
	}
	return interceptor(ctx, in, info, handler)
}

func _TracertService_AlumniAppraiserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlumniAppraiser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracertServiceServer).AlumniAppraiserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TracertService/AlumniAppraiserCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracertServiceServer).AlumniAppraiserCreate(ctx, req.(*AlumniAppraiser))
	}
	return interceptor(ctx, in, info, handler)
}

func _TracertService_AlumniAppraiserList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TracertServiceServer).AlumniAppraiserList(m, &tracertServiceAlumniAppraiserListServer{stream})
}

type TracertService_AlumniAppraiserListServer interface {
	Send(*AlumniAppraiserListResponse) error
	grpc.ServerStream
}

type tracertServiceAlumniAppraiserListServer struct {
	grpc.ServerStream
}

func (x *tracertServiceAlumniAppraiserListServer) Send(m *AlumniAppraiserListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TracertService_AlumniAppraiserGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlumniAppraiser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracertServiceServer).AlumniAppraiserGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TracertService/AlumniAppraiserGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracertServiceServer).AlumniAppraiserGet(ctx, req.(*AlumniAppraiser))
	}
	return interceptor(ctx, in, info, handler)
}

func _TracertService_UserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracertServiceServer).UserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TracertService/UserCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracertServiceServer).UserCreate(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _TracertService_UserList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TracertServiceServer).UserList(m, &tracertServiceUserListServer{stream})
}

type TracertService_UserListServer interface {
	Send(*UserListResponse) error
	grpc.ServerStream
}

type tracertServiceUserListServer struct {
	grpc.ServerStream
}

func (x *tracertServiceUserListServer) Send(m *UserListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TracertService_UserGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracertServiceServer).UserGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TracertService/UserGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracertServiceServer).UserGet(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _TracertService_UserAnswerCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAnswer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracertServiceServer).UserAnswerCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TracertService/UserAnswerCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracertServiceServer).UserAnswerCreate(ctx, req.(*UserAnswer))
	}
	return interceptor(ctx, in, info, handler)
}

var _TracertService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TracertService",
	HandlerType: (*TracertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _TracertService_Login_Handler,
		},
		{
			MethodName: "QuestionList",
			Handler:    _TracertService_QuestionList_Handler,
		},
		{
			MethodName: "AlumniCreate",
			Handler:    _TracertService_AlumniCreate_Handler,
		},
		{
			MethodName: "AlumniGet",
			Handler:    _TracertService_AlumniGet_Handler,
		},
		{
			MethodName: "AlumniAppraiserCreate",
			Handler:    _TracertService_AlumniAppraiserCreate_Handler,
		},
		{
			MethodName: "AlumniAppraiserGet",
			Handler:    _TracertService_AlumniAppraiserGet_Handler,
		},
		{
			MethodName: "UserCreate",
			Handler:    _TracertService_UserCreate_Handler,
		},
		{
			MethodName: "UserGet",
			Handler:    _TracertService_UserGet_Handler,
		},
		{
			MethodName: "UserAnswerCreate",
			Handler:    _TracertService_UserAnswerCreate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AlumniList",
			Handler:       _TracertService_AlumniList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AlumniAppraiserList",
			Handler:       _TracertService_AlumniAppraiserList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UserList",
			Handler:       _TracertService_UserList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tracert_service.proto",
}
