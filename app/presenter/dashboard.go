package presenter

import (
	"dienlanhphongvan/app/view"
	"fmt"
)

type DashboardPagePresenter struct {
	global
}

type DashboardProductListPresenter struct {
	global
	Products []view.Product
}

func NewDashboardPagePresenter(isAdmin bool, br string) DashboardPagePresenter {
	return DashboardPagePresenter{
		global: global{
			Categories: globalCategories,
			IsAdmin:    isAdmin,
			CurrentPageBreadCrumbs: []string{
				br,
			},
			CurrentPageTitle: fmt.Sprintf("%s - %s", websiteName, br),
		},
	}
}

func NewDashboardProductListPresenter(products []view.Product, isAdmin bool) DashboardProductListPresenter {
	return DashboardProductListPresenter{
		global: global{
			Categories: globalCategories,
			IsAdmin:    isAdmin,
			CurrentPageBreadCrumbs: []string{
				"Danh sách sản phẩm",
			},
			CurrentPageTitle: fmt.Sprintf("%s - %s", websiteName, "Danh sách sản phẩm"),
		},
		Products: products,
	}
}
