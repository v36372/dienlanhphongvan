package repo

import (
	"dienlanhphongvan/infra"
	"dienlanhphongvan/models"

	"github.com/jinzhu/gorm"
)

type user struct {
	base
}

var User IUser

func init() {
	User = user{}
}

type IUser interface {
	GetByUsername(username string) (*models.User, error)
	Create(*models.User) error
	Update(*models.User) error
}

func (u user) Create(user *models.User) error {
	return u.create(user)
}

func (u user) Update(user *models.User) error {
	return u.save(user)
}

func (user) GetByUsername(username string) (*models.User, error) {
	var user models.User

	err := infra.PostgreSql.Model(models.User{}).Where("username = ?", username).Limit(1).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &user, err
}
