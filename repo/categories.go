package repo

import (
	"dienlanhphongvan/infra"
	"dienlanhphongvan/models"
	"utilities/uerror"
)

type category struct {
	base
}

var Category ICategory

func init() {
	Category = category{}
}

type ICategory interface {
	GetAll() (categories []models.Category, err error)
	GetList(limit, offset int) (categories []models.Category, err error)
	GetById(id int) (category models.Category, err error)
}

func (p category) Create(category *models.Category) error {
	return p.create(category)
}

func (p category) Update(category *models.Category) error {
	return p.save(category)
}

func (category) GetAll() (categories []models.Category, err error) {
	err = infra.PostgreSql.Model(models.Category{}).
		Find(&categories).
		Error
	if err != nil {
		err = uerror.StackTrace(err)
		return
	}

	return categories, nil
}

func (category) GetList(limit, offset int) (categories []models.Category, err error) {
	err = infra.PostgreSql.Model(models.Category{}).
		Offset(offset).
		Limit(limit).
		Find(&categories).
		Error
	if err != nil {
		err = uerror.StackTrace(err)
		return
	}

	return categories, nil
}

func (category) GetById(id int) (category models.Category, err error) {
	err = infra.PostgreSql.Model(models.Category{}).
		Where("id = ?", id).
		Find(&category).
		Error
	if err != nil {
		err = uerror.StackTrace(err)
		return
	}

	return category, nil
}
