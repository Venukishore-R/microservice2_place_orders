package models

type Order struct {
	Id       int64  `json:"id" gorm:"primary_key;autoIncrement"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	UserId   int64  `json:"user_id" gorm:"default:null"`
}
