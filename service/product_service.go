package service

import (
	"clean-architecture/model"
)
type ProductService interface {
	Create(request model.CreateProductRequest)(response model.CreateProductRequest)
	list() (response []model.GetProductResponse)
}