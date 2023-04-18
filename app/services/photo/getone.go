package photo

import (
	"mygram294/app/models"
)

func (service PhotoService) GetOne(PhotoID uint) (photo *models.Photo, err error) {
	photo, err = service.PhotoRepository.GetOne(service.DB, PhotoID)
	if err != nil {
		return nil, err
	}

	return photo, nil
}
