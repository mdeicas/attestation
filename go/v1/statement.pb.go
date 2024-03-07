// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.24.4
// source: in_toto_attestation/v1/statement.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Proto representation of the in-toto v1 Statement.
// https://github.com/in-toto/attestation/tree/main/spec/v1
// Validation of all fields is left to the users of this proto.
type Statement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Expected to always be "https://in-toto.io/Statement/v1"
	Type          string                `protobuf:"bytes,1,opt,name=type,json=_type,proto3" json:"type,omitempty"`
	Subject       []*ResourceDescriptor `protobuf:"bytes,2,rep,name=subject,proto3" json:"subject,omitempty"`
	PredicateType string                `protobuf:"bytes,3,opt,name=predicate_type,json=predicateType,proto3" json:"predicate_type,omitempty"`
	Predicate     *structpb.Struct      `protobuf:"bytes,4,opt,name=predicate,proto3" json:"predicate,omitempty"`
}

func (x *Statement) Reset() {
	*x = Statement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_toto_attestation_v1_statement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_v1_statement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement.ProtoReflect.Descriptor instead.
func (*Statement) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_v1_statement_proto_rawDescGZIP(), []int{0}
}

func (x *Statement) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Statement) GetSubject() []*ResourceDescriptor {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *Statement) GetPredicateType() string {
	if x != nil {
		return x.PredicateType
	}
	return ""
}

func (x *Statement) GetPredicate() *structpb.Struct {
	if x != nil {
		return x.Predicate
	}
	return nil
}

var File_in_toto_attestation_v1_statement_proto protoreflect.FileDescriptor

var file_in_toto_attestation_v1_statement_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74,
	0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30,
	0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc4, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x13,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x35, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x09, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x47, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x2d, 0x74, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_in_toto_attestation_v1_statement_proto_rawDescOnce sync.Once
	file_in_toto_attestation_v1_statement_proto_rawDescData = file_in_toto_attestation_v1_statement_proto_rawDesc
)

func file_in_toto_attestation_v1_statement_proto_rawDescGZIP() []byte {
	file_in_toto_attestation_v1_statement_proto_rawDescOnce.Do(func() {
		file_in_toto_attestation_v1_statement_proto_rawDescData = protoimpl.X.CompressGZIP(file_in_toto_attestation_v1_statement_proto_rawDescData)
	})
	return file_in_toto_attestation_v1_statement_proto_rawDescData
}

var file_in_toto_attestation_v1_statement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_in_toto_attestation_v1_statement_proto_goTypes = []interface{}{
	(*Statement)(nil),          // 0: in_toto_attestation.v1.Statement
	(*ResourceDescriptor)(nil), // 1: in_toto_attestation.v1.ResourceDescriptor
	(*structpb.Struct)(nil),    // 2: google.protobuf.Struct
}
var file_in_toto_attestation_v1_statement_proto_depIdxs = []int32{
	1, // 0: in_toto_attestation.v1.Statement.subject:type_name -> in_toto_attestation.v1.ResourceDescriptor
	2, // 1: in_toto_attestation.v1.Statement.predicate:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_in_toto_attestation_v1_statement_proto_init() }
func file_in_toto_attestation_v1_statement_proto_init() {
	if File_in_toto_attestation_v1_statement_proto != nil {
		return
	}
	file_in_toto_attestation_v1_resource_descriptor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_in_toto_attestation_v1_statement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statement); i {
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
			RawDescriptor: file_in_toto_attestation_v1_statement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_in_toto_attestation_v1_statement_proto_goTypes,
		DependencyIndexes: file_in_toto_attestation_v1_statement_proto_depIdxs,
		MessageInfos:      file_in_toto_attestation_v1_statement_proto_msgTypes,
	}.Build()
	File_in_toto_attestation_v1_statement_proto = out.File
	file_in_toto_attestation_v1_statement_proto_rawDesc = nil
	file_in_toto_attestation_v1_statement_proto_goTypes = nil
	file_in_toto_attestation_v1_statement_proto_depIdxs = nil
}
