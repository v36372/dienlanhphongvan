package form

import (
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	Username string `form:"username" json:"username" validator:"required"`
	Password string `form:"password" json:"password" validator:"required"`
}

func (inputForm *UserLogin) FromCtx(c *gin.Context) error {
	if err := c.Bind(inputForm); err != nil {
		return uer.BadRequestError(err)
	}

	return nil
}
