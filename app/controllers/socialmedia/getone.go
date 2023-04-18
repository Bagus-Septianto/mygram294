package socialmedia

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReadSocialMedia godoc
//	@Summary		Read a SocialMedia by ID
//	@Description	Reading a SocialMedia based on SocialMediaID
//	@Tags			SocialMedia
//	@Accept			json
//	@Produce		json
//	@Param			SocialMediaID	path		int	true	"SocialMedia ID"
//	@Success		200				{object}	models.SocialMedia
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/socialmedia/GetOne/{SocialMediaID} [get]
//	@Security		BearerAuth
func (controller SocialMediaController) GetOne(c *gin.Context) {
	SocialMediaID, _ := strconv.Atoi(c.Param("SocialMediaID"))
	socialMedia, err := controller.SocialMediaService.GetOne(uint(SocialMediaID))
	
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"socialMedia": socialMedia,
	})
}
