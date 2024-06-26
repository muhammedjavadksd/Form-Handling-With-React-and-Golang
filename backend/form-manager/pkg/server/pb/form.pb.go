// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: form-manager/pkg/server/pb/form.proto

package server

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

type FormReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName string `protobuf:"bytes,1,opt,name=FullName,proto3" json:"FullName,omitempty"`
	Phone    int64  `protobuf:"varint,2,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *FormReq) Reset() {
	*x = FormReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_manager_pkg_server_pb_form_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormReq) ProtoMessage() {}

func (x *FormReq) ProtoReflect() protoreflect.Message {
	mi := &file_form_manager_pkg_server_pb_form_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormReq.ProtoReflect.Descriptor instead.
func (*FormReq) Descriptor() ([]byte, []int) {
	return file_form_manager_pkg_server_pb_form_proto_rawDescGZIP(), []int{0}
}

func (x *FormReq) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *FormReq) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *FormReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *FormReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type FormResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *FormResponse) Reset() {
	*x = FormResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_form_manager_pkg_server_pb_form_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormResponse) ProtoMessage() {}

func (x *FormResponse) ProtoReflect() protoreflect.Message {
	mi := &file_form_manager_pkg_server_pb_form_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormResponse.ProtoReflect.Descriptor instead.
func (*FormResponse) Descriptor() ([]byte, []int) {
	return file_form_manager_pkg_server_pb_form_proto_rawDescGZIP(), []int{1}
}

func (x *FormResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_form_manager_pkg_server_pb_form_proto protoreflect.FileDescriptor

var file_form_manager_pkg_server_pb_form_proto_rawDesc = []byte{
	0x0a, 0x25, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x07, 0x46, 0x6f, 0x72, 0x6d, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x1e, 0x0a, 0x0c, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x64, 0x32, 0x2f, 0x0a, 0x04, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x27, 0x0a, 0x0a,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x08, 0x2e, 0x46, 0x6f, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x2d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_form_manager_pkg_server_pb_form_proto_rawDescOnce sync.Once
	file_form_manager_pkg_server_pb_form_proto_rawDescData = file_form_manager_pkg_server_pb_form_proto_rawDesc
)

func file_form_manager_pkg_server_pb_form_proto_rawDescGZIP() []byte {
	file_form_manager_pkg_server_pb_form_proto_rawDescOnce.Do(func() {
		file_form_manager_pkg_server_pb_form_proto_rawDescData = protoimpl.X.CompressGZIP(file_form_manager_pkg_server_pb_form_proto_rawDescData)
	})
	return file_form_manager_pkg_server_pb_form_proto_rawDescData
}

var file_form_manager_pkg_server_pb_form_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_form_manager_pkg_server_pb_form_proto_goTypes = []interface{}{
	(*FormReq)(nil),      // 0: FormReq
	(*FormResponse)(nil), // 1: FormResponse
}
var file_form_manager_pkg_server_pb_form_proto_depIdxs = []int32{
	0, // 0: Form.InsertForm:input_type -> FormReq
	1, // 1: Form.InsertForm:output_type -> FormResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_form_manager_pkg_server_pb_form_proto_init() }
func file_form_manager_pkg_server_pb_form_proto_init() {
	if File_form_manager_pkg_server_pb_form_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_form_manager_pkg_server_pb_form_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FormReq); i {
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
		file_form_manager_pkg_server_pb_form_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FormResponse); i {
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
			RawDescriptor: file_form_manager_pkg_server_pb_form_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_form_manager_pkg_server_pb_form_proto_goTypes,
		DependencyIndexes: file_form_manager_pkg_server_pb_form_proto_depIdxs,
		MessageInfos:      file_form_manager_pkg_server_pb_form_proto_msgTypes,
	}.Build()
	File_form_manager_pkg_server_pb_form_proto = out.File
	file_form_manager_pkg_server_pb_form_proto_rawDesc = nil
	file_form_manager_pkg_server_pb_form_proto_goTypes = nil
	file_form_manager_pkg_server_pb_form_proto_depIdxs = nil
}
