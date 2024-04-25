package endpoints

import (
	"github.com/Venukishore-R/microservice2_place_orders/services"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type Endpoints struct {
	CreateOrder endpoint.Endpoint
}

func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		CreateOrder: makeCreateOrderEndpoint(s),
	}
}

func makeCreateOrderEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateOrderReq)
		status, err := s.CreateOrder(ctx, req.Name, req.Quantity, req.Price)
		return CreateOrderResp{Status: status}, err
	}
}
