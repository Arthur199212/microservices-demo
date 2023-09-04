// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: services/checkout/v1/checkout_service.proto

package checkoutv1

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

const (
	CheckoutService_PlaceOrder_FullMethodName = "/services.checkout.v1.CheckoutService/PlaceOrder"
)

// CheckoutServiceClient is the client API for CheckoutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckoutServiceClient interface {
	PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error)
}

type checkoutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckoutServiceClient(cc grpc.ClientConnInterface) CheckoutServiceClient {
	return &checkoutServiceClient{cc}
}

func (c *checkoutServiceClient) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error) {
	out := new(PlaceOrderResponse)
	err := c.cc.Invoke(ctx, CheckoutService_PlaceOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckoutServiceServer is the server API for CheckoutService service.
// All implementations must embed UnimplementedCheckoutServiceServer
// for forward compatibility
type CheckoutServiceServer interface {
	PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error)
	mustEmbedUnimplementedCheckoutServiceServer()
}

// UnimplementedCheckoutServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCheckoutServiceServer struct {
}

func (UnimplementedCheckoutServiceServer) PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedCheckoutServiceServer) mustEmbedUnimplementedCheckoutServiceServer() {}

// UnsafeCheckoutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckoutServiceServer will
// result in compilation errors.
type UnsafeCheckoutServiceServer interface {
	mustEmbedUnimplementedCheckoutServiceServer()
}

func RegisterCheckoutServiceServer(s grpc.ServiceRegistrar, srv CheckoutServiceServer) {
	s.RegisterService(&CheckoutService_ServiceDesc, srv)
}

func _CheckoutService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckoutService_PlaceOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).PlaceOrder(ctx, req.(*PlaceOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckoutService_ServiceDesc is the grpc.ServiceDesc for CheckoutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckoutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.checkout.v1.CheckoutService",
	HandlerType: (*CheckoutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _CheckoutService_PlaceOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/checkout/v1/checkout_service.proto",
}
