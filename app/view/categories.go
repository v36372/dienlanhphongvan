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

func NewCategory(category models.Category) Category {
	return Category{
		Id:   category.Id,
		Name: category.Name,
		Url:  fmt.Sprintf("/categories/%s", category.Slug),
	}
}

func NewCategories(categories []models.Category) (categoryViews []Category) {
	categoryViews = make([]Category, len(categories))
	for i, cate := range categories {
		categoryViews[i] = NewCategory(cate)
	}

	return
}
