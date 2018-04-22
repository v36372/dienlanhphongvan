package handler

import (
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/app/form"
	"dienlanhphongvan/app/params"
	"dienlanhphongvan/app/presenter"
	"dienlanhphongvan/app/view"
	"dienlanhphongvan/middleware"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productEntity entity.Product
	imageEntity   entity.Image
}

func (h productHandler) GetByCategory(c *gin.Context) {
	admin := middleware.Auth.GetCurrentUser(c)

	categorySlug := params.NewGetProductSlugParam(c)
	products, categoryName, err := h.productEntity.GetByCategorySlug(categorySlug)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	productsView, err := view.NewProducts(products)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	productsPagePresenter := presenter.NewProductsPagePresenter(productsView, categoryName, admin != nil)
	c.HTML(200, "product-list", productsPagePresenter)
}

func (h productHandler) GetDetail(c *gin.Context) {
	admin := middleware.Auth.GetCurrentUser(c)

	productSlug := params.NewGetProductSlugParam(c)
	product, err := h.productEntity.GetBySlug(productSlug)
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

	productPagePresenter := presenter.NewProductPagePresenter(productView, admin != nil)
	c.HTML(200, "product-detail", productPagePresenter)
}

func (h productHandler) GetList(c *gin.Context) {
	limit, offset, page := params.NewGetProductsParams(c)

	products, total, err := h.productEntity.GetList(limit, offset)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	productsView, err := view.NewProducts(products)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}
	pagination := view.NewPagination(total, limit, page)
	view.ResponseOKWithPagination(c, productsView, &pagination)
}

func (h productHandler) Create(c *gin.Context) {
	var productForm form.Product
	err := productForm.FromCtx(c)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	productModelDb := productForm.ToModelDb()
	err = h.productEntity.Create(productModelDb, h.imageEntity)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	c.Redirect(302, "/dashboard/product-list")
}
