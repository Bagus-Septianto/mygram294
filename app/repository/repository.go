package repository

import (
	"mygram294/app/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Register(db *gorm.DB, user *models.User) (err error)
	Login(db *gorm.DB, username string) (user *models.User, err error)
}

type IPhotoRepository interface {
	GetAll(db *gorm.DB) (photo []*models.Photo, err error)
	GetOne(db *gorm.DB, PhotoID uint) (photo *models.Photo, err error)
	CreatePhoto(db *gorm.DB, p *models.Photo) (photo *models.Photo, err error)
	UpdatePhoto(db *gorm.DB, p *models.Photo, PhotoID uint) (photo *models.Photo, err error)
	DeletePhoto(db *gorm.DB, PhotoID uint) (err error)
}

type ISocialMediaRepository interface {
	GetAll(db *gorm.DB) (socialMedia []*models.SocialMedia, err error)
	GetOne(db *gorm.DB, SocialMediaID uint) (socialMedia *models.SocialMedia, err error)
	CreateSocialMedia(db *gorm.DB, p *models.SocialMedia) (socialMedia *models.SocialMedia, err error)
	UpdateSocialMedia(db *gorm.DB, p *models.SocialMedia, SocialMediaID uint) (socialMedia *models.SocialMedia, err error)
	DeleteSocialMedia(db *gorm.DB, SocialMediaID uint) (err error)
}

type ICommentRepository interface {
	GetAll(db *gorm.DB) (comment []*models.Comment, err error)
	GetOne(db *gorm.DB, CommentID uint) (comment *models.Comment, err error)
	CreateComment(db *gorm.DB, p *models.Comment) (comment *models.Comment, err error)
	UpdateComment(db *gorm.DB, p *models.Comment, CommentID uint) (comment *models.Comment, err error)
	DeleteComment(db *gorm.DB, PhotoID uint) (err error)
}