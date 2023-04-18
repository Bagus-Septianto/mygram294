package socialmedia

import (
	"mygram294/app/models"
)

func (service SocialMediaService) GetOne(SocialMediaID uint) (socialMedia *models.SocialMedia, err error) {
	socialMedia, err = service.SocialMediaRepository.GetOne(service.DB, SocialMediaID)
	if err != nil {
		return nil, err
	}

	return socialMedia, nil
}
