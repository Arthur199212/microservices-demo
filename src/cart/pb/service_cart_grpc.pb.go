// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: service_cart.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CartClient is the client API for Cart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartClient interface {
	AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error)
}

type cartClient struct {
	cc grpc.ClientConnInterface
}

func NewCartClient(cc grpc.ClientConnInterface) CartClient {
	return &cartClient{cc}
}

func (c *cartClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Cart/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Cart/ClearCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error) {
	out := new(GetCartResponse)
	err := c.cc.Invoke(ctx, "/pb.Cart/GetCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServer is the server API for Cart service.
// All implementations must embed UnimplementedCartServer
// for forward compatibility
type CartServer interface {
	AddItem(context.Context, *AddItemRequest) (*emptypb.Empty, error)
	ClearCart(context.Context, *ClearCartRequest) (*emptypb.Empty, error)
	GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error)
	mustEmbedUnimplementedCartServer()
}

// UnimplementedCartServer must be embedded to have forward compatible implementations.
type UnimplementedCartServer struct {
}

func (UnimplementedCartServer) AddItem(context.Context, *AddItemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (UnimplementedCartServer) ClearCart(context.Context, *ClearCartRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCart not implemented")
}
func (UnimplementedCartServer) GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServer) mustEmbedUnimplementedCartServer() {}

// UnsafeCartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServer will
// result in compilation errors.
type UnsafeCartServer interface {
	mustEmbedUnimplementedCartServer()
}

func RegisterCartServer(s grpc.ServiceRegistrar, srv CartServer) {
	s.RegisterService(&Cart_ServiceDesc, srv)
}

func _Cart_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Cart/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).AddItem(ctx, req.(*AddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_ClearCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).ClearCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Cart/ClearCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).ClearCart(ctx, req.(*ClearCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Cart/GetCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).GetCart(ctx, req.(*GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cart_ServiceDesc is the grpc.ServiceDesc for Cart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Cart",
	HandlerType: (*CartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddItem",
			Handler:    _Cart_AddItem_Handler,
		},
		{
			MethodName: "ClearCart",
			Handler:    _Cart_ClearCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _Cart_GetCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_cart.proto",
}
