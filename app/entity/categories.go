package entity

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/repo"
	"dienlanhphongvan/utilities/uer"
)

type categoryEntity struct{}

type Category interface {
	GetForHomePage(limit int) ([]models.Category, error)
	GetForDashboard() (categories []models.Category, err error)
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

func (categoryEntity) GetForDashboard() (categories []models.Category, err error) {
	categories, err = repo.Category.GetAll()
	if err != nil {
		err = uer.InternalError(err)
		return
	}

	return
}
