package main

import (
	router "mygram294/app/controllers"
	commentRepo "mygram294/app/repository/comment"
	commentCont "mygram294/app/controllers/comment"
	commentServ "mygram294/app/services/comment"
	photoRepo "mygram294/app/repository/photo"
	photoCont "mygram294/app/controllers/photo"
	photoServ "mygram294/app/services/photo"
	socialMediaRepo "mygram294/app/repository/socialmedia"
	socialMediaCont "mygram294/app/controllers/socialmedia"
	socialMediaServ "mygram294/app/services/socialmedia"
	userRepo "mygram294/app/repository/user"
	userCont "mygram294/app/controllers/user"
	userServ "mygram294/app/services/user"
	"mygram294/pkg/database"
)

//	@title			My Gram
//	@version		1.0
//	@description	DTS FGA KOMINFO Hacktiv8 Golang-6

//	@contact.name	Bagus Septianto
//	@contact.url	http://bagusseptianto.com/
//	@contact.email	mygram@bagusseptianto.com

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apiKey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Bearer Token. Please add "Bearer " + your token
func main() {
	db, err := database.StartDB()
	if err != nil {
		panic(err)
	}
	userRepository := userRepo.NewUserRepository()
	userService := userServ.NewUserService(userRepository, db)
	userController := userCont.NewUserController(userService, db)

	photoRepository := photoRepo.NewPhotoRepository()
	photoService := photoServ.NewPhotoService(photoRepository, db)
	photoController := photoCont.NewPhotoController(photoService, db)

	socialMediaRepository := socialMediaRepo.NewSocialMediaRepository()
	socialMediaService := socialMediaServ.NewSocialMediaService(socialMediaRepository, db)
	socialMediaController := socialMediaCont.NewSocialMediaController(socialMediaService, db)

	commentRepository := commentRepo.NewCommentRepository()
	commentService := commentServ.NewCommentService(commentRepository, photoRepository, db)
	commentController := commentCont.NewCommentController(commentService, db)

	r := router.StartApp(userController, photoController, socialMediaController, commentController)
	r.Run("localhost:8080")
}