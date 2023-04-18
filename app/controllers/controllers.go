package controllers

import "github.com/gin-gonic/gin"

type IUserController interface {
	UserRegister(c *gin.Context)
	UserLogin(c *gin.Context)
}

type IPhotoController interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	CreatePhoto(c *gin.Context)
	UpdatePhoto(c *gin.Context)
	DeletePhoto(c *gin.Context)
}

type ISocialMediaController interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	CreateSocialMedia(c *gin.Context)
	UpdateSocialMedia(c *gin.Context)
	DeleteSocialMedia(c *gin.Context)
}

type ICommentController interface {
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	CreateComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
}