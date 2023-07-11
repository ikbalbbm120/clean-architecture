package repository

import (
	"clean-architecture/entity"
)

type ProductRepository interface {
	Insert(product entity.Product)

	FindAll() (products []entity.Product)

	DeleteAll()
}