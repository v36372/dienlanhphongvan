package entity

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/repo"
	"dienlanhphongvan/utilities/uer"
	"errors"
)

type productEntity struct{}

type Product interface {
	GetBySlug(slug string) (*models.Product, error)
	GetList(limit, offset int) (products []models.Product, total int, err error)
	Create(product models.Product) (err error)
}

func NewProduct() Product {
	return &productEntity{}
}

func (productEntity) GetBySlug(slug string) (*models.Product, error) {
	product, err := repo.Product.GetBySlug(slug)
	if err != nil {
		return product, uer.InternalError(err)
	}

	if product == nil {
		return product, uer.NotFoundError(errors.New("product not found"))
	}

	return product, nil
}

func (productEntity) GetList(limit, offset int) (products []models.Product, total int, err error) {
	products, total, err = repo.Product.GetList(limit, offset)
	if err != nil {
		err = uer.InternalError(err)
		return
	}

	return
}

func (productEntity) Create(product models.Product) (err error) {
	err = repo.Product.Create(&product)
	if err != nil {
		err = uer.InternalError(err)
		return
	}

	return
}
