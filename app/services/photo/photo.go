package photo

import (
	repo "mygram294/app/repository"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type PhotoService struct {
	PhotoRepository repo.IPhotoRepository
	DB              *gorm.DB
}

func NewPhotoService(repository repo.IPhotoRepository, db *gorm.DB) *PhotoService {
	return &PhotoService{
		PhotoRepository: repository,
		DB:              db,
	}
}
