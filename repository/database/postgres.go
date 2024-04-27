package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"storage/config"
	"storage/models"
)

func ConnectPostgres() *gorm.DB {
	dns := getPostgresDns()
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//migrate tables
	err = db.AutoMigrate(models.Product{})
	fmt.Println("Successfully connected!")
	return db
}

func getPostgresDns() string {
	cfg := config.GetPostgres()
	dns := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)
	return dns
}
