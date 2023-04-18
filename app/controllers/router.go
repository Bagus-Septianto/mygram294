package controllers

import (
	"mygram294/app/controllers/middlewares"

	"github.com/gin-gonic/gin"

	_ "mygram294/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

func StartApp(userController IUserController, photoController IPhotoController, socialMediaController ISocialMediaController, commentController ICommentController) *gin.Engine {

	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", userController.UserRegister)

		userRouter.POST("/login", userController.UserLogin)
	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())

		photoRouter.GET("/GetAll", photoController.GetAll)
		photoRouter.GET("/GetOne/:PhotoID", photoController.GetOne)
		photoRouter.POST("/CreatePhoto", photoController.CreatePhoto)
		photoRouter.PUT("/UpdatePhoto/:PhotoID", middlewares.PhotoAuthorization(), photoController.UpdatePhoto)
		photoRouter.DELETE("/DeletePhoto/:PhotoID", middlewares.PhotoAuthorization(), photoController.DeletePhoto)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())

		commentRouter.GET("/GetAll", commentController.GetAll)
		commentRouter.GET("/GetOne/:CommentID", commentController.GetOne)
		commentRouter.POST("/CreateComment", commentController.CreateComment)
		commentRouter.PUT("/UpdateComment/:CommentID", middlewares.CommentAuthorization(), commentController.UpdateComment)
		commentRouter.DELETE("/DeleteComment/:CommentID", middlewares.CommentAuthorization(), commentController.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedia")
	{
		socialMediaRouter.Use(middlewares.Authentication())

		socialMediaRouter.GET("/GetAll", socialMediaController.GetAll)
		socialMediaRouter.GET("/GetOne/:SocialMediaID", socialMediaController.GetOne)
		socialMediaRouter.POST("/CreateSocialMedia", socialMediaController.CreateSocialMedia)
		socialMediaRouter.PUT("/UpdateSocialMedia/:SocialMediaID", middlewares.SocialMediaAuthorization(), socialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/DeleteSocialMedia/:SocialMediaID", middlewares.SocialMediaAuthorization(), socialMediaController.DeleteSocialMedia)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
