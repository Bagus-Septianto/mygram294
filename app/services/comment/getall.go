package comment

import (
	"errors"
	"mygram294/app/models"
)

func (service CommentService) GetAll() (comment []*models.Comment, err error) {
	comment, err = service.CommentRepository.GetAll(service.DB)
	if err != nil {
		return nil, err
	}
	if len(comment) != 0 {
		return comment, nil
	}

	return nil, errors.New("record not found")
}
