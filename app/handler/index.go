package handler

import (
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/app/view"
	"dienlanhphongvan/middleware"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
)

type indexHandler struct {
	Category entity.Category
}

const (
	limitCategoryHomePage = 10
)

func (h indexHandler) Index(c *gin.Context) {
	admin := middleware.Auth.GetCurrentUser(c)

	categories, err := h.Category.GetForHomePage(limitCategoryHomePage)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	categoriesView, err := view.NewCategories(categories)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	homePageView := struct {
		Categories []view.Category
		IsAdmin    bool
	}{
		Categories: categoriesView,
		IsAdmin:    admin != nil,
	}

	c.HTML(200, "index", homePageView)
}
