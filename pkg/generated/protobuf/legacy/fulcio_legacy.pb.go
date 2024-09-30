//
// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: fulcio_legacy.proto

package legacy

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateSigningCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The public key to be stored in the requested certificate
	//
	// Deprecated: Marked as deprecated in fulcio_legacy.proto.
	PublicKey *PublicKey `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	// Proof that the client possesses the private key
	//
	// Deprecated: Marked as deprecated in fulcio_legacy.proto.
	SignedEmailAddress []byte `protobuf:"bytes,2,opt,name=signedEmailAddress,proto3" json:"signedEmailAddress,omitempty"`
	// Optional: PKCS#10 PEM-encoded certificate signing request
	// Contains the public key to be stored in the requested
	// certificate. All other CSR fields are ignored. Since
	// the CSR is self-signed, it also acts as a proof of
	// possession of the private key.
	//
	// In particular, the CSR's subject name is not verified, or tested for
	// compatibility with its specified X.509 name type (e.g. email address).
	//
	// Deprecated: Marked as deprecated in fulcio_legacy.proto.
	CertificateSigningRequest []byte `protobuf:"bytes,3,opt,name=certificateSigningRequest,proto3" json:"certificateSigningRequest,omitempty"`
}

func (x *CreateSigningCertificateRequest) Reset() {
	*x = CreateSigningCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcio_legacy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSigningCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSigningCertificateRequest) ProtoMessage() {}

func (x *CreateSigningCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulcio_legacy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSigningCertificateRequest.ProtoReflect.Descriptor instead.
func (*CreateSigningCertificateRequest) Descriptor() ([]byte, []int) {
	return file_fulcio_legacy_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in fulcio_legacy.proto.
func (x *CreateSigningCertificateRequest) GetPublicKey() *PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// Deprecated: Marked as deprecated in fulcio_legacy.proto.
func (x *CreateSigningCertificateRequest) GetSignedEmailAddress() []byte {
	if x != nil {
		return x.SignedEmailAddress
	}
	return nil
}

// Deprecated: Marked as deprecated in fulcio_legacy.proto.
func (x *CreateSigningCertificateRequest) GetCertificateSigningRequest() []byte {
	if x != nil {
		return x.CertificateSigningRequest
	}
	return nil
}

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The cryptographic algorithm to use with the key material
	//
	// Deprecated: Marked as deprecated in fulcio_legacy.proto.
	Algorithm string `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	// PKIX, ASN.1 DER or PEM-encoded public key. PEM is typically
	// of type PUBLIC KEY.
	//
	// Deprecated: Marked as deprecated in fulcio_legacy.proto.
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulcio_legacy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_fulcio_legacy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_fulcio_legacy_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in fulcio_legacy.proto.
func (x *PublicKey) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

// Deprecated: Marked as deprecated in fulcio_legacy.proto.
func (x *PublicKey) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_fulcio_legacy_proto protoreflect.FileDescriptor

var file_fulcio_legacy_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x75, 0x6c, 0x63, 0x69, 0x6f, 0x5f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x66, 0x75, 0x6c, 0x63, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x1f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x66,
	0x75, 0x6c, 0x63, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x06, 0xe2, 0x41, 0x01, 0x01, 0x18, 0x01, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x12, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0xe2, 0x41, 0x01, 0x01, 0x18, 0x01, 0x52, 0x12, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x44, 0x0a, 0x19, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0xe2, 0x41, 0x01, 0x01, 0x18, 0x01, 0x52, 0x19, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0xe2, 0x41, 0x01, 0x02, 0x18, 0x01, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xf8, 0x01, 0x0a, 0x02, 0x43, 0x41, 0x12,
	0x90, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e,
	0x67, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x75, 0x6c, 0x63,
	0x69, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x65, 0x72, 0x74, 0x88,
	0x02, 0x01, 0x12, 0x5f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74,
	0x88, 0x02, 0x01, 0x42, 0xa1, 0x03, 0x92, 0x41, 0xb8, 0x02, 0x12, 0xc0, 0x01, 0x0a, 0x0d, 0x46,
	0x75, 0x6c, 0x63, 0x69, 0x6f, 0x20, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x22, 0x5c, 0x0a, 0x17,
	0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x46, 0x75, 0x6c, 0x63, 0x69, 0x6f, 0x20,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x67, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x75, 0x6c, 0x63, 0x69, 0x6f, 0x1a, 0x1d, 0x73, 0x69, 0x67,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x64, 0x65, 0x76, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x4a, 0x0a, 0x12, 0x41, 0x70,
	0x61, 0x63, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x32, 0x2e, 0x30,
	0x12, 0x34, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x75,
	0x6c, 0x63, 0x69, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x4c,
	0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x1a, 0x13, 0x66,
	0x75, 0x6c, 0x63, 0x69, 0x6f, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x37, 0x0a, 0x11, 0x4d, 0x6f, 0x72,
	0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x46, 0x75, 0x6c, 0x63, 0x69, 0x6f, 0x12, 0x22,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x75, 0x6c, 0x63,
	0x69, 0x6f, 0x0a, 0x1a, 0x64, 0x65, 0x76, 0x2e, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x66, 0x75, 0x6c, 0x63, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0b,
	0x46, 0x75, 0x6c, 0x63, 0x69, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x67, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x66, 0x75, 0x6c, 0x63, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fulcio_legacy_proto_rawDescOnce sync.Once
	file_fulcio_legacy_proto_rawDescData = file_fulcio_legacy_proto_rawDesc
)

func file_fulcio_legacy_proto_rawDescGZIP() []byte {
	file_fulcio_legacy_proto_rawDescOnce.Do(func() {
		file_fulcio_legacy_proto_rawDescData = protoimpl.X.CompressGZIP(file_fulcio_legacy_proto_rawDescData)
	})
	return file_fulcio_legacy_proto_rawDescData
}

var file_fulcio_legacy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fulcio_legacy_proto_goTypes = []any{
	(*CreateSigningCertificateRequest)(nil), // 0: dev.sigstore.fulcio.v1beta.CreateSigningCertificateRequest
	(*PublicKey)(nil),                       // 1: dev.sigstore.fulcio.v1beta.PublicKey
	(*emptypb.Empty)(nil),                   // 2: google.protobuf.Empty
	(*httpbody.HttpBody)(nil),               // 3: google.api.HttpBody
}
var file_fulcio_legacy_proto_depIdxs = []int32{
	1, // 0: dev.sigstore.fulcio.v1beta.CreateSigningCertificateRequest.publicKey:type_name -> dev.sigstore.fulcio.v1beta.PublicKey
	0, // 1: dev.sigstore.fulcio.v1beta.CA.CreateSigningCertificate:input_type -> dev.sigstore.fulcio.v1beta.CreateSigningCertificateRequest
	2, // 2: dev.sigstore.fulcio.v1beta.CA.GetRootCertificate:input_type -> google.protobuf.Empty
	3, // 3: dev.sigstore.fulcio.v1beta.CA.CreateSigningCertificate:output_type -> google.api.HttpBody
	3, // 4: dev.sigstore.fulcio.v1beta.CA.GetRootCertificate:output_type -> google.api.HttpBody
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fulcio_legacy_proto_init() }
func file_fulcio_legacy_proto_init() {
	if File_fulcio_legacy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fulcio_legacy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSigningCertificateRequest); i {
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
		file_fulcio_legacy_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PublicKey); i {
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
			RawDescriptor: file_fulcio_legacy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fulcio_legacy_proto_goTypes,
		DependencyIndexes: file_fulcio_legacy_proto_depIdxs,
		MessageInfos:      file_fulcio_legacy_proto_msgTypes,
	}.Build()
	File_fulcio_legacy_proto = out.File
	file_fulcio_legacy_proto_rawDesc = nil
	file_fulcio_legacy_proto_goTypes = nil
	file_fulcio_legacy_proto_depIdxs = nil
}
