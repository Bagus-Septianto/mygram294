package socialmedia

import (
	repo "mygram294/app/repository"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type SocialMediaService struct {
	SocialMediaRepository repo.ISocialMediaRepository
	DB              *gorm.DB
}

func NewSocialMediaService(repository repo.ISocialMediaRepository, db *gorm.DB) *SocialMediaService {
	return &SocialMediaService{
		SocialMediaRepository: repository,
		DB:              db,
	}
}
