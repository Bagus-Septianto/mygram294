package socialmedia

import (
	"errors"
	"mygram294/app/models"
)

func (service SocialMediaService) GetAll() (socialMedia []*models.SocialMedia, err error) {
	socialMedia, err = service.SocialMediaRepository.GetAll(service.DB)
	if err != nil {
		return nil, err
	}
	if len(socialMedia) != 0 {
		return socialMedia, nil
	}

	return nil, errors.New("record not found")
}
