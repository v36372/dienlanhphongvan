package entity

import (
	"dienlanhphongvan/models"
	"dienlanhphongvan/repo"
	"dienlanhphongvan/utilities/uer"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userEntity struct{}

type User interface {
	Login(username, password string) (*models.User, error)
}

func NewUser() User {
	return &userEntity{}
}

func (userEntity) Login(username, password string) (*models.User, error) {
	user, err := repo.User.GetByUsername(username)
	if err != nil {
		return nil, uer.InternalError(err)
	}

	if user == nil {
		return nil, uer.NotFoundError(errors.New("user not found"))
	}

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, uer.NotAuthorizedError(errors.New("unauthorized"))
	}

	return user, nil
}
