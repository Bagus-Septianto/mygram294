package photo

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r PhotoRepository) GetOne(db *gorm.DB, PhotoID uint) (photo *models.Photo, err error) {
	err = db.Debug().First(&photo, "id = ?", PhotoID).Error

	return photo, err
}
