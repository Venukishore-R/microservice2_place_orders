package transports

import (
	"context"
	"github.com/Venukishore-R/microservice2_place_orders/endpoints"
	"github.com/Venukishore-R/microservice2_place_orders/protos"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type MyServer struct {
	createOrder grpc.Handler
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
	}
}

func (s *MyServer) CreateOrder(ctx context.Context, request *protos.Order) (*protos.CreateOrderResp, error) {
	_, resp, err := s.createOrder.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*protos.CreateOrderResp), nil
}

func decodeCreateOrderReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protos.Order)
	return endpoints.CreateOrderReq{Name: req.Name, Quantity: req.Quantity, Price: req.Price}, nil
}

func encodeCreateOrderResp(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.CreateOrderResp)
	return &protos.CreateOrderResp{Status: resp.Status}, nil
}
