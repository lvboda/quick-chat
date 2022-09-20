package apiV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvboda/quick-chat/model"
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
)

// Register 注册
func Register(c *gin.Context) {
	var user model.UserEntity
	var query struct {
		UserId string
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	query.UserId = user.UserId
	if _, code := user.SelectBy(query); code == status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USERNAME_USED, nil, nil))
		return
	}

	if code := user.Insert(); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_USER_REGISTER, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}

// Login 登录
func Login(c *gin.Context) {
	var user model.UserEntity
	var query struct {
		UserId   string `binding:"required,min=6,max=15"`
		Password string `binding:"required,min=6,max=15"`
	}

	if err := c.ShouldBindJSON(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	data, code := user.SelectBy(query)
	if code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_PASSWORD_WRONG, nil, nil))
		return
	}

	token, err := utils.CreateToken(data.UserId, data.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_TOKEN_CREATE, err, nil))
		return
	}

	data.Password = ""
	resData, _ := utils.MergeJson(data, map[string]any{"token": token})

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, resData))
}

// QueryUserByUid 查询用户信息
func QueryUserByUid(c *gin.Context) {
	var user model.UserEntity
	var query struct {
		UserId string
	}
	query.UserId = c.Param("uid")

	if res, code := user.SelectBy(query); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NOT_EXIST, nil, nil))
		return
	} else {
		c.JSON(http.StatusOK, status.GetResponse(code, nil, res))
	}
}

// EditUserById 修改用户信息
func EditUserById(c *gin.Context) {
	var user model.UserEntity
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil || id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	if code := user.Update(id); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_USER_UPDATE, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}

// RemoveUserById 注销用户
func RemoveUserById(c *gin.Context) {
	var user model.UserEntity
	id := c.Param("id")

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, nil, nil))
		return
	}

	if code := user.Delete(id); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_USER_DELETE, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}
