package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsimao/bookstore_users-api/domain/users"
	"github.com/hsimao/bookstore_users-api/services"
	"github.com/hsimao/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User

	// 參數格式錯誤
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
