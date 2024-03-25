// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: ordering.proto

package proto

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

type OrderingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host          string                    `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port          uint32                    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelName   string                    `protobuf:"bytes,3,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
	JoinedOrderer []*OrderingRequest_Joined `protobuf:"bytes,4,rep,name=joined_orderer,json=joinedOrderer,proto3" json:"joined_orderer,omitempty"`
}

func (x *OrderingRequest) Reset() {
	*x = OrderingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderingRequest) ProtoMessage() {}

func (x *OrderingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderingRequest.ProtoReflect.Descriptor instead.
func (*OrderingRequest) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{0}
}

func (x *OrderingRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *OrderingRequest) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *OrderingRequest) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *OrderingRequest) GetJoinedOrderer() []*OrderingRequest_Joined {
	if x != nil {
		return x.JoinedOrderer
	}
	return nil
}

type OrderingRemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderingRemoveResponse) Reset() {
	*x = OrderingRemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderingRemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderingRemoveResponse) ProtoMessage() {}

func (x *OrderingRemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderingRemoveResponse.ProtoReflect.Descriptor instead.
func (*OrderingRemoveResponse) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{1}
}

type OrderingJoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *OrderingJoinResponse) Reset() {
	*x = OrderingJoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderingJoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderingJoinResponse) ProtoMessage() {}

func (x *OrderingJoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderingJoinResponse.ProtoReflect.Descriptor instead.
func (*OrderingJoinResponse) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{2}
}

func (x *OrderingJoinResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type OrderingListInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*OrderingListInfoResponse_ChannelInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *OrderingListInfoResponse) Reset() {
	*x = OrderingListInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderingListInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderingListInfoResponse) ProtoMessage() {}

func (x *OrderingListInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderingListInfoResponse.ProtoReflect.Descriptor instead.
func (*OrderingListInfoResponse) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{3}
}

func (x *OrderingListInfoResponse) GetInfo() []*OrderingListInfoResponse_ChannelInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type OrderingRequest_Joined struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *OrderingRequest_Joined) Reset() {
	*x = OrderingRequest_Joined{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderingRequest_Joined) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderingRequest_Joined) ProtoMessage() {}

func (x *OrderingRequest_Joined) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderingRequest_Joined.ProtoReflect.Descriptor instead.
func (*OrderingRequest_Joined) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OrderingRequest_Joined) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *OrderingRequest_Joined) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type OrderingListInfoResponse_ChannelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status            string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	ConsensusRelation string `protobuf:"bytes,3,opt,name=consensus_relation,json=consensusRelation,proto3" json:"consensus_relation,omitempty"`
	Height            uint64 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *OrderingListInfoResponse_ChannelInfo) Reset() {
	*x = OrderingListInfoResponse_ChannelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ordering_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderingListInfoResponse_ChannelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderingListInfoResponse_ChannelInfo) ProtoMessage() {}

func (x *OrderingListInfoResponse_ChannelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ordering_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderingListInfoResponse_ChannelInfo.ProtoReflect.Descriptor instead.
func (*OrderingListInfoResponse_ChannelInfo) Descriptor() ([]byte, []int) {
	return file_ordering_proto_rawDescGZIP(), []int{3, 0}
}

func (x *OrderingListInfoResponse_ChannelInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderingListInfoResponse_ChannelInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderingListInfoResponse_ChannelInfo) GetConsensusRelation() string {
	if x != nil {
		return x.ConsensusRelation
	}
	return ""
}

func (x *OrderingListInfoResponse_ChannelInfo) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_ordering_proto protoreflect.FileDescriptor

var file_ordering_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x0d, 0x6a,
	0x6f, 0x69, 0x6e, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x1a, 0x30, 0x0a, 0x06,
	0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x18,
	0x0a, 0x16, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0xde, 0x01, 0x0a, 0x18, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x80, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x6e, 0x2d, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x68, 0x6c, 0x66, 0x2d, 0x74, 0x6f, 0x6f, 0x6c,
	0x2f, 0x68, 0x6c, 0x66, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ordering_proto_rawDescOnce sync.Once
	file_ordering_proto_rawDescData = file_ordering_proto_rawDesc
)

func file_ordering_proto_rawDescGZIP() []byte {
	file_ordering_proto_rawDescOnce.Do(func() {
		file_ordering_proto_rawDescData = protoimpl.X.CompressGZIP(file_ordering_proto_rawDescData)
	})
	return file_ordering_proto_rawDescData
}

var file_ordering_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ordering_proto_goTypes = []interface{}{
	(*OrderingRequest)(nil),                      // 0: proto.OrderingRequest
	(*OrderingRemoveResponse)(nil),               // 1: proto.OrderingRemoveResponse
	(*OrderingJoinResponse)(nil),                 // 2: proto.OrderingJoinResponse
	(*OrderingListInfoResponse)(nil),             // 3: proto.OrderingListInfoResponse
	(*OrderingRequest_Joined)(nil),               // 4: proto.OrderingRequest.Joined
	(*OrderingListInfoResponse_ChannelInfo)(nil), // 5: proto.OrderingListInfoResponse.ChannelInfo
}
var file_ordering_proto_depIdxs = []int32{
	4, // 0: proto.OrderingRequest.joined_orderer:type_name -> proto.OrderingRequest.Joined
	5, // 1: proto.OrderingListInfoResponse.info:type_name -> proto.OrderingListInfoResponse.ChannelInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ordering_proto_init() }
func file_ordering_proto_init() {
	if File_ordering_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ordering_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderingRequest); i {
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
		file_ordering_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderingRemoveResponse); i {
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
		file_ordering_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderingJoinResponse); i {
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
		file_ordering_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderingListInfoResponse); i {
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
		file_ordering_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderingRequest_Joined); i {
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
		file_ordering_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderingListInfoResponse_ChannelInfo); i {
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
			RawDescriptor: file_ordering_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ordering_proto_goTypes,
		DependencyIndexes: file_ordering_proto_depIdxs,
		MessageInfos:      file_ordering_proto_msgTypes,
	}.Build()
	File_ordering_proto = out.File
	file_ordering_proto_rawDesc = nil
	file_ordering_proto_goTypes = nil
	file_ordering_proto_depIdxs = nil
}