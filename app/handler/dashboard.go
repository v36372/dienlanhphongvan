package handler

import (
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/app/params"
	"dienlanhphongvan/app/presenter"
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
	dashboardPagePresenter := presenter.NewDashboardPagePresenter(admin != nil, "Tạo sản phẩm")
	c.HTML(200, "create-product", dashboardPagePresenter)
}

func (h dashboardHandler) UpdateProduct(c *gin.Context) {
	admin := middleware.Auth.GetCurrentUser(c)

	productSlug := params.NewGetProductSlugParam(c)
	product, err := h.product.GetBySlug(productSlug)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	if product == nil {
		uer.HandleNotFound(c)
		return
	}

	productView, err := view.NewProduct(*product)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	dashboardUpdateProductPresenter := presenter.NewDashboardUpdateProductPresenter(productView, admin != nil)
	c.HTML(200, "update-product", dashboardUpdateProductPresenter)
}

func (h dashboardHandler) CreateCategory(c *gin.Context) {
	admin := middleware.Auth.GetCurrentUser(c)
	dashboardPagePresenter := presenter.NewDashboardPagePresenter(admin != nil, "Tạo phan loại")
	c.HTML(200, "create-category", dashboardPagePresenter)
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

	dashboardProductList := presenter.NewDashboardProductListPresenter(productViews, admin != nil)
	c.HTML(200, "list-product", dashboardProductList)
}
