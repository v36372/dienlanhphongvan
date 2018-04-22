package presenter

import (
	"dienlanhphongvan/app/view"
	"fmt"
)

type ProductPage struct {
	global
	Product view.Product
}

type ProductsPage struct {
	global
	Products []view.Product
}

func NewProductPagePresenter(product view.Product, isAdmin bool) ProductPage {
	return ProductPage{
		global: global{
			Categories: globalCategories,
			IsAdmin:    isAdmin,
			CurrentPageBreadCrumbs: []string{
				product.Category,
				product.Name,
			},
			CurrentPageTitle: fmt.Sprintf("%s - %s", websiteName, product.Name),
		},
		Product: product,
	}
}

func NewProductsPagePresenter(products []view.Product, categoryName string, isAdmin bool) ProductsPage {
	return ProductsPage{
		global: global{
			Categories: globalCategories,
			IsAdmin:    isAdmin,
			CurrentPageBreadCrumbs: []string{
				categoryName,
			},
			CurrentPageTitle: fmt.Sprintf("%s - %s", websiteName, categoryName),
		},
		Products: products,
	}
}
