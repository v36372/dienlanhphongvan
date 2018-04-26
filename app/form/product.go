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
	Active      bool    `form:"active" json:"active" validator:"nonzero"`
	Image0      string  `form:"image0" json:"image01" validator:"nonzero"`
	Image1      string  `form:"image1" json:"image02" validator:"nonzero"`
	Image2      string  `form:"image2" json:"image03" validator:"nonzero"`
	Image3      string  `form:"image3" json:"image04" validator:"nonzero"`
	Image4      string  `form:"image4" json:"image05" validator:"nonzero"`
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
		Active:      inputForm.Active,
		Image01:     inputForm.Image0,
		Image02:     inputForm.Image1,
		Image03:     inputForm.Image2,
		Image04:     inputForm.Image3,
		Image05:     inputForm.Image4,
	}
}
