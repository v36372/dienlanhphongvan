package repo

import (
	"dienlanhphongvan/infra"
	"dienlanhphongvan/models"
	"dienlanhphongvan/utilities/uerror"

	"github.com/jinzhu/gorm"
)

type product struct {
	base
}

var Product IProduct

func init() {
	Product = product{}
}

type IProduct interface {
	GetBySlug(slug string) (*models.Product, error)
	GetList(limit, offset int) (products []models.Product, total int, err error)
	GetByCategory(category string, limit, offset int) (products []models.Product, total int, err error)
	Create(*models.Product) error
	Update(*models.Product) error
}

func (p product) Create(product *models.Product) error {
	return p.create(product)
}

func (p product) Update(product *models.Product) error {
	return p.save(product)
}

func (product) GetBySlug(slug string) (*models.Product, error) {
	var product models.Product
	if slug == "" {
		return nil, nil
	}
	err := infra.PostgreSql.Model(models.Product{}).Where("slug = ?", slug).Limit(1).Find(&product).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &product, err
}

func (product) GetByCategory(category string, limit, offset int) (products []models.Product, total int, err error) {
	query := infra.PostgreSql.Model(models.Product{}).Joins("JOIN categories ON categories.id = products.category_id").
		Where("categories.name = ?", category)

	err = query.Count(&total).Error
	if err != nil {
		err = uerror.StackTrace(err)
		return
	}

	err = query.Order("products.created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&products).
		Error
	if err != nil {
		err = uerror.StackTrace(err)
		return
	}

	return products, total, nil
}

func (product) GetList(limit, offset int) (products []models.Product, total int, err error) {
	query := infra.PostgreSql.Model(models.Product{})

	err = query.Count(&total).Error
	if err != nil {
		err = uerror.StackTrace(err)
		return
	}

	err = query.Order("products.created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&products).
		Error
	if err != nil {
		err = uerror.StackTrace(err)
		return
	}

	return products, total, nil
}
