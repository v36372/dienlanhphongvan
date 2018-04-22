package presenter

import (
	"dienlanhphongvan/app/view"
	"dienlanhphongvan/repo"
	"fmt"
)

var (
	globalCategories []view.Category
)

const (
	limitCategoriesHomePage = 100
	websiteName             = "Điện lạnh Phong Van"
)

type global struct {
	Categories             []view.Category
	CurrentPageTitle       string
	CurrentPageBreadCrumbs []string
	IsAdmin                bool
}

func InitGlobalPresenter() {
	categories, err := repo.Category.GetList(limitCategoriesHomePage, 0)
	if err != nil {
		panic(fmt.Errorf("Cant init presenter: %s", err))
	}

	globalCategories, err = view.NewCategories(categories)
	if err != nil {
		panic(fmt.Errorf("Cant init presenter: %s", err))
	}
}
