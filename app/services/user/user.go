package user

import (
	repo "mygram294/app/repository"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type UserService struct {
	UserRepository repo.IUserRepository
	DB             *gorm.DB
}

func NewUserService(repository repo.IUserRepository, db *gorm.DB) *UserService {
	return &UserService{
		UserRepository: repository,
		DB:             db,
	}
}
