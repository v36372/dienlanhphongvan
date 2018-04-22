package presenter

import (
	"dienlanhphongvan/app/view"
	"fmt"
)

type IndexPage struct {
	global
	Products []view.Product
}

func NewIndexPagePresenter(products []view.Product, isAdmin bool) IndexPage {
	return IndexPage{
		global: global{
			Categories:       globalCategories,
			IsAdmin:          isAdmin,
			CurrentPageTitle: fmt.Sprintf("%s - %s", websiteName, "Trang chủ"),
		},
		Products: products,
	}
}
