package photo

import (
	"mygram294/app/models"

	"github.com/go-playground/validator"
)

func (service PhotoService) UpdatePhoto(p *models.Photo, UserID uint, PhotoID uint) (photo *models.Photo, err error) {
	validate = validator.New()
	err = validate.Struct(p)
	if err != nil {
		return nil, err
	}
	p.UserID = UserID
	photo, err = service.PhotoRepository.UpdatePhoto(service.DB, p, PhotoID)
	if err != nil {
		return nil, err
	}
	return photo, nil
}
