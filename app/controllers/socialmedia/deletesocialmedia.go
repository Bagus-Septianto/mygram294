package socialmedia

import (
	"mygram294/app/models"
	"mygram294/pkg/common"
	"mygram294/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteSocialMedia godoc
//	@Summary		Delete a SocialMedia
//	@Description	Deleting a SocialMedia by SocialMediaID
//	@Tags			SocialMedia
//	@Produce		json
//	@Param			SocialMediaID	path	int	true	"SocialMedia ID"
//	@Success		204
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/socialmedia/DeleteSocialMedia/{SocialMediaID} [delete]
//	@Security		BearerAuth
func (controller SocialMediaController) DeleteSocialMedia(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	socialMedia := models.SocialMedia{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	SocialMediaID, _ := strconv.Atoi(c.Param("SocialMediaID"))
	err := controller.SocialMediaService.DeleteSocialMedia(uint(SocialMediaID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
