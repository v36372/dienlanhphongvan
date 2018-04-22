package handler

import (
	"dienlanhphongvan/app/presenter"
	"dienlanhphongvan/middleware"

	"github.com/gin-gonic/gin"
)

type contactHandler struct {
}

func (h contactHandler) ContactPage(c *gin.Context) {
	admin := middleware.Auth.GetCurrentUser(c)

	contactPresenter := presenter.NewContactPagePresenter(admin != nil)
	c.HTML(200, "contact", contactPresenter)
}
