package handler

import (
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/app/view"
	"dienlanhphongvan/middleware"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
)

type dashboardHandler struct {
	product  entity.Product
	category entity.Category
	image    entity.Image
}

func (h dashboardHandler) CreateProduct(c *gin.Context) {
	admin := middleware.Auth.GetCurrentUser(c)
	categories, err := h.category.GetForDashboard()
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}
	categoriesView := view.NewCategoriesForDashboard(categories)
	createProductPageView := struct {
		Categories []view.Category
		IsAdmin    bool
	}{
		Categories: categoriesView,
		IsAdmin:    admin != nil,
	}
	c.HTML(200, "create-product", createProductPageView)
}

func (h dashboardHandler) CreateCategory(c *gin.Context) {
	admin := middleware.Auth.GetCurrentUser(c)
	createCategoryPageView := struct {
		IsAdmin bool
	}{
		IsAdmin: admin != nil,
	}
	c.HTML(200, "create-category", createCategoryPageView)
}

func (h dashboardHandler) ListProduct(c *gin.Context) {
	admin := middleware.Auth.GetCurrentUser(c)

	limit := 100
	offset := 0
	products, _, err := h.product.GetList(limit, offset)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	productViews, err := view.NewProducts(products)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	productListPageView := struct {
		Products []view.Product
		IsAdmin  bool
	}{
		Products: productViews,
		IsAdmin:  admin != nil,
	}
	c.HTML(200, "list-product", productListPageView)
}
