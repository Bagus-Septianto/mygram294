package socialmedia

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r SocialMediaRepository) GetOne(db *gorm.DB, SocialMediaID uint) (socialMedia *models.SocialMedia, err error) {
	err = db.Debug().First(&socialMedia, "id = ?", SocialMediaID).Error

	return socialMedia, err
}
