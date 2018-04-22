package handler

import (
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/app/presenter"
	"dienlanhphongvan/app/view"
	"dienlanhphongvan/middleware"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
)

type indexHandler struct {
	Category entity.Category
	Product  entity.Product
}

const (
	limitCategoriesHomePage = 100
	limitProductsHomePage   = 12
)

func (h indexHandler) Index(c *gin.Context) {
	admin := middleware.Auth.GetCurrentUser(c)

	products, err := h.Product.GetNewest(limitProductsHomePage)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}
	productsView, err := view.NewProducts(products)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	indexPresenter := presenter.NewIndexPagePresenter(productsView, admin != nil)
	c.HTML(200, "index", indexPresenter)
}
