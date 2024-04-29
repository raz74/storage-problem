package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"storage/config"
	"storage/handler"
	"storage/repository"
	"storage/repository/database"
	service2 "storage/service"
)

func main() {
	fmt.Println("init program...")
	Init()
}

func Init() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	//init data bases
	postgresDb := database.ConnectPostgres()
	redis := database.ConnectToRedis()
	//init repositories
	pRepo := repository.NewProductRepository(postgresDb, redis)
	//init service
	service := service2.NewProductService(pRepo)

	// process csv file
	go func() {
		service.ProcessCsvData()
	}()

	echo := echo.New()
	handler.InitRest(echo, service)
	echo.Logger.Fatal(echo.Start(":3000"))

}
