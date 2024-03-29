// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: services/cart/v1/rpc_clear_cart.proto

package cartv1

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

type ClearCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *ClearCartRequest) Reset() {
	*x = ClearCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cart_v1_rpc_clear_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCartRequest) ProtoMessage() {}

func (x *ClearCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cart_v1_rpc_clear_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCartRequest.ProtoReflect.Descriptor instead.
func (*ClearCartRequest) Descriptor() ([]byte, []int) {
	return file_services_cart_v1_rpc_clear_cart_proto_rawDescGZIP(), []int{0}
}

func (x *ClearCartRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type ClearCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClearCartResponse) Reset() {
	*x = ClearCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_cart_v1_rpc_clear_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCartResponse) ProtoMessage() {}

func (x *ClearCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_cart_v1_rpc_clear_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCartResponse.ProtoReflect.Descriptor instead.
func (*ClearCartResponse) Descriptor() ([]byte, []int) {
	return file_services_cart_v1_rpc_clear_cart_proto_rawDescGZIP(), []int{1}
}

var File_services_cart_v1_rpc_clear_cart_proto protoreflect.FileDescriptor

var file_services_cart_v1_rpc_clear_cart_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x31, 0x0a, 0x10, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xd3, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x52, 0x70, 0x63, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x74, 0x68,
	0x75, 0x72, 0x31, 0x39, 0x39, 0x32, 0x31, 0x32, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x61, 0x72, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x43, 0x58, 0xaa, 0x02, 0x10,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x43, 0x61, 0x72, 0x74,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x43,
	0x61, 0x72, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x43,
	0x61, 0x72, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_cart_v1_rpc_clear_cart_proto_rawDescOnce sync.Once
	file_services_cart_v1_rpc_clear_cart_proto_rawDescData = file_services_cart_v1_rpc_clear_cart_proto_rawDesc
)

func file_services_cart_v1_rpc_clear_cart_proto_rawDescGZIP() []byte {
	file_services_cart_v1_rpc_clear_cart_proto_rawDescOnce.Do(func() {
		file_services_cart_v1_rpc_clear_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_cart_v1_rpc_clear_cart_proto_rawDescData)
	})
	return file_services_cart_v1_rpc_clear_cart_proto_rawDescData
}

var file_services_cart_v1_rpc_clear_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_cart_v1_rpc_clear_cart_proto_goTypes = []interface{}{
	(*ClearCartRequest)(nil),  // 0: services.cart.v1.ClearCartRequest
	(*ClearCartResponse)(nil), // 1: services.cart.v1.ClearCartResponse
}
var file_services_cart_v1_rpc_clear_cart_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_cart_v1_rpc_clear_cart_proto_init() }
func file_services_cart_v1_rpc_clear_cart_proto_init() {
	if File_services_cart_v1_rpc_clear_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_cart_v1_rpc_clear_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearCartRequest); i {
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
		file_services_cart_v1_rpc_clear_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearCartResponse); i {
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
			RawDescriptor: file_services_cart_v1_rpc_clear_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_cart_v1_rpc_clear_cart_proto_goTypes,
		DependencyIndexes: file_services_cart_v1_rpc_clear_cart_proto_depIdxs,
		MessageInfos:      file_services_cart_v1_rpc_clear_cart_proto_msgTypes,
	}.Build()
	File_services_cart_v1_rpc_clear_cart_proto = out.File
	file_services_cart_v1_rpc_clear_cart_proto_rawDesc = nil
	file_services_cart_v1_rpc_clear_cart_proto_goTypes = nil
	file_services_cart_v1_rpc_clear_cart_proto_depIdxs = nil
}
