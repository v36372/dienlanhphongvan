package entity

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/repo"
	"dienlanhphongvan/utilities/uer"
	"errors"

	"github.com/gosimple/slug"
)

type productEntity struct{}

type Product interface {
	GetBySlug(slug string) (*models.Product, error)
	GetList(limit, offset int) (products []models.Product, total int, err error)
	Create(product models.Product, imgx Image) (err error)
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

func (productEntity) Create(product models.Product, imgx Image) (err error) {
	product.Slug = slug.Make(product.Name)

	uploadImages := []string{
		product.Image01,
		product.Image02,
		product.Image03,
		product.Image04,
		product.Image05,
	}
	originalImages, err := imgx.MoveImagesOfProduct(uploadImages)
	if err != nil {
		return uer.InternalError(err)
	}

	product.Thumbnail = originalImages[0]
	product.Image01 = originalImages[0]
	product.Image02 = originalImages[1]
	product.Image03 = originalImages[2]
	product.Image04 = originalImages[3]
	product.Image05 = originalImages[4]

	err = repo.Product.Create(&product)
	if err != nil {
		err = uer.InternalError(err)
		return
	}

	return
}
