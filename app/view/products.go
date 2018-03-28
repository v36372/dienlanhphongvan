package view

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/repo"
	"dienlanhphongvan/utilities/uer"
	"fmt"

	"github.com/leekchan/accounting"
)

type Product struct {
	Name        string   `json:"name"`
	Description string   `json:"desc"`
	Category    string   `json:"category"`
	Price       string   `json:"price"`
	Slug        string   `json:"slug"`
	Url         string   `json:"url"`
	Thumbnail   string   `json:"thumbnail"`
	Images      []string `json:"images"`
}

func NewProduct(product models.Product) (Product, error) {
	category, err := repo.Category.GetById(product.CategoryId)
	if err != nil {
		return Product{}, uer.InternalError(err)
	}

	ac := accounting.Accounting{
		Symbol:   "đ",
		Thousand: ".",
		Format:   "%v %s",
	}

	return Product{
		Name:        product.Name,
		Description: product.Description,
		Category:    category.Name,
		Price:       ac.FormatMoney(product.Price),
		Slug:        product.Slug,
		Url:         fmt.Sprintf("products/%s", product.Slug),
		Thumbnail:   NewImage(product.Thumbnail),
		Images: []string{
			NewImage(product.Image01),
			NewImage(product.Image02),
			NewImage(product.Image03),
			NewImage(product.Image04),
			NewImage(product.Image05),
			NewImage(product.Image06),
		},
	}, nil
}

func NewProducts(products []models.Product) (productsView []Product, err error) {
	productsView = make([]Product, len(products))
	for i, product := range products {
		productsView[i], err = NewProduct(product)
		if err != nil {
			return
		}
	}

	return
}