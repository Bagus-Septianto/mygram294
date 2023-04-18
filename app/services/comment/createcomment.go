package comment

import (
	"errors"
	"mygram294/app/models"

	"github.com/go-playground/validator"
)

func (service CommentService) CreateComment(p *models.Comment, UserID uint) (comment *models.Comment, err error) {
	validate = validator.New()
	err = validate.Struct(p)
	if err != nil {
		return nil, err
	}
	// check if photo is available
	_, err = service.PhotoRepository.GetOne(service.DB, p.PhotoID)
	if err != nil {
		return nil, errors.New("photo_id is invalid")
	}
	p.UserID = UserID
	comment, err = service.CommentRepository.CreateComment(service.DB, p)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
