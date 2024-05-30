package endpoints

import "github.com/Venukishore-R/microservice2_place_orders/models"

type Empty struct{}

type CreateOrderReq struct {
	Name     string
	Quantity string
	Price    string
}

type CreateUpdateDeleteOrderResp struct {
	Status int64
}

type GetOrdersResp struct {
	Order []*models.Order
}

type GetDeleteOrderReq struct {
	Id int64
}

type GetOrderResp struct {
	Order *models.Order
}

type UpdateOrderReq struct {
	Id    int64
	Order *models.Order
}
