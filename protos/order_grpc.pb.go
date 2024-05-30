// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: protos/order.proto

package protos

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
	OrderService_CreateOrder_FullMethodName = "/OrderService/CreateOrder"
	OrderService_GetOrders_FullMethodName   = "/OrderService/GetOrders"
	OrderService_GetOrder_FullMethodName    = "/OrderService/GetOrder"
	OrderService_UpdateOrder_FullMethodName = "/OrderService/UpdateOrder"
	OrderService_DeleteOrder_FullMethodName = "/OrderService/DeleteOrder"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*CreateUpdateDeleteOrderResp, error)
	GetOrders(ctx context.Context, in *OrderEmpty, opts ...grpc.CallOption) (*GetOrdersResp, error)
	GetOrder(ctx context.Context, in *GetDeleteOrderReq, opts ...grpc.CallOption) (*GetOrderResp, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*CreateUpdateDeleteOrderResp, error)
	DeleteOrder(ctx context.Context, in *GetDeleteOrderReq, opts ...grpc.CallOption) (*CreateUpdateDeleteOrderResp, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*CreateUpdateDeleteOrderResp, error) {
	out := new(CreateUpdateDeleteOrderResp)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrders(ctx context.Context, in *OrderEmpty, opts ...grpc.CallOption) (*GetOrdersResp, error) {
	out := new(GetOrdersResp)
	err := c.cc.Invoke(ctx, OrderService_GetOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *GetDeleteOrderReq, opts ...grpc.CallOption) (*GetOrderResp, error) {
	out := new(GetOrderResp)
	err := c.cc.Invoke(ctx, OrderService_GetOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*CreateUpdateDeleteOrderResp, error) {
	out := new(CreateUpdateDeleteOrderResp)
	err := c.cc.Invoke(ctx, OrderService_UpdateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteOrder(ctx context.Context, in *GetDeleteOrderReq, opts ...grpc.CallOption) (*CreateUpdateDeleteOrderResp, error) {
	out := new(CreateUpdateDeleteOrderResp)
	err := c.cc.Invoke(ctx, OrderService_DeleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	CreateOrder(context.Context, *Order) (*CreateUpdateDeleteOrderResp, error)
	GetOrders(context.Context, *OrderEmpty) (*GetOrdersResp, error)
	GetOrder(context.Context, *GetDeleteOrderReq) (*GetOrderResp, error)
	UpdateOrder(context.Context, *UpdateOrderReq) (*CreateUpdateDeleteOrderResp, error)
	DeleteOrder(context.Context, *GetDeleteOrderReq) (*CreateUpdateDeleteOrderResp, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *Order) (*CreateUpdateDeleteOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrders(context.Context, *OrderEmpty) (*GetOrdersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrderServiceServer) GetOrder(context.Context, *GetDeleteOrderReq) (*GetOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrder(context.Context, *UpdateOrderReq) (*CreateUpdateDeleteOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderServiceServer) DeleteOrder(context.Context, *GetDeleteOrderReq) (*CreateUpdateDeleteOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrders(ctx, req.(*OrderEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeleteOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*GetDeleteOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder(ctx, req.(*UpdateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeleteOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteOrder(ctx, req.(*GetDeleteOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OrderService_GetOrders_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _OrderService_UpdateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrderService_DeleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/order.proto",
}
