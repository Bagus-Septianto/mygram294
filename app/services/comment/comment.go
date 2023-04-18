package comment

import (
	repo "mygram294/app/repository"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type CommentService struct {
	CommentRepository repo.ICommentRepository
	PhotoRepository repo.IPhotoRepository
	DB              *gorm.DB
}

func NewCommentService(repository repo.ICommentRepository, prepo repo.IPhotoRepository, db *gorm.DB) *CommentService {
	return &CommentService{
		CommentRepository: repository,
		PhotoRepository: prepo,
		DB:              db,
	}
}
