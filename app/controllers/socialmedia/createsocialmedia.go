package socialmedia

import (
	"mygram294/app/models"
	"mygram294/pkg/common"
	"mygram294/pkg/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateSocialMedia godoc
//	@Summary		Create a SocialMedia
//	@Description	Creating a SocialMedia
//	@Tags			SocialMedia
//	@Accept			json
//	@Produce		json
//	@Param			SocialMedia	body		string				true	"SocialMedia Fields"	SchemaExample({"name":"bob","social_media_url":"https://bagusseptianto.com"})
//	@Success		201			{object}	models.SocialMedia	"SocialMedia"
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/socialmedia/CreateSocialMedia [post]
//	@Security		BearerAuth
func (controller SocialMediaController) CreateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	socialMedia := models.SocialMedia{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialMediaReturn, err := controller.SocialMediaService.CreateSocialMedia(&socialMedia, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"socialMedia": socialMediaReturn,
	})
}
