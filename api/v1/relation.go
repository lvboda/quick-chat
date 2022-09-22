package apiV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvboda/quick-chat/model"
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
)

// SendValidate 发送验证信息
func SendValidate(c *gin.Context) {
	var relation model.RelationEntity
	var query struct {
		UserId   string `binding:"required"`
		FriendId string `binding:"required"`
		Memo     string `binding:"required"`
		RoleType int    `binding:"required"`
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
	relation.RoleType = query.RoleType
	relation.RelationType = 1

	if code := relation.Insert(); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_RELATION_VALIDATE_SEND, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}

// AddRelation 添加关系
func AddRelation(c *gin.Context) {
	var relation model.RelationEntity
	var query struct {
		UserId   string `binding:"required"`
		FriendId string `binding:"required"`
		RoleType int    `binding:"required"`
	}

	if err := c.ShouldBindJSON(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	if query.RoleType == 1 && !utils.CheckAuthByUserId(c, query.UserId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NO_RIGHT, nil, nil))
		return
	}

	relation.UserId = query.FriendId
	relation.FriendId = query.UserId
	relation.RoleType = query.RoleType
	relation.RelationType = 2
	if code := relation.AddFriend(); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_RELATION_ADD, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}

// RemoveRelation 删除关系
func RemoveRelation(c *gin.Context) {
	var relation model.RelationEntity
	var query struct {
		UserId   string `binding:"required"`
		FriendId string `binding:"required"`
		RoleType int    `binding:"required"`
	}

	if err := c.ShouldBindJSON(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	if query.RoleType == 1 && !utils.CheckAuthByUserId(c, query.UserId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NO_RIGHT, nil, nil))
		return
	}

	relation.UserId = query.FriendId
	relation.FriendId = query.UserId
	relation.RoleType = query.RoleType
	relation.RelationType = 3
	if code := relation.RemoveFriend(); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_RELATION_DELETE, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}

// QueryList 查询关系信息列表
func QueryRelationList(c *gin.Context) {
	var relation model.RelationEntity
	var query struct {
		FriendId     string
		RelationType int
		RoleType     int
	}

	if err := c.ShouldBindJSON(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	if query.RoleType == 1 && !utils.CheckAuthByUserId(c, query.FriendId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NO_RIGHT, nil, nil))
		return
	}

	if res, code := relation.SelectListBy(query, query.RoleType); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_RELATION_VALIDATE_SELECT, nil, nil))
		return
	} else {
		c.JSON(http.StatusOK, status.GetResponse(code, nil, res))
	}
}
