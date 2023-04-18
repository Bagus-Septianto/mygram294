package user

import (
	"mygram294/app/models"
	"mygram294/pkg/common"
	"mygram294/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegister godoc
//	@Summary		Register User
//	@Description	Creating new user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			User	body		string	true	"User"	SchemaExample({"username":"username","password":"12345678","email":"email@bagusseptianto.com","age":21})
//	@Success		201		{object}	string
//	@Failure		400
//	@Router			/user/register [post]
func (controller UserController) UserRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == common.AppJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	userReturn, err := controller.UserService.Register(&User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": userReturn.ID,
		"username": userReturn.Username,
	})
}