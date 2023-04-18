package socialmedia

import (
	"mygram294/app/services"

	"gorm.io/gorm"
)

type SocialMediaController struct {
	SocialMediaService services.ISocialMediaService
	DB           *gorm.DB
}

func NewSocialMediaController(socialMediaService services.ISocialMediaService, db *gorm.DB) *SocialMediaController {
	return &SocialMediaController{
		SocialMediaService: socialMediaService,
		DB:           db,
	}
}
