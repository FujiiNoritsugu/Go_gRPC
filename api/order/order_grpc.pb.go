// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: order.proto

package order

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

// OrderToClient is the client API for OrderTo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderToClient interface {
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
}

type orderToClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderToClient(cc grpc.ClientConnInterface) OrderToClient {
	return &orderToClient{cc}
}

func (c *orderToClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "/order.OrderTo/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderToServer is the server API for OrderTo service.
// All implementations must embed UnimplementedOrderToServer
// for forward compatibility
type OrderToServer interface {
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	mustEmbedUnimplementedOrderToServer()
}

// UnimplementedOrderToServer must be embedded to have forward compatible implementations.
type UnimplementedOrderToServer struct {
}

func (UnimplementedOrderToServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderToServer) mustEmbedUnimplementedOrderToServer() {}

// UnsafeOrderToServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderToServer will
// result in compilation errors.
type UnsafeOrderToServer interface {
	mustEmbedUnimplementedOrderToServer()
}

func RegisterOrderToServer(s grpc.ServiceRegistrar, srv OrderToServer) {
	s.RegisterService(&OrderTo_ServiceDesc, srv)
}

func _OrderTo_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderToServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderTo/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderToServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderTo_ServiceDesc is the grpc.ServiceDesc for OrderTo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderTo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderTo",
	HandlerType: (*OrderToServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _OrderTo_GetOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}