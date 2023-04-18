package photo

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

func (r PhotoRepository) CreatePhoto(db *gorm.DB, p *models.Photo) (photo *models.Photo, err error) {
	err = db.Debug().Create(&p).Error
	return p, err
}
