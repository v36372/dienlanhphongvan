package form

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/utilities/uer"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/validator.v2"
)

type Product struct {
	Name  string  `json:"name" validator:"required"`
	Price float32 `json:"price" validator:"required"`
}

func (inputForm *Product) FromCtx(c *gin.Context) error {
	if err := c.Bind(&inputForm); err != nil {
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
		Name:  inputForm.Name,
		Price: inputForm.Price,
	}
}
