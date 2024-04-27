package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"storage/models"
)

type IProductRepository interface {
	Get(ctx context.Context, id uint) (*models.Product, error)
	Create(ctx context.Context, product []*models.Product) error
}

type ProductRepository struct {
	postgres *gorm.DB
	redis    redis.UniversalClient
}

func (p *ProductRepository) Create(ctx context.Context, product []*models.Product) error {
	return p.postgres.WithContext(ctx).Create(&product).Error
}

func (p *ProductRepository) Get(ctx context.Context, id uint) (*models.Product, error) {
	var product *models.Product
	err := p.postgres.WithContext(ctx).Where("id =?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func NewProductRepository(postgres *gorm.DB, redis redis.UniversalClient) *ProductRepository {
	return &ProductRepository{postgres: postgres, redis: redis}
}
