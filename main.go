package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"storage/handler"
	"storage/repository"
	"storage/repository/database"
	service2 "storage/service"
)

func main() {
	fmt.Println("init program...")
	postgresDb := database.ConnectPostgres()
	redis := database.ConnectToRedis()

	pRepo := repository.NewProductRepository(postgresDb, redis)
	service := service2.NewProductService(pRepo)
	//logger := InitLogger()
	InitRest(service)
}

func InitRest(service *service2.ProductService) {

	echo := echo.New()

	h := handler.NewProductHandler(service)

	echo.GET("promotions/id", h.GetProducts)
	echo.Logger.Fatal(echo.Start(":3000"))
}
