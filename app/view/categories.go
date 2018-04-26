package view

import (
	"dienlanhphongvan/models"
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

func NewCategory(category models.Category) Category {
	return Category{
		Name: category.Name,
		Url:  fmt.Sprintf("/categories/%s", category.Slug),
	}
}

func NewCategoriesForDashboard(categories []models.Category) (categoryViews []Category) {
	categoryViews = make([]Category, len(categories))
	for i, cate := range categories {
		categoryViews[i] = NewCategoryForDashboard(cate)
	}

	return
}
func NewCategories(categories []models.Category) (categoryViews []Category) {
	categoryViews = make([]Category, len(categories))
	for i, cate := range categories {
		categoryViews[i] = NewCategory(cate)
	}

	return
}
