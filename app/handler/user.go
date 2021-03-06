package handler

import (
	"dienlanhphongvan/app/entity"
	"dienlanhphongvan/app/form"
	"dienlanhphongvan/app/presenter"
	"dienlanhphongvan/middleware"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	user      entity.User
	secCookie *middleware.SecCookie
}

func (h userHandler) LoginPage(c *gin.Context) {
	admin := middleware.Auth.GetCurrentUser(c)

	if admin != nil {
		c.Redirect(302, "/dashboard/product-list")
		return
	}

	loginPagePresenter := presenter.NewLoginPagePresenter()
	c.HTML(200, "login", loginPagePresenter)
}

func (h userHandler) Login(c *gin.Context) {
	var loginForm form.UserLogin
	err := loginForm.FromCtx(c)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	user, err := h.user.Login(loginForm.Username, loginForm.Password)
	if err != nil {
		uer.HandleErrorGin(err, c)
		return
	}

	if user == nil {
		c.Redirect(200, "/user/login")
		return
	}

	_, err = h.secCookie.SetAuthorizationToken("auth", user.Username, "/", c.Writer)
	if err != nil {
		err = uer.InternalError(err)
		uer.HandleErrorGin(err, c)
	}

	c.Redirect(302, "/dashboard/product-list")
}
