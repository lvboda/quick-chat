package apiV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lvboda/quick-chat/model"
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
)

// CreateCommunity 新建群聊
func CreateCommunity(c *gin.Context) {
	var community model.CommunityEntity
	var query struct {
		CommunityId string
	}

	if err := c.ShouldBindJSON(&community); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	if !utils.CheckAuthByUserId(c, community.OwnerId) {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_USER_NO_RIGHT, nil, nil))
		return
	}

	query.CommunityId = community.CommunityId
	if _, code := community.SelectBy(query); code == status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusBadRequest, status.GetResponse(status.ERROR_COMMUNITY_ID_USED, nil, nil))
		return
	}

	if code := community.Insert(); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_COMMUNITY_CREATE, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}

// EditCommunityById 修改群聊信息
func EditCommunityById(c *gin.Context) {
	var community model.CommunityEntity
	var query struct {
		CommunityId string
	}
	query.CommunityId = c.Param("cid")

	if err := c.ShouldBindJSON(&community); err != nil || query.CommunityId == "" {
		c.AbortWithStatusJSON(http.StatusOK, status.GetResponse(status.ERROR_REQUEST_PARAM, err, nil))
		return
	}

	v, code := community.SelectBy(query)
	if code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusOK, status.GetResponse(status.ERROR_COMMUNITY_NOT_EXIST, nil, nil))
		return
	}

	if !utils.CheckAuthByUserId(c, v.OwnerId) {
		c.AbortWithStatusJSON(http.StatusOK, status.GetResponse(status.ERROR_USER_NO_RIGHT, nil, nil))
		return
	}

	if code := community.Update(query.CommunityId); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusOK, status.GetResponse(status.ERROR_COMMUNITY_UPDATE, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}

// RemoveCommunityById 解散群
func RemoveCommunityById(c *gin.Context) {
	var community model.CommunityEntity
	var query struct {
		CommunityId string
	}
	query.CommunityId = c.Param("cid")

	if query.CommunityId == "" {
		c.AbortWithStatusJSON(http.StatusOK, status.GetResponse(status.ERROR_REQUEST_PARAM, nil, nil))
		return
	}

	v, code := community.SelectBy(query)
	if code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusOK, status.GetResponse(status.ERROR_COMMUNITY_NOT_EXIST, nil, nil))
		return
	}

	if !utils.CheckAuthByUserId(c, v.OwnerId) {
		c.AbortWithStatusJSON(http.StatusOK, status.GetResponse(status.ERROR_USER_NO_RIGHT, nil, nil))
		return
	}

	if code := community.Delete(query.CommunityId); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusOK, status.GetResponse(status.ERROR_COMMUNITY_DELETE, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, nil))
}

// QueryCommunityById 查询群聊
func QueryCommunityByCid(c *gin.Context) {
	var community model.CommunityEntity
	var query struct {
		CommunityId string
	}
	query.CommunityId = c.Param("cid")

	if res, code := community.SelectBy(query); code != status.SUCCESS {
		c.AbortWithStatusJSON(http.StatusOK, status.GetResponse(status.ERROR_COMMUNITY_NOT_EXIST, nil, nil))
		return
	} else {
		c.JSON(http.StatusOK, status.GetResponse(code, nil, res))
	}
}
