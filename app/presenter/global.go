package presenter

import (
	"dienlanhphongvan/app/view"
	"dienlanhphongvan/repo"
	"dienlanhphongvan/utilities/ulog"
)

const (
	limitCategoriesHomePage = 100
	websiteName             = "Điện lạnh Phong Vân"
)

type global struct {
	Categories             []view.Category
	CurrentPageTitle       string
	CurrentPageBreadCrumbs []string
	IsAdmin                bool
}

func getGlobalCategories() []view.Category {
	categories, err := repo.Category.GetList(limitCategoriesHomePage, 0)
	if err != nil {
		ulog.Logger().LogErrorObjectManual(err, "error when get global categories", nil)
		return []view.Category{}
	}

	globalCategories := view.NewCategories(categories)
	if err != nil {
		ulog.Logger().LogErrorObjectManual(err, "error when populate view global categories", nil)
		return []view.Category{}
	}

	return globalCategories
}
