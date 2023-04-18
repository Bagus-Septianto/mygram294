package services

import "mygram294/app/models"

type IUserService interface {
	Register(u *models.User) (user *models.User, err error)
	Login(username, password string) (token string, err error)
}

type IPhotoService interface {
	GetAll() (photo []*models.Photo, err error)
	GetOne(PhotoID uint) (photo *models.Photo, err error)
	CreatePhoto(p *models.Photo, UserID uint) (photo *models.Photo, err error)
	UpdatePhoto(p *models.Photo, UserID uint, PhotoID uint) (photo *models.Photo, err error)
	DeletePhoto(UserID uint) (err error)
}

type ISocialMediaService interface {
	GetAll() (socialMedia []*models.SocialMedia, err error)
	GetOne(SocialMediaID uint) (socialMedia *models.SocialMedia, err error)
	CreateSocialMedia(p *models.SocialMedia, UserID uint) (socialMedia *models.SocialMedia, err error)
	UpdateSocialMedia(p *models.SocialMedia, UserID uint, SocialMediaID uint) (socialMedia *models.SocialMedia, err error)
	DeleteSocialMedia(UserID uint) (err error)
}

type ICommentService interface {
	GetAll() (comment []*models.Comment, err error)
	GetOne(CommentID uint) (comment *models.Comment, err error)
	CreateComment(p *models.Comment, UserID uint) (comment *models.Comment, err error)
	UpdateComment(p *models.Comment, UserID uint, CommentID uint) (comment *models.Comment, err error)
	DeleteComment(UserID uint) (err error)
}