// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: rpc_get_quote.proto

package pb

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

type GetQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  *Address   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Products []*Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GetQuoteRequest) Reset() {
	*x = GetQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_quote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteRequest) ProtoMessage() {}

func (x *GetQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_quote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteRequest.ProtoReflect.Descriptor instead.
func (*GetQuoteRequest) Descriptor() ([]byte, []int) {
	return file_rpc_get_quote_proto_rawDescGZIP(), []int{0}
}

func (x *GetQuoteRequest) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *GetQuoteRequest) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type GetQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote        float32 `protobuf:"fixed32,1,opt,name=quote,proto3" json:"quote,omitempty"`
	CurrencyCode string  `protobuf:"bytes,2,opt,name=currencyCode,proto3" json:"currencyCode,omitempty"`
}

func (x *GetQuoteResponse) Reset() {
	*x = GetQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_quote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteResponse) ProtoMessage() {}

func (x *GetQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_quote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteResponse.ProtoReflect.Descriptor instead.
func (*GetQuoteResponse) Descriptor() ([]byte, []int) {
	return file_rpc_get_quote_proto_rawDescGZIP(), []int{1}
}

func (x *GetQuoteResponse) GetQuote() float32 {
	if x != nil {
		return x.Quote
	}
	return 0
}

func (x *GetQuoteResponse) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

var File_rpc_get_quote_proto protoreflect.FileDescriptor

var file_rpc_get_quote_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x27, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x4c, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x74, 0x68, 0x75, 0x72, 0x31, 0x39, 0x39,
	0x32, 0x31, 0x32, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_get_quote_proto_rawDescOnce sync.Once
	file_rpc_get_quote_proto_rawDescData = file_rpc_get_quote_proto_rawDesc
)

func file_rpc_get_quote_proto_rawDescGZIP() []byte {
	file_rpc_get_quote_proto_rawDescOnce.Do(func() {
		file_rpc_get_quote_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_get_quote_proto_rawDescData)
	})
	return file_rpc_get_quote_proto_rawDescData
}

var file_rpc_get_quote_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_get_quote_proto_goTypes = []interface{}{
	(*GetQuoteRequest)(nil),  // 0: pb.GetQuoteRequest
	(*GetQuoteResponse)(nil), // 1: pb.GetQuoteResponse
	(*Address)(nil),          // 2: pb.Address
	(*Product)(nil),          // 3: pb.Product
}
var file_rpc_get_quote_proto_depIdxs = []int32{
	2, // 0: pb.GetQuoteRequest.address:type_name -> pb.Address
	3, // 1: pb.GetQuoteRequest.products:type_name -> pb.Product
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_get_quote_proto_init() }
func file_rpc_get_quote_proto_init() {
	if File_rpc_get_quote_proto != nil {
		return
	}
	file_address_proto_init()
	file_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_get_quote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteRequest); i {
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
		file_rpc_get_quote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteResponse); i {
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
			RawDescriptor: file_rpc_get_quote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_get_quote_proto_goTypes,
		DependencyIndexes: file_rpc_get_quote_proto_depIdxs,
		MessageInfos:      file_rpc_get_quote_proto_msgTypes,
	}.Build()
	File_rpc_get_quote_proto = out.File
	file_rpc_get_quote_proto_rawDesc = nil
	file_rpc_get_quote_proto_goTypes = nil
	file_rpc_get_quote_proto_depIdxs = nil
}
