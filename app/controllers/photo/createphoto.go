package photo

import (
	"mygram294/app/models"
	"mygram294/pkg/common"
	"mygram294/pkg/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


// CreatePhoto godoc
//	@Summary		Create a Photo
//	@Description	Creating a Photo
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			Photo	body		string			true	"Photo Fields"	SchemaExample({"title":"Lorem Ipsum","caption":"Lorem ipsum dolor sit amet","photo_url":"https://css.category.id/cat.jpg"})
//	@Success		201		{object}	models.Photo	"Photo"
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/photo/CreatePhoto [post]
//	@Security		BearerAuth
func (controller PhotoController) CreatePhoto(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	photo := models.Photo{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photoReturn, err := controller.PhotoService.CreatePhoto(&photo, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"photo": photoReturn,
	})
}
