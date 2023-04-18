package user

import (
	"mygram294/app/models"
	"mygram294/pkg/helpers"

	"github.com/go-playground/validator"
)

func (service UserService) Register(u *models.User) (user *models.User, err error) {
	validate = validator.New()
	err = validate.Struct(u)
	if err != nil {
		return nil, err
	}
	u.Password = helpers.HashPass(u.Password)
	err = service.UserRepository.Register(service.DB, u)
	if err != nil {
		return nil, err
	}
	user = u
	return user, nil
}
