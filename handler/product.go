package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"storage/models"
	"storage/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (s *ProductHandler) GetProducts(c echo.Context) error {
	req := models.ProductReq{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	product, err := s.service.GetProduct(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"result": product,
	})
}
