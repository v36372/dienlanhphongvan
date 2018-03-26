package handler

import (
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/app/view"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
)

type indexHandler struct {
	Category entity.Category
}

const (
	limitCategoryHomePage = 3
)

func (h indexHandler) Index(c *gin.Context) {
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
	}{
		Categories: categoriesView,
	}

	c.HTML(200, "index.html", homePageView)
}
