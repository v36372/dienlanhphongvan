package view

import (
	"dienlanhphongvan/models"
	"fmt"
)

type Product struct {
	Name      string   `json:"name"`
	Price     float32  `json:"price"`
	Slug      string   `json:"slug"`
	Url       string   `json:"url"`
	Thumbnail string   `json:"thumbnail"`
	Images    []string `json:"images"`
}

func NewProduct(product models.Product) Product {
	return Product{
		Name:      product.Name,
		Price:     product.Price,
		Slug:      product.Slug,
		Url:       fmt.Sprintf("products/%s", product.Slug),
		Thumbnail: product.Thumbnail,
		Images: []string{
			product.Image01,
			product.Image02,
			product.Image03,
			product.Image04,
			product.Image05,
		},
	}
}

func NewProducts(products []models.Product) (productsView []Product) {
	productsView = make([]Product, len(products))
	for i, product := range products {
		productsView[i] = NewProduct(product)
	}

	return
}
