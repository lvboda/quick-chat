package model

import (
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
)

type CommunityEntity struct {
	BaseEntity
	CommunityId string `gorm:"type:varchar(32);not null;INDEX" json:"communityId" binding:"required,min=6,max=12" label:"群id"`
	Name        string `gorm:"type:varchar(32);not null" json:"Name" binding:"required,min=1,max=12" label:"群名称"`
	OwnerId     string `gorm:"type:varchar(32);not null" json:"ownerId" binding:"required" label:"群主id"`
	Face        string `gorm:"type:varchar(100);not null" json:"face" label:"群头像"`
	Memo        string `gorm:"type:varchar(100);" json:"memo" label:"群描述"`
	Extend      string `gorm:"-" json:"extend"`
}

func (CommunityEntity) TableName() string {
	return "community"
}

// Insert 增
func (community CommunityEntity) Insert() int {
	community.Id = utils.UUID()

	if err := utils.Db.Create(&community).Error; err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// Delete 删
func (community CommunityEntity) Delete(cid string) int {
	if err := utils.Db.Where("community_id = ? ", cid).Delete(&community).Error; err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// Update 改
func (community CommunityEntity) Update(cid string) int {
	updateFields := map[string]any{
		"name":     community.Name,
		"owner_id": community.OwnerId,
		"face":     community.Face,
		"memo":     community.Memo,
	}

	if err := utils.Db.Model(&community).Where("community_id = ?", cid).Updates(&updateFields).Error; err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// SelectBy 查单条
func (community CommunityEntity) SelectBy(query any) (CommunityEntity, int) {
	var res CommunityEntity

	if err := utils.Db.Where(query).First(&res).Error; err != nil {
		return res, status.ERROR
	}
	return res, status.SUCCESS
}
