package form

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/validator.v2"
)

type Product struct {
	Name        string  `form:"name" json:"name" validator:"nonzero"`
	Price       float32 `form:"price" json:"price" validator:"nonzero"`
	Description string  `form:"desc" json:"desc" validator:"nonzero"`
	CategoryId  int     `form:"categoryId" json:"categoryId" validator:"nonzero"`
	Image01     string  `form:"image01" json:"image01" validator:"nonzero"`
	Image02     string  `form:"image02" json:"image02" validator:"nonzero"`
	Image03     string  `form:"image03" json:"image03" validator:"nonzero"`
	Image04     string  `form:"image04" json:"image04" validator:"nonzero"`
	Image05     string  `form:"image05" json:"image05" validator:"nonzero"`
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
	}
}
