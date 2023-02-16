// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: services/cart/v1/cart_service.proto

package cartv1

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

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error)
	ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error)
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error) {
	out := new(AddItemResponse)
	err := c.cc.Invoke(ctx, "/services.cart.v1.CartService/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error) {
	out := new(ClearCartResponse)
	err := c.cc.Invoke(ctx, "/services.cart.v1.CartService/ClearCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error) {
	out := new(GetCartResponse)
	err := c.cc.Invoke(ctx, "/services.cart.v1.CartService/GetCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error)
	ClearCart(context.Context, *ClearCartRequest) (*ClearCartResponse, error)
	GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedCartServiceServer) ClearCart(context.Context, *ClearCartRequest) (*ClearCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCart not implemented")
}
func (UnimplementedCartServiceServer) GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.cart.v1.CartService/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddItem(ctx, req.(*AddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_ClearCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).ClearCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.cart.v1.CartService/ClearCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).ClearCart(ctx, req.(*ClearCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.cart.v1.CartService/GetCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCart(ctx, req.(*GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.cart.v1.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _CartService_AddItem_Handler,
		},
		{
			MethodName: "ClearCart",
			Handler:    _CartService_ClearCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _CartService_GetCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/cart/v1/cart_service.proto",
}
