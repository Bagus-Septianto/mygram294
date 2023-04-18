package photo

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r PhotoRepository) GetAll(db *gorm.DB) (photo []*models.Photo, err error) {
	err = db.Debug().Find(&photo).Error
	return photo, err
}
