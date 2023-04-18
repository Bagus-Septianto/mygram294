package comment

import (
	"errors"
	"mygram294/app/models"

	"github.com/go-playground/validator"
)

func (service CommentService) UpdateComment(p *models.Comment, UserID uint, CommentID uint) (comment *models.Comment, err error) {
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
	comment, err = service.CommentRepository.UpdateComment(service.DB, p, CommentID)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
