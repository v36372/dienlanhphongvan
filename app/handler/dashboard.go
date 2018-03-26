package handler

import (
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/app/view"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
)

type dashboardHandler struct {
	product  entity.Product
	category entity.Category
	image    entity.Image
}

func (h dashboardHandler) CreateProduct(c *gin.Context) {
	categories, err := h.category.GetForDashboard()
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}
	categoriesView := view.NewCategoriesForDashboard(categories)
	createProductPageView := struct {
		Categories []view.Category
	}{
		Categories: categoriesView,
	}
	c.HTML(200, "create-product.html", createProductPageView)
}
