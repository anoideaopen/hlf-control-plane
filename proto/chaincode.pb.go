// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: chaincode.proto

package proto

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// Request and response for chaincode installation method
type ChaincodeInstallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source      string            `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	AuthHeaders map[string]string `protobuf:"bytes,2,rep,name=auth_headers,json=authHeaders,proto3" json:"auth_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChaincodeInstallRequest) Reset() {
	*x = ChaincodeInstallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInstallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInstallRequest) ProtoMessage() {}

func (x *ChaincodeInstallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInstallRequest.ProtoReflect.Descriptor instead.
func (*ChaincodeInstallRequest) Descriptor() ([]byte, []int) {
	return file_chaincode_proto_rawDescGZIP(), []int{0}
}

func (x *ChaincodeInstallRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ChaincodeInstallRequest) GetAuthHeaders() map[string]string {
	if x != nil {
		return x.AuthHeaders
	}
	return nil
}

type ChaincodeInstallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    []*ChaincodeInstallResponse_Result `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	PackageId string                             `protobuf:"bytes,2,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
}

func (x *ChaincodeInstallResponse) Reset() {
	*x = ChaincodeInstallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInstallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInstallResponse) ProtoMessage() {}

func (x *ChaincodeInstallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInstallResponse.ProtoReflect.Descriptor instead.
func (*ChaincodeInstallResponse) Descriptor() ([]byte, []int) {
	return file_chaincode_proto_rawDescGZIP(), []int{1}
}

func (x *ChaincodeInstallResponse) GetResult() []*ChaincodeInstallResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *ChaincodeInstallResponse) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

type ChaincodeInstallExternalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type            string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Label           string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Address         string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Timeout         string `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	TlsRequired     bool   `protobuf:"varint,5,opt,name=tls_required,json=tlsRequired,proto3" json:"tls_required,omitempty"`
	TlsClientAuth   bool   `protobuf:"varint,6,opt,name=tls_client_auth,json=tlsClientAuth,proto3" json:"tls_client_auth,omitempty"`
	TlsKey          []byte `protobuf:"bytes,7,opt,name=tls_key,json=tlsKey,proto3" json:"tls_key,omitempty"`
	TlsCert         []byte `protobuf:"bytes,8,opt,name=tls_cert,json=tlsCert,proto3" json:"tls_cert,omitempty"`
	TlsRootCert     []byte `protobuf:"bytes,9,opt,name=tls_root_cert,json=tlsRootCert,proto3" json:"tls_root_cert,omitempty"`
	EnableConnCheck bool   `protobuf:"varint,10,opt,name=enable_conn_check,json=enableConnCheck,proto3" json:"enable_conn_check,omitempty"`
}

func (x *ChaincodeInstallExternalRequest) Reset() {
	*x = ChaincodeInstallExternalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInstallExternalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInstallExternalRequest) ProtoMessage() {}

func (x *ChaincodeInstallExternalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInstallExternalRequest.ProtoReflect.Descriptor instead.
func (*ChaincodeInstallExternalRequest) Descriptor() ([]byte, []int) {
	return file_chaincode_proto_rawDescGZIP(), []int{2}
}

func (x *ChaincodeInstallExternalRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ChaincodeInstallExternalRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ChaincodeInstallExternalRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ChaincodeInstallExternalRequest) GetTimeout() string {
	if x != nil {
		return x.Timeout
	}
	return ""
}

func (x *ChaincodeInstallExternalRequest) GetTlsRequired() bool {
	if x != nil {
		return x.TlsRequired
	}
	return false
}

func (x *ChaincodeInstallExternalRequest) GetTlsClientAuth() bool {
	if x != nil {
		return x.TlsClientAuth
	}
	return false
}

func (x *ChaincodeInstallExternalRequest) GetTlsKey() []byte {
	if x != nil {
		return x.TlsKey
	}
	return nil
}

func (x *ChaincodeInstallExternalRequest) GetTlsCert() []byte {
	if x != nil {
		return x.TlsCert
	}
	return nil
}

func (x *ChaincodeInstallExternalRequest) GetTlsRootCert() []byte {
	if x != nil {
		return x.TlsRootCert
	}
	return nil
}

func (x *ChaincodeInstallExternalRequest) GetEnableConnCheck() bool {
	if x != nil {
		return x.EnableConnCheck
	}
	return false
}

type ChaincodeInstalledResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*ChaincodeInstalledResponse_Result `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ChaincodeInstalledResponse) Reset() {
	*x = ChaincodeInstalledResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInstalledResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInstalledResponse) ProtoMessage() {}

func (x *ChaincodeInstalledResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInstalledResponse.ProtoReflect.Descriptor instead.
func (*ChaincodeInstalledResponse) Descriptor() ([]byte, []int) {
	return file_chaincode_proto_rawDescGZIP(), []int{3}
}

func (x *ChaincodeInstalledResponse) GetResult() []*ChaincodeInstalledResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type ChaincodeInstalledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChaincodeInstalledRequest) Reset() {
	*x = ChaincodeInstalledRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInstalledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInstalledRequest) ProtoMessage() {}

func (x *ChaincodeInstalledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInstalledRequest.ProtoReflect.Descriptor instead.
func (*ChaincodeInstalledRequest) Descriptor() ([]byte, []int) {
	return file_chaincode_proto_rawDescGZIP(), []int{4}
}

type ChaincodeInstallResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer string `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	// Types that are assignable to Result:
	//
	//	*ChaincodeInstallResponse_Result_Existed
	//	*ChaincodeInstallResponse_Result_Err
	Result isChaincodeInstallResponse_Result_Result `protobuf_oneof:"result"`
	Label  string                                   `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *ChaincodeInstallResponse_Result) Reset() {
	*x = ChaincodeInstallResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInstallResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInstallResponse_Result) ProtoMessage() {}

func (x *ChaincodeInstallResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInstallResponse_Result.ProtoReflect.Descriptor instead.
func (*ChaincodeInstallResponse_Result) Descriptor() ([]byte, []int) {
	return file_chaincode_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ChaincodeInstallResponse_Result) GetPeer() string {
	if x != nil {
		return x.Peer
	}
	return ""
}

func (m *ChaincodeInstallResponse_Result) GetResult() isChaincodeInstallResponse_Result_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ChaincodeInstallResponse_Result) GetExisted() bool {
	if x, ok := x.GetResult().(*ChaincodeInstallResponse_Result_Existed); ok {
		return x.Existed
	}
	return false
}

func (x *ChaincodeInstallResponse_Result) GetErr() string {
	if x, ok := x.GetResult().(*ChaincodeInstallResponse_Result_Err); ok {
		return x.Err
	}
	return ""
}

func (x *ChaincodeInstallResponse_Result) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type isChaincodeInstallResponse_Result_Result interface {
	isChaincodeInstallResponse_Result_Result()
}

type ChaincodeInstallResponse_Result_Existed struct {
	Existed bool `protobuf:"varint,2,opt,name=existed,proto3,oneof"`
}

type ChaincodeInstallResponse_Result_Err struct {
	Err string `protobuf:"bytes,3,opt,name=err,proto3,oneof"`
}

func (*ChaincodeInstallResponse_Result_Existed) isChaincodeInstallResponse_Result_Result() {}

func (*ChaincodeInstallResponse_Result_Err) isChaincodeInstallResponse_Result_Result() {}

type ChaincodeInstalledResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageId string   `protobuf:"bytes,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	Label     string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Peers     []string `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *ChaincodeInstalledResponse_Result) Reset() {
	*x = ChaincodeInstalledResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInstalledResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInstalledResponse_Result) ProtoMessage() {}

func (x *ChaincodeInstalledResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInstalledResponse_Result.ProtoReflect.Descriptor instead.
func (*ChaincodeInstalledResponse_Result) Descriptor() ([]byte, []int) {
	return file_chaincode_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ChaincodeInstalledResponse_Result) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *ChaincodeInstalledResponse_Result) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ChaincodeInstalledResponse_Result) GetPeers() []string {
	if x != nil {
		return x.Peers
	}
	return nil
}

var File_chaincode_proto protoreflect.FileDescriptor

var file_chaincode_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x02, 0x0a, 0x17, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0x92, 0x41, 0x24, 0x32, 0x22, 0x55, 0x52, 0x4c, 0x20, 0x6c,
	0x69, 0x6e, 0x6b, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x20, 0x77,
	0x69, 0x74, 0x68, 0x20, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x27, 0x92, 0x41,
	0x24, 0x32, 0x22, 0x55, 0x52, 0x4c, 0x20, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x74, 0x6f, 0x20, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x3a, 0x5f, 0x92, 0x41, 0x5c, 0x0a, 0x5a, 0x2a, 0x17, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x32, 0x36, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x6e, 0x64, 0x20,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0xd2, 0x01, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0xd4, 0x03, 0x0a, 0x18, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x69, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x29, 0x92, 0x41, 0x26, 0x32, 0x24, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x65, 0x61, 0x63, 0x68, 0x20, 0x70,
	0x65, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x1a, 0xad, 0x02, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1d, 0x92, 0x41, 0x1a, 0x32, 0x18, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x70, 0x65,
	0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x5f, 0x0a, 0x07, 0x65, 0x78, 0x69, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x43, 0x92, 0x41, 0x40, 0x32, 0x3e,
	0x46, 0x6c, 0x61, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x20, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x6f, 0x6e, 0x20, 0x70,
	0x65, 0x65, 0x72, 0x2c, 0x20, 0x74, 0x72, 0x75, 0x65, 0x20, 0x69, 0x66, 0x20, 0x61, 0x6c, 0x72,
	0x65, 0x61, 0x64, 0x79, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x07, 0x65, 0x78, 0x69, 0x73, 0x74, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0x92, 0x41, 0x2d, 0x32, 0x2b, 0x4f, 0x63, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x64, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x77, 0x69, 0x74, 0x68,
	0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x66, 0x20,
	0x68, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x48, 0x00, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12,
	0x3f, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29,
	0x92, 0x41, 0x26, 0x32, 0x24, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x20, 0x6f, 0x66, 0x20, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x20, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x20, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x80, 0x04, 0x0a, 0x1f, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0x92, 0x41,
	0x19, 0x32, 0x17, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x2e, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x18, 0x92, 0x41, 0x15, 0x32, 0x13, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x2e, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x14, 0x92, 0x41, 0x11, 0x32, 0x0f, 0x42, 0x61, 0x73, 0x65, 0x20, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x3a, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x20, 0x92, 0x41, 0x1d, 0x32, 0x1b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x20, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c,
	0x74, 0x6c, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x19, 0x92, 0x41, 0x16, 0x32, 0x14, 0x46, 0x6c, 0x61, 0x67, 0x20, 0x6f, 0x66,
	0x20, 0x74, 0x6c, 0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x0b, 0x74,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x0f, 0x74, 0x6c,
	0x73, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x25, 0x92, 0x41, 0x22, 0x32, 0x20, 0x46, 0x6c, 0x61, 0x67, 0x20, 0x6f,
	0x66, 0x20, 0x74, 0x6c, 0x73, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x61, 0x75, 0x74,
	0x68, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6c, 0x73,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x6c, 0x73, 0x4b,
	0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x6c, 0x73, 0x43, 0x65, 0x72, 0x74, 0x12, 0x22, 0x0a,
	0x0d, 0x74, 0x6c, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74, 0x6c, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x22, 0xb3, 0x01,
	0x0a, 0x1a, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x53,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65,
	0x65, 0x72, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x2d, 0x74, 0x2e, 0x69,
	0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x68,
	0x6c, 0x66, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x68, 0x6c, 0x66, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chaincode_proto_rawDescOnce sync.Once
	file_chaincode_proto_rawDescData = file_chaincode_proto_rawDesc
)

func file_chaincode_proto_rawDescGZIP() []byte {
	file_chaincode_proto_rawDescOnce.Do(func() {
		file_chaincode_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaincode_proto_rawDescData)
	})
	return file_chaincode_proto_rawDescData
}

var file_chaincode_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chaincode_proto_goTypes = []interface{}{
	(*ChaincodeInstallRequest)(nil),           // 0: proto.ChaincodeInstallRequest
	(*ChaincodeInstallResponse)(nil),          // 1: proto.ChaincodeInstallResponse
	(*ChaincodeInstallExternalRequest)(nil),   // 2: proto.ChaincodeInstallExternalRequest
	(*ChaincodeInstalledResponse)(nil),        // 3: proto.ChaincodeInstalledResponse
	(*ChaincodeInstalledRequest)(nil),         // 4: proto.ChaincodeInstalledRequest
	nil,                                       // 5: proto.ChaincodeInstallRequest.AuthHeadersEntry
	(*ChaincodeInstallResponse_Result)(nil),   // 6: proto.ChaincodeInstallResponse.Result
	(*ChaincodeInstalledResponse_Result)(nil), // 7: proto.ChaincodeInstalledResponse.Result
}
var file_chaincode_proto_depIdxs = []int32{
	5, // 0: proto.ChaincodeInstallRequest.auth_headers:type_name -> proto.ChaincodeInstallRequest.AuthHeadersEntry
	6, // 1: proto.ChaincodeInstallResponse.result:type_name -> proto.ChaincodeInstallResponse.Result
	7, // 2: proto.ChaincodeInstalledResponse.result:type_name -> proto.ChaincodeInstalledResponse.Result
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chaincode_proto_init() }
func file_chaincode_proto_init() {
	if File_chaincode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chaincode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInstallRequest); i {
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
		file_chaincode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInstallResponse); i {
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
		file_chaincode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInstallExternalRequest); i {
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
		file_chaincode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInstalledResponse); i {
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
		file_chaincode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInstalledRequest); i {
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
		file_chaincode_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInstallResponse_Result); i {
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
		file_chaincode_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInstalledResponse_Result); i {
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
	file_chaincode_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ChaincodeInstallResponse_Result_Existed)(nil),
		(*ChaincodeInstallResponse_Result_Err)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chaincode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chaincode_proto_goTypes,
		DependencyIndexes: file_chaincode_proto_depIdxs,
		MessageInfos:      file_chaincode_proto_msgTypes,
	}.Build()
	File_chaincode_proto = out.File
	file_chaincode_proto_rawDesc = nil
	file_chaincode_proto_goTypes = nil
	file_chaincode_proto_depIdxs = nil
}
