package models

import "time"

type ProductReq struct {
	Id uint `json:"id"`
}

type ProductResp struct {
	Id             string    `json:"id"`
	Price          float64   `json:"price"`
	ExpirationData time.Time `json:"expiration_data"`
}

type Product struct {
	Id             string
	Price          float64
	ExpirationData time.Time
}
