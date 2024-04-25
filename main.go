package main

import (
	"fmt"
	"github.com/Venukishore-R/microservice2_place_orders/endpoints"
	"github.com/Venukishore-R/microservice2_place_orders/middlewares"
	"github.com/Venukishore-R/microservice2_place_orders/models"
	"github.com/Venukishore-R/microservice2_place_orders/protos"
	"github.com/Venukishore-R/microservice2_place_orders/services"
	"github.com/Venukishore-R/microservice2_place_orders/transports"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	var logger log.Logger

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		logger.Log("during", "db", "err", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		logger.Log("during", "db", "err", err)
		os.Exit(1)
	}
}

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := services.NewLoggerService(logger)
	makeEndpoints := endpoints.MakeEndpoints(service)
	makeEndpoints.CreateOrder = middlewares.Mid()(makeEndpoints.CreateOrder)
	server := transports.NewMyServer(logger, makeEndpoints)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":5010")
	if err != nil {
		logger.Log("during", "listen", "err", err)
		os.Exit(1)
	}

	go func() {
		serverRegistrar := grpc.NewServer()
		protos.RegisterOrderServiceServer(serverRegistrar, &server)
		level.Info(logger).Log("msg", "server created successfully")
		serverRegistrar.Serve(grpcListener)
	}()

	logger.Log("exit", <-errs)

}
