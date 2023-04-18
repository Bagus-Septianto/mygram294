package photo

import (
	"errors"
	"mygram294/app/models"
)

func (service PhotoService) GetAll() (photo []*models.Photo, err error) {
	photo, err = service.PhotoRepository.GetAll(service.DB)
	if err != nil {
		return nil, err
	}
	if len(photo) != 0 {
		return photo, nil
	}

	return nil, errors.New("record not found")
}
