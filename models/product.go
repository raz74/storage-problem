package models

import "time"

type ProductReq struct {
	Id uint `json:"id"`
}

type ProductResp struct {
	Id             uint          `json:"id"`
	Price          int           `json:"price"`
	ExpirationData time.Duration `json:"expiration_data"`
}

type Product struct {
	Id             uint
	Price          int
	ExpirationData time.Duration
}
