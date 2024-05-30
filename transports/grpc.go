package transports

import (
	"context"

	"github.com/Venukishore-R/microservice2_place_orders/endpoints"
	"github.com/Venukishore-R/microservice2_place_orders/models"
	"github.com/Venukishore-R/microservice2_place_orders/protos"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type MyServer struct {
	createOrder grpc.Handler
	getOrders   grpc.Handler
	getOrder    grpc.Handler
	updateOrder grpc.Handler
	deleteOrder grpc.Handler
	protos.UnimplementedOrderServiceServer
}

func NewMyServer(logger log.Logger, endpoints endpoints.Endpoints) MyServer {
	options := grpc.ServerBefore(jwt.GRPCToContext())
	return MyServer{
		createOrder: grpc.NewServer(
			endpoints.CreateOrder,
			decodeCreateOrderReq,
			encodeCreateOrderResp,
			options,
		),
		getOrders: grpc.NewServer(
			endpoints.GetOrders,
			decodeGetOrdersReq,
			encodeGetOrdersResp,
			options,
		),
		getOrder: grpc.NewServer(
			endpoints.GetOrder,
			decodeGetOrderReq,
			encodeGetOrderResp,
			options,
		),
		updateOrder: grpc.NewServer(
			endpoints.UpdateOrder,
			decodeUpdateOrderReq,
			encodeUpdateOrderResp,
			options,
		),
		deleteOrder: grpc.NewServer(
			endpoints.DeleteOrder,
			decodeDeleteOrderReq,
			encodeDeleteOrderResp,
			options,
		),
	}
}

func (s *MyServer) CreateOrder(ctx context.Context, request *protos.Order) (*protos.CreateUpdateDeleteOrderResp, error) {
	_, resp, err := s.createOrder.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*protos.CreateUpdateDeleteOrderResp), nil
}

func decodeCreateOrderReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protos.Order)
	return endpoints.CreateOrderReq{Name: req.Name, Quantity: req.Quantity, Price: req.Price}, nil
}

func encodeCreateOrderResp(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.CreateUpdateDeleteOrderResp)
	return &protos.CreateUpdateDeleteOrderResp{Status: resp.Status}, nil
}

func (s *MyServer) GetOrders(ctx context.Context, request *protos.OrderEmpty) (*protos.GetOrdersResp, error) {
	_, resp, err := s.getOrders.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*protos.GetOrdersResp), nil
}

func decodeGetOrdersReq(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(*protos.OrderEmpty)
	return endpoints.Empty{}, nil
}

func encodeGetOrdersResp(_ context.Context, response interface{}) (interface{}, error) {
	var orders []*protos.Order

	resp := response.(endpoints.GetOrdersResp)

	for _, order := range resp.Order {
		orderS := &protos.Order{
			Name:     order.Name,
			Quantity: order.Quantity,
			Price:    order.Price,
		}
		orders = append(orders, orderS)
	}
	return &protos.GetOrdersResp{Order: orders}, nil
}

func (s *MyServer) GetOrder(ctx context.Context, request *protos.GetDeleteOrderReq) (*protos.GetOrderResp, error) {
	_, resp, err := s.getOrder.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*protos.GetOrderResp), nil
}

func decodeGetOrderReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protos.GetDeleteOrderReq)
	return endpoints.GetDeleteOrderReq{Id: req.Id}, nil
}

func encodeGetOrderResp(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GetOrderResp)
	orderS := &protos.Order{
		Name:     resp.Order.Name,
		Quantity: resp.Order.Quantity,
		Price:    resp.Order.Price,
	}
	return &protos.GetOrderResp{Order: orderS}, nil
}

func (s *MyServer) UpdateOrder(ctx context.Context, request *protos.UpdateOrderReq) (*protos.CreateUpdateDeleteOrderResp, error) {
	_, resp, err := s.updateOrder.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*protos.CreateUpdateDeleteOrderResp), nil
}

func decodeUpdateOrderReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protos.UpdateOrderReq)

	orderS := &models.Order{
		Name:     req.Order.Name,
		Quantity: req.Order.Quantity,
		Price:    req.Order.Price,
	}

	return endpoints.UpdateOrderReq{Id: req.Id, Order: orderS}, nil
}

func encodeUpdateOrderResp(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.CreateUpdateDeleteOrderResp)

	return &protos.CreateUpdateDeleteOrderResp{Status: resp.Status}, nil
}

func (s *MyServer) DeleteOrder(ctx context.Context, request *protos.GetDeleteOrderReq) (*protos.CreateUpdateDeleteOrderResp, error) {
	_, resp, err := s.deleteOrder.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*protos.CreateUpdateDeleteOrderResp), nil
}

func decodeDeleteOrderReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protos.GetDeleteOrderReq)
	return endpoints.GetDeleteOrderReq{Id: req.Id}, nil
}

func encodeDeleteOrderResp(_ context.Context, resposne interface{}) (interface{}, error) {
	resp := resposne.(endpoints.CreateUpdateDeleteOrderResp)
	return &protos.CreateUpdateDeleteOrderResp{Status: resp.Status}, nil
}
