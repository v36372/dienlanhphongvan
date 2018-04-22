package repo

import (
	"dienlanhphongvan/infra"
	"dienlanhphongvan/models"
	"utilities/uerror"

	"github.com/jinzhu/gorm"
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
	GetById(id int) (*models.Category, error)
	GetBySlug(slug string) (*models.Category, error)
	Create(*models.Category) error
	Update(*models.Category) error
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

func (category) GetBySlug(slug string) (*models.Category, error) {
	var category models.Category
	err := infra.PostgreSql.Model(models.Category{}).
		Where("slug = ?", slug).
		Find(&category).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &category, err
}
func (category) GetById(id int) (*models.Category, error) {
	var category models.Category
	err := infra.PostgreSql.Model(models.Category{}).
		Where("id = ?", id).
		Find(&category).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &category, err
}
