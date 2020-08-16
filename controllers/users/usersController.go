package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hsimao/bookstore_users-api/domain/users"
	"github.com/hsimao/bookstore_users-api/services"
	"github.com/hsimao/bookstore_users-api/utils/errors"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
	var user users.User

	// 參數格式錯誤
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)

	// 判斷儲存用戶資料到 database 是否失敗
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("userId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("userId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User

	// 參數格式錯誤
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	// 判斷是否是使用 patch 方法
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("userId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
