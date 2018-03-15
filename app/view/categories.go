package view

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/repo"
	"dienlanhphongvan/utilities/uer"
)

type Category struct {
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

func NewCategory(category models.Category) (cate Category, err error) {
	limit, offset := 10, 0
	products, _, err := repo.Product.GetByCategory(category.Name, limit, offset)
	if err != nil {
		err = uer.InternalError(err)
		return
	}

	return Category{
		Name:     category.Name,
		Products: NewProducts(products),
	}, nil
}

func NewCategories(categories []models.Category) (categoryViews []Category, err error) {
	categoryViews = make([]Category, len(categories))
	for i, cate := range categories {
		categoryViews[i], err = NewCategory(cate)
		if err != nil {
			err = uer.InternalError(err)
			return
		}
	}

	return
}
