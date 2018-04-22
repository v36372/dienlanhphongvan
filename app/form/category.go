package form

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	validator "gopkg.in/validator.v2"
)

type Category struct {
	Name string `form:"name" json:"name" validator:"required"`
}

func (inputForm *Category) FromCtx(c *gin.Context) error {
	if err := c.Bind(inputForm); err != nil {
		return uer.BadRequestError(err)
	}

	if err := inputForm.Validate(); err != nil {
		return uer.BadRequestError(err)
	}

	return nil
}
func (inputForm *Category) Validate() error {
	if errs := validator.Validate(inputForm); errs != nil {
		return errs
	}

	return nil
}

func (inputForm *Category) ToModelDb() models.Category {
	return models.Category{
		Name: inputForm.Name,
		Slug: slug.Make(inputForm.Name),
	}
}
