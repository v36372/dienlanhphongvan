package handler

import (
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/app/form"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	category entity.Category
}

func (h categoryHandler) Create(c *gin.Context) {
	var categoryForm form.Category
	err := categoryForm.FromCtx(c)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	categoryModelDb := categoryForm.ToModelDb()
	err = h.category.Create(&categoryModelDb)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	c.Redirect(302, "/dashboard/product-list")
}
