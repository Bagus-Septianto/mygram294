package socialmedia

import (
	"mygram294/app/models"
	"mygram294/pkg/common"
	"mygram294/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UpdateSocialMedia godoc
//
//	@Summary		Update a SocialMedia
//	@Description	Updating a SocialMedia by SocialMediaID
//	@Tags			SocialMedia
//	@Accept			json
//	@Produce		json
//	@Param			SocialMediaID	path		int					true	"SocialMedia ID"
//	@Param			SocialMedia		body		string				true	"SocialMedia Fields"	SchemaExample({"name":"bob","social_media_url":"https://bagusseptianto.com"})
//	@Success		200				{object}	models.SocialMedia	"SocialMedia"
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/socialmedia/UpdateSocialMedia/{SocialMediaID} [put]
//	@Security		BearerAuth
func (controller SocialMediaController) UpdateSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	socialMedia := models.SocialMedia{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	SocialMediaID, _ := strconv.Atoi(c.Param("SocialMediaID"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialMediaReturn, err := controller.SocialMediaService.UpdateSocialMedia(&socialMedia, userID, uint(SocialMediaID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"socialMedia": socialMediaReturn,
	})
}
