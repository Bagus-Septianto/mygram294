package photo

import (
	"mygram294/app/services"

	"gorm.io/gorm"
)

type PhotoController struct {
	PhotoService services.IPhotoService
	DB           *gorm.DB
}

func NewPhotoController(photoService services.IPhotoService, db *gorm.DB) *PhotoController {
	return &PhotoController{
		PhotoService: photoService,
		DB:           db,
	}
}
