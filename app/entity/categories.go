package entity

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/repo"
	"dienlanhphongvan/utilities/uer"
)

type categoryEntity struct{}

type Category interface {
	GetForHomePage(limit int) ([]models.Category, error)
}

func NewCategory() Category {
	return &categoryEntity{}
}

func (categoryEntity) GetForHomePage(limit int) (categories []models.Category, err error) {
	offset := 0
	categories, err = repo.Category.GetList(limit, offset)
	if err != nil {
		err = uer.InternalError(err)
		return
	}

	return
}
