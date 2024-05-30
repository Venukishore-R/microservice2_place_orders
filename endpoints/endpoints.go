package endpoints

import (
	"github.com/Venukishore-R/microservice2_place_orders/models"
	"github.com/Venukishore-R/microservice2_place_orders/services"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type Endpoints struct {
	CreateOrder endpoint.Endpoint
	GetOrders   endpoint.Endpoint
	GetOrder    endpoint.Endpoint
	UpdateOrder endpoint.Endpoint
	DeleteOrder endpoint.Endpoint
}

func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		CreateOrder: makeCreateOrderEndpoint(s),
		GetOrders:   makeGetOrdersEndpoint(s),
		GetOrder:    makeGetOrderEndpoint(s),
		UpdateOrder: makeUpdateOrderEndpoint(s),
		DeleteOrder: makeDeleteOrderEndpoint(s),
	}
}

func makeCreateOrderEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateOrderReq)
		status, err := s.CreateOrder(ctx, req.Name, req.Quantity, req.Price)
		return CreateUpdateDeleteOrderResp{Status: status}, err
	}
}

func makeGetOrdersEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		orders, err := s.GetOrders(ctx)
		return GetOrdersResp{Order: orders}, err
	}
}

func makeGetOrderEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDeleteOrderReq)
		name, quantity, price, err := s.GetOrder(ctx, req.Id)
		order := &models.Order{
			Name:     name,
			Quantity: quantity,
			Price:    price,
		}
		return GetOrderResp{Order: order}, err
	}
}

func makeUpdateOrderEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateOrderReq)
		status, err := s.UpdateOrder(ctx, req.Id, req.Order.Name, req.Order.Quantity, req.Order.Price)
		return CreateUpdateDeleteOrderResp{Status: status}, err
	}
}

func makeDeleteOrderEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDeleteOrderReq)
		status, err := s.DeleteOrder(ctx, req.Id)
		return CreateUpdateDeleteOrderResp{Status: status}, err
	}
}
