package apiV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvboda/quick-chat/model"
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
)

// SendValidateInfo 发送验证信息
func SendValidateInfo(c *gin.Context) {
	var relation model.RelationEntity
	var query struct {
		UserId   string `binding:"required"`
		FriendId string `binding:"required"`
		Memo     string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	if !utils.CheckAuthByUserId(c, query.UserId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NO_RIGHT, nil, nil))
		return
	}

	relation.UserId = query.UserId
	relation.FriendId = query.FriendId
	relation.Memo = query.Memo
	relation.RelationType = 1

	if code := relation.Insert(); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_RELATION_VALIDATE_SEND, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}

// AddFriend 添加好友
func AddFriend(c *gin.Context) {
	var relation model.RelationEntity
	var query struct {
		UserId   string `binding:"required"`
		FriendId string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	if !utils.CheckAuthByUserId(c, query.UserId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NO_RIGHT, nil, nil))
		return
	}

	relation.UserId = query.FriendId
	relation.FriendId = query.UserId
	relation.RelationType = 2
	if code := relation.AddFriend(); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_RELATION_ADD, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}

// RemoveFriend 删除好友
func RemoveFriend(c *gin.Context) {
	var relation model.RelationEntity
	var query struct {
		UserId   string `binding:"required"`
		FriendId string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	if !utils.CheckAuthByUserId(c, query.UserId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NO_RIGHT, nil, nil))
		return
	}

	relation.UserId = query.FriendId
	relation.FriendId = query.UserId
	relation.RelationType = 3
	if code := relation.RemoveFriend(); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_RELATION_DELETE, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}

// QueryValidateInfoList 查询验证信息列表
func QueryValidateInfoList(c *gin.Context) {
	var relation model.RelationEntity
	var query struct {
		FriendId     string
		RelationType int
	}
	query.FriendId = c.Param("uid")
	query.RelationType = 1

	if !utils.CheckAuthByUserId(c, query.FriendId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NO_RIGHT, nil, nil))
		return
	}

	if res, code := relation.SelectListBy(query); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_RELATION_VALIDATE_SELECT, nil, nil))
		return
	} else {
		c.JSON(http.StatusOK, status.GetResponse(code, nil, res))
	}
}

// QueryFriendList 查询好友列表
func QueryFriendList(c *gin.Context) {
	var relation model.RelationEntity

	if !utils.CheckAuthByUserId(c, c.Param("uid")) {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NO_RIGHT, nil, nil))
		return
	}

	if res, code := relation.SelectListBy("user_id = ? AND relation_type != ?", c.Param("uid"), 1); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_RELATION_SELECT, nil, nil))
		return
	} else {
		c.JSON(http.StatusOK, status.GetResponse(code, nil, res))
	}
}
