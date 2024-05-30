package services

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Venukishore-R/microservice2_place_orders/models"
	"github.com/go-kit/log"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LoggerService struct {
	logger log.Logger
}

type Service interface {
	CreateOrder(ctx context.Context, name, quantity, price string) (int64, error)
	GetOrders(ctx context.Context) ([]*models.Order, error)
	GetOrder(ctx context.Context, id int64) (string, string, string, error)
	UpdateOrder(ctx context.Context, id int64, name, quantity, price string) (int64, error)
	DeleteOrder(ctx context.Context, id int64) (int64, error)
}

func NewLoggerService(logger log.Logger) *LoggerService {
	return &LoggerService{
		logger: logger,
	}
}

func ConnectDb() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *LoggerService) CreateOrder(ctx context.Context, name, quantity, price string) (int64, error) {

	s.logger.Log("received: ", "name", name, "quantity", quantity, "price", price)

	db, err := ConnectDb()
	if err != nil {
		s.logger.Log("error while connecting to db", err)
		return http.StatusInternalServerError, err
	}

	order := &models.Order{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}

	err = db.Model(&models.Order{}).Create(order).Error
	if err != nil {
		s.logger.Log("error while creating order", err)
		return http.StatusInternalServerError, err
	}

	s.logger.Log("Name: ", ctx.Value("Name"), "Email: ", ctx.Value("Email"), "Phone: ", ctx.Value("Phone"))
	return http.StatusOK, nil
}

func (s *LoggerService) GetOrders(ctx context.Context) ([]*models.Order, error) {
	var orders []*models.Order

	db, err := ConnectDb()
	if err != nil {
		s.logger.Log("error while connecting to db", err)
		return nil, err
	}

	err = db.Model(&models.Order{}).Find(&orders).Error
	if err != nil {
		s.logger.Log("cannot able to get orders ", err)
		return nil, err
	}

	return orders, nil
}

func (s *LoggerService) GetOrder(ctx context.Context, id int64) (string, string, string, error) {
	var order *models.Order

	db, err := ConnectDb()
	if err != nil {
		s.logger.Log("error while connecting to db", err)
		return "", "", "", err
	}

	err = db.Model(&models.Order{}).Where("id = ?", id).First(&order).Error
	if err != nil {
		s.logger.Log("cannot get order ", err)
		return "", "", "", err
	}

	return order.Name, order.Quantity, order.Price, nil
}

func (s *LoggerService) UpdateOrder(ctx context.Context, id int64, name, quantity, price string) (int64, error) {

	db, err := ConnectDb()
	if err != nil {
		s.logger.Log("cannot able to connect to db ", err)
		return http.StatusInternalServerError, err
	}

	order := &models.Order{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}

	err = db.Model(&models.Order{}).Where("id = ?", id).Updates(&order).Error
	if err != nil {
		s.logger.Log("unable to update order ", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *LoggerService) DeleteOrder(ctx context.Context, id int64) (int64, error) {

	db, err := ConnectDb()
	if err != nil {
		s.logger.Log("cannot able to connect to db ", err)
		return http.StatusInternalServerError, err
	}

	err = db.Delete(&models.Order{}, id).Error
	if err != nil {
		s.logger.Log("unable to delete order ", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
