package photo

import (
	"mygram294/app/models"

	"github.com/go-playground/validator"
)

func (service PhotoService) CreatePhoto(p *models.Photo, UserID uint) (photo *models.Photo, err error) {
	validate = validator.New()
	err = validate.Struct(p)
	if err != nil {
		return nil, err
	}
	p.UserID = UserID
	photo, err = service.PhotoRepository.CreatePhoto(service.DB, p)
	if err != nil {
		return nil, err
	}
	return photo, nil
}
