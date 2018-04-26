package presenter

import (
	"dienlanhphongvan/app/view"
	"fmt"
)

type DashboardPagePresenter struct {
	global
}

type DashboardUpdateProductPresenter struct {
	global
	Product view.Product
}

type DashboardProductListPresenter struct {
	global
	Products []view.Product
}

func NewDashboardPagePresenter(isAdmin bool, br string) DashboardPagePresenter {
	return DashboardPagePresenter{
		global: global{
			Categories: getGlobalCategories(),
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
			Categories: getGlobalCategories(),
			IsAdmin:    isAdmin,
			CurrentPageBreadCrumbs: []string{
				"Danh sách sản phẩm",
			},
			CurrentPageTitle: fmt.Sprintf("%s - %s", websiteName, "Danh sách sản phẩm"),
		},
		Products: products,
	}
}

func NewDashboardUpdateProductPresenter(product view.Product, isAdmin bool) DashboardUpdateProductPresenter {
	return DashboardUpdateProductPresenter{
		global: global{
			Categories: getGlobalCategories(),
			IsAdmin:    isAdmin,
			CurrentPageBreadCrumbs: []string{
				"Cập nhật sản phẩm",
			},
			CurrentPageTitle: fmt.Sprintf("%s - %s", websiteName, "Cập nhật sản phẩm"),
		},
		Product: product,
	}
}
