// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service_shipping.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShippingClient is the client API for Shipping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShippingClient interface {
	GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*GetQuoteResponse, error)
	ShipOrder(ctx context.Context, in *ShipOrderRequest, opts ...grpc.CallOption) (*ShipOrderResponse, error)
}

type shippingClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingClient(cc grpc.ClientConnInterface) ShippingClient {
	return &shippingClient{cc}
}

func (c *shippingClient) GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*GetQuoteResponse, error) {
	out := new(GetQuoteResponse)
	err := c.cc.Invoke(ctx, "/pb.Shipping/GetQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingClient) ShipOrder(ctx context.Context, in *ShipOrderRequest, opts ...grpc.CallOption) (*ShipOrderResponse, error) {
	out := new(ShipOrderResponse)
	err := c.cc.Invoke(ctx, "/pb.Shipping/ShipOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServer is the server API for Shipping service.
// All implementations must embed UnimplementedShippingServer
// for forward compatibility
type ShippingServer interface {
	GetQuote(context.Context, *GetQuoteRequest) (*GetQuoteResponse, error)
	ShipOrder(context.Context, *ShipOrderRequest) (*ShipOrderResponse, error)
	mustEmbedUnimplementedShippingServer()
}

// UnimplementedShippingServer must be embedded to have forward compatible implementations.
type UnimplementedShippingServer struct {
}

func (UnimplementedShippingServer) GetQuote(context.Context, *GetQuoteRequest) (*GetQuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuote not implemented")
}
func (UnimplementedShippingServer) ShipOrder(context.Context, *ShipOrderRequest) (*ShipOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShipOrder not implemented")
}
func (UnimplementedShippingServer) mustEmbedUnimplementedShippingServer() {}

// UnsafeShippingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShippingServer will
// result in compilation errors.
type UnsafeShippingServer interface {
	mustEmbedUnimplementedShippingServer()
}

func RegisterShippingServer(s grpc.ServiceRegistrar, srv ShippingServer) {
	s.RegisterService(&Shipping_ServiceDesc, srv)
}

func _Shipping_GetQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServer).GetQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Shipping/GetQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServer).GetQuote(ctx, req.(*GetQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shipping_ShipOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShipOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServer).ShipOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Shipping/ShipOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServer).ShipOrder(ctx, req.(*ShipOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shipping_ServiceDesc is the grpc.ServiceDesc for Shipping service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shipping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Shipping",
	HandlerType: (*ShippingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuote",
			Handler:    _Shipping_GetQuote_Handler,
		},
		{
			MethodName: "ShipOrder",
			Handler:    _Shipping_ShipOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_shipping.proto",
}