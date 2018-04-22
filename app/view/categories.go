package view

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/utilities/uer"
	"fmt"
)

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func NewCategoryForDashboard(category models.Category) Category {
	return Category{
		Id:   category.Id,
		Name: category.Name,
	}
}

func NewCategory(category models.Category) (cate Category, err error) {
	return Category{
		Name: category.Name,
		Url:  fmt.Sprintf("/categories/%s", category.Slug),
	}, nil
}

func NewCategoriesForDashboard(categories []models.Category) (categoryViews []Category) {
	categoryViews = make([]Category, len(categories))
	for i, cate := range categories {
		categoryViews[i] = NewCategoryForDashboard(cate)
	}

	return
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
