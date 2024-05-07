// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: in_toto_attestation/predicates/vsa/v0/vsa.proto

package v0

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Proto representation of predicate type https://slsa.dev/verification_summary/v0.2
// Validation of all fields is left to the users of this proto.
type VerificationSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verifier           *VerificationSummary_Verifier           `protobuf:"bytes,1,opt,name=verifier,proto3" json:"verifier,omitempty"`
	TimeVerified       *timestamppb.Timestamp                  `protobuf:"bytes,2,opt,name=time_verified,proto3" json:"time_verified,omitempty"`
	ResourceUri        string                                  `protobuf:"bytes,3,opt,name=resource_uri,proto3" json:"resource_uri,omitempty"`
	Policy             *VerificationSummary_Policy             `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	InputAttestations  []*VerificationSummary_InputAttestation `protobuf:"bytes,5,rep,name=input_attestations,proto3" json:"input_attestations,omitempty"`
	VerificationResult string                                  `protobuf:"bytes,6,opt,name=verification_result,proto3" json:"verification_result,omitempty"`
	PolicyLevel        string                                  `protobuf:"bytes,7,opt,name=policy_level,proto3" json:"policy_level,omitempty"`
	DependencyLevels   map[string]uint64                       `protobuf:"bytes,8,rep,name=dependency_levels,proto3" json:"dependency_levels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *VerificationSummary) Reset() {
	*x = VerificationSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationSummary) ProtoMessage() {}

func (x *VerificationSummary) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationSummary.ProtoReflect.Descriptor instead.
func (*VerificationSummary) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDescGZIP(), []int{0}
}

func (x *VerificationSummary) GetVerifier() *VerificationSummary_Verifier {
	if x != nil {
		return x.Verifier
	}
	return nil
}

func (x *VerificationSummary) GetTimeVerified() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeVerified
	}
	return nil
}

func (x *VerificationSummary) GetResourceUri() string {
	if x != nil {
		return x.ResourceUri
	}
	return ""
}

func (x *VerificationSummary) GetPolicy() *VerificationSummary_Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

func (x *VerificationSummary) GetInputAttestations() []*VerificationSummary_InputAttestation {
	if x != nil {
		return x.InputAttestations
	}
	return nil
}

func (x *VerificationSummary) GetVerificationResult() string {
	if x != nil {
		return x.VerificationResult
	}
	return ""
}

func (x *VerificationSummary) GetPolicyLevel() string {
	if x != nil {
		return x.PolicyLevel
	}
	return ""
}

func (x *VerificationSummary) GetDependencyLevels() map[string]uint64 {
	if x != nil {
		return x.DependencyLevels
	}
	return nil
}

type VerificationSummary_Verifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VerificationSummary_Verifier) Reset() {
	*x = VerificationSummary_Verifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationSummary_Verifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationSummary_Verifier) ProtoMessage() {}

func (x *VerificationSummary_Verifier) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationSummary_Verifier.ProtoReflect.Descriptor instead.
func (*VerificationSummary_Verifier) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDescGZIP(), []int{0, 0}
}

func (x *VerificationSummary_Verifier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VerificationSummary_Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri    string            `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Digest map[string]string `protobuf:"bytes,2,rep,name=digest,proto3" json:"digest,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VerificationSummary_Policy) Reset() {
	*x = VerificationSummary_Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationSummary_Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationSummary_Policy) ProtoMessage() {}

func (x *VerificationSummary_Policy) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationSummary_Policy.ProtoReflect.Descriptor instead.
func (*VerificationSummary_Policy) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDescGZIP(), []int{0, 1}
}

func (x *VerificationSummary_Policy) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *VerificationSummary_Policy) GetDigest() map[string]string {
	if x != nil {
		return x.Digest
	}
	return nil
}

type VerificationSummary_InputAttestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri    string            `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Digest map[string]string `protobuf:"bytes,2,rep,name=digest,proto3" json:"digest,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VerificationSummary_InputAttestation) Reset() {
	*x = VerificationSummary_InputAttestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationSummary_InputAttestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationSummary_InputAttestation) ProtoMessage() {}

func (x *VerificationSummary_InputAttestation) ProtoReflect() protoreflect.Message {
	mi := &file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationSummary_InputAttestation.ProtoReflect.Descriptor instead.
func (*VerificationSummary_InputAttestation) Descriptor() ([]byte, []int) {
	return file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDescGZIP(), []int{0, 2}
}

func (x *VerificationSummary_InputAttestation) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *VerificationSummary_InputAttestation) GetDigest() map[string]string {
	if x != nil {
		return x.Digest
	}
	return nil
}

var File_in_toto_attestation_predicates_vsa_v0_vsa_proto protoreflect.FileDescriptor

var file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x2f, 0x76, 0x73, 0x61, 0x2f, 0x76, 0x30, 0x2f, 0x76, 0x73, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x25, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x73, 0x61, 0x2e, 0x76, 0x30, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x08, 0x0a, 0x13, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x5f, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x73, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x40, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x12, 0x59, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x69, 0x6e, 0x5f, 0x74, 0x6f,
	0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x73, 0x61, 0x2e, 0x76, 0x30,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x7b, 0x0a, 0x12, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x61, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x4b, 0x2e, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x73, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x30, 0x0a, 0x13, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x7e, 0x0a, 0x11, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x50, 0x2e, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x73, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x44, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x11, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x1a, 0x1a, 0x0a, 0x08, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x1a, 0xbc, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12,
	0x65, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x4d, 0x2e, 0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x73, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0xd0, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x6f, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x57, 0x2e, 0x69, 0x6e, 0x5f, 0x74, 0x6f,
	0x74, 0x6f, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x73, 0x61, 0x2e, 0x76, 0x30,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x44, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x43, 0x0a, 0x15, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x65, 0x0a, 0x2e, 0x69, 0x6f, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6e, 0x74, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x73, 0x61, 0x2e, 0x76, 0x30, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x2d, 0x74, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x73, 0x61, 0x2f, 0x76, 0x30,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDescOnce sync.Once
	file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDescData = file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDesc
)

func file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDescGZIP() []byte {
	file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDescOnce.Do(func() {
		file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDescData = protoimpl.X.CompressGZIP(file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDescData)
	})
	return file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDescData
}

var file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_in_toto_attestation_predicates_vsa_v0_vsa_proto_goTypes = []interface{}{
	(*VerificationSummary)(nil),                  // 0: in_toto_attestation.predicates.vsa.v0.VerificationSummary
	(*VerificationSummary_Verifier)(nil),         // 1: in_toto_attestation.predicates.vsa.v0.VerificationSummary.Verifier
	(*VerificationSummary_Policy)(nil),           // 2: in_toto_attestation.predicates.vsa.v0.VerificationSummary.Policy
	(*VerificationSummary_InputAttestation)(nil), // 3: in_toto_attestation.predicates.vsa.v0.VerificationSummary.InputAttestation
	nil,                           // 4: in_toto_attestation.predicates.vsa.v0.VerificationSummary.DependencyLevelsEntry
	nil,                           // 5: in_toto_attestation.predicates.vsa.v0.VerificationSummary.Policy.DigestEntry
	nil,                           // 6: in_toto_attestation.predicates.vsa.v0.VerificationSummary.InputAttestation.DigestEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_in_toto_attestation_predicates_vsa_v0_vsa_proto_depIdxs = []int32{
	1, // 0: in_toto_attestation.predicates.vsa.v0.VerificationSummary.verifier:type_name -> in_toto_attestation.predicates.vsa.v0.VerificationSummary.Verifier
	7, // 1: in_toto_attestation.predicates.vsa.v0.VerificationSummary.time_verified:type_name -> google.protobuf.Timestamp
	2, // 2: in_toto_attestation.predicates.vsa.v0.VerificationSummary.policy:type_name -> in_toto_attestation.predicates.vsa.v0.VerificationSummary.Policy
	3, // 3: in_toto_attestation.predicates.vsa.v0.VerificationSummary.input_attestations:type_name -> in_toto_attestation.predicates.vsa.v0.VerificationSummary.InputAttestation
	4, // 4: in_toto_attestation.predicates.vsa.v0.VerificationSummary.dependency_levels:type_name -> in_toto_attestation.predicates.vsa.v0.VerificationSummary.DependencyLevelsEntry
	5, // 5: in_toto_attestation.predicates.vsa.v0.VerificationSummary.Policy.digest:type_name -> in_toto_attestation.predicates.vsa.v0.VerificationSummary.Policy.DigestEntry
	6, // 6: in_toto_attestation.predicates.vsa.v0.VerificationSummary.InputAttestation.digest:type_name -> in_toto_attestation.predicates.vsa.v0.VerificationSummary.InputAttestation.DigestEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_in_toto_attestation_predicates_vsa_v0_vsa_proto_init() }
func file_in_toto_attestation_predicates_vsa_v0_vsa_proto_init() {
	if File_in_toto_attestation_predicates_vsa_v0_vsa_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationSummary); i {
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
		file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationSummary_Verifier); i {
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
		file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationSummary_Policy); i {
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
		file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationSummary_InputAttestation); i {
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
			RawDescriptor: file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_in_toto_attestation_predicates_vsa_v0_vsa_proto_goTypes,
		DependencyIndexes: file_in_toto_attestation_predicates_vsa_v0_vsa_proto_depIdxs,
		MessageInfos:      file_in_toto_attestation_predicates_vsa_v0_vsa_proto_msgTypes,
	}.Build()
	File_in_toto_attestation_predicates_vsa_v0_vsa_proto = out.File
	file_in_toto_attestation_predicates_vsa_v0_vsa_proto_rawDesc = nil
	file_in_toto_attestation_predicates_vsa_v0_vsa_proto_goTypes = nil
	file_in_toto_attestation_predicates_vsa_v0_vsa_proto_depIdxs = nil
}
