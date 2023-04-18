package socialmedia

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r SocialMediaRepository) GetAll(db *gorm.DB) (socialMedia []*models.SocialMedia, err error) {
	err = db.Debug().Find(&socialMedia).Error
	return socialMedia, err
}
