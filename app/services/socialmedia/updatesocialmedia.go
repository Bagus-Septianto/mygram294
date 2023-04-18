package socialmedia

import (
	"mygram294/app/models"

	"github.com/go-playground/validator"
)

func (service SocialMediaService) UpdateSocialMedia(p *models.SocialMedia, UserID uint, SocialMediaID uint) (socialMedia *models.SocialMedia, err error) {
	validate = validator.New()
	err = validate.Struct(p)
	if err != nil {
		return nil, err
	}
	p.UserID = UserID
	socialMedia, err = service.SocialMediaRepository.UpdateSocialMedia(service.DB, p, SocialMediaID)
	if err != nil {
		return nil, err
	}
	return socialMedia, nil
}
