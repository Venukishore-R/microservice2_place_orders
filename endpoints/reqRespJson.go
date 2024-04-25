package endpoints

type CreateOrderReq struct {
	Name     string
	Quantity string
	Price    string
}

type CreateOrderResp struct {
	Status int64
}
