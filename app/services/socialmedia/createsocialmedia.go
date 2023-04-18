package socialmedia

import (
	"mygram294/app/models"

	"github.com/go-playground/validator"
)

func (service SocialMediaService) CreateSocialMedia(p *models.SocialMedia, UserID uint) (socialMedia *models.SocialMedia, err error) {
	validate = validator.New()
	err = validate.Struct(p)
	if err != nil {
		return nil, err
	}
	p.UserID = UserID
	socialMedia, err = service.SocialMediaRepository.CreateSocialMedia(service.DB, p)
	if err != nil {
		return nil, err
	}
	return socialMedia, nil
}
