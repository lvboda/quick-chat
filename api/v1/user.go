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
	var user model.User
	var query struct {
		UserId string
	}
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	query.UserId = user.UserId
	_, code := user.SelectBy(query)
	if code == status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USERNAME_USED, nil, nil))
		return
	}

	code = user.Insert()
	if code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR, "注册失败", nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(code, nil, nil))
}

// Login 登录
func Login(c *gin.Context) {
	var user model.User
	var query struct {
		UserId   string `binding:"required,min=6,max=15"`
		Password string `binding:"required,min=6,max=15"`
	}
	err := c.ShouldBindJSON(&query)

	if err != nil {
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
	var user model.User
	var query struct {
		UserId string
	}
	query.UserId = c.Param("uid")

	res, code := user.SelectBy(query)

	if code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NOT_EXIST, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(code, nil, res))
}

// EditUserById 修改用户信息
func EditUserById(c *gin.Context) {
	var user model.User
	var query struct {
		Id string
	}
	query.Id = c.Param("id")
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	_, code := user.SelectBy(query)
	if code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NOT_EXIST, nil, nil))
		return
	}

	code = user.Update(query.Id)
	if code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_UPDATE, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(code, nil, nil))
}

// RemoveUserById 注销用户
func RemoveUserById(c *gin.Context) {
	var user model.User
	var query struct {
		Id string
	}
	query.Id = c.Param("id")

	_, code := user.SelectBy(query)
	if code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NOT_EXIST, nil, nil))
		return
	}

	code = user.Delete(query.Id)
	if code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_DELETE, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(code, nil, nil))
}
