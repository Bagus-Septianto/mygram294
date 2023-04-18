package socialmedia

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r SocialMediaRepository) CreateSocialMedia(db *gorm.DB, p *models.SocialMedia) (socialMedia *models.SocialMedia, err error) {
	err = db.Debug().Create(&p).Error
	return p, err
}
