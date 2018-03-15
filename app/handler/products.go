package handler

import (
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/app/form"
	"dienlanhphongvan/app/params"
	"dienlanhphongvan/app/view"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productEntity entity.Product
}

func (h productHandler) GetDetail(c *gin.Context) {
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

	productView := view.NewProduct(*product)
	c.HTML(200, "product-detail.html", productView)
}

func (h productHandler) GetList(c *gin.Context) {
	limit, offset, page := params.NewGetProductsParams(c)

	products, total, err := h.productEntity.GetList(limit, offset)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	productsView := view.NewProducts(products)
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
	err = h.productEntity.Create(productModelDb)
	if err != nil {
		uer.HandleErrorGin(err, c)
	}

	productView := view.NewProduct(productModelDb)
	view.ResponseOK(c, productView)
}
