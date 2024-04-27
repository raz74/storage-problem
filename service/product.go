package service

import (
	"context"
	"storage/models"
	"storage/repository"
)

type ProductService struct {
	repo repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProduct(ctx context.Context, req models.ProductReq) (*models.ProductResp, error) {
	product, err := s.repo.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return mapProductToDto(product), nil
}
func mapProductToDto(product *models.Product) *models.ProductResp {
	var p = &models.ProductResp{
		Id:             product.Id,
		Price:          product.Price,
		ExpirationData: product.ExpirationData,
	}
	return p
}
