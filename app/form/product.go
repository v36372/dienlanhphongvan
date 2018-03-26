package form

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/validator.v2"
)

type Product struct {
	Name        string  `form:"name" json:"name" validator:"required"`
	Price       float32 `form:"price" json:"price" validator:"required"`
	Description string  `form:"desc" json:"desc" validator:"required"`
	CategoryId  int     `form:"categoryId" json:"categoryId" validator:"required"`
	Image01     string  `form:"image01" json:"image01" validator:"required"`
	Image02     string  `form:"image02" json:"image02" validator:"required"`
	Image03     string  `form:"image03" json:"image03" validator:"required"`
	Image04     string  `form:"image04" json:"image04" validator:"required"`
	Image05     string  `form:"image05" json:"image05" validator:"required"`
	Image06     string  `form:"image06" json:"image06" validator:"required"`
}

func (inputForm *Product) FromCtx(c *gin.Context) error {
	if err := c.Bind(inputForm); err != nil {
		return uer.BadRequestError(err)
	}

	if err := inputForm.Validate(); err != nil {
		return uer.BadRequestError(err)
	}

	return nil
}
func (inputForm *Product) Validate() error {
	if errs := validator.Validate(inputForm); errs != nil {
		return errs
	}

	return nil
}

func (inputForm *Product) ToModelDb() models.Product {
	return models.Product{
		Name:        inputForm.Name,
		Price:       inputForm.Price,
		Description: inputForm.Description,
		CategoryId:  inputForm.CategoryId,
		Image01:     inputForm.Image01,
		Image02:     inputForm.Image02,
		Image03:     inputForm.Image03,
		Image04:     inputForm.Image04,
		Image05:     inputForm.Image05,
		Image06:     inputForm.Image06,
	}
}
