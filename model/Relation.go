package model

import (
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
	"gorm.io/gorm"
)

type RelationEntity struct {
	BaseEntity
	UserId        string          `gorm:"type:varchar(32);not null" json:"userId" binding:"required,min=6,max=15" label:"用户id"`
	FriendId      string          `gorm:"type:varchar(32);not null" json:"friendId" binding:"required,min=6,max=15" label:"好友id"`
	RelationType  int             `gorm:"type:int;DEFAULT:1" json:"relationType" binding:"required" label:"关系类型"`
	RoleType      int             `gorm:"type:int;DEFAULT:1" json:"roleType" binding:"required" label:"角色类型"`
	Memo          string          `gorm:"type:varchar(120);DEFAULT:NULL" json:"memo" label:"描述"`
	Extend        string          `gorm:"-" json:"extend"`
	FriendInfo    UserEntity      `gorm:"foreignKey:UserId;references:UserId;" json:"friendInfo"`
	ProposerInfo  UserEntity      `gorm:"foreignKey:UserId;references:FriendId;" json:"proposerInfo"`
	CommunityInfo CommunityEntity `gorm:"foreignKey:CommunityId;references:FriendId;" json:"communityInfo"`
}

func (RelationEntity) TableName() string {
	return "user_relation"
}

// Insert 增
func (relation RelationEntity) Insert() int {
	relation.Id = utils.UUID()

	if err := utils.Db.Create(&relation).Error; err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// Delete 删
func (relation RelationEntity) Delete(id string) int {
	if err := utils.Db.Where("id = ? ", id).Delete(&relation).Error; err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// Update 改
func (relation RelationEntity) Update(query any, args ...any) int {
	updateFields := map[string]any{
		"user_id":       relation.UserId,
		"friend_id":     relation.FriendId,
		"relation_type": relation.RelationType,
	}

	if err := utils.Db.Model(&relation).Where(query, args...).Updates(&updateFields).Error; err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// SelectBy 查单条
func (relation RelationEntity) SelectBy(query any, role int) (RelationEntity, int) {
	var res RelationEntity

	db := utils.Db.Where(query)
	if err := utils.If(role == 1, db.Preload("FriendInfo").Find(&res), db.Preload("ProposerInfo").Preload("CommunityInfo").First(&res)).Error; err != nil {
		return res, status.ERROR
	}
	return res, status.SUCCESS
}

// SelectListBy 查多条
func (relation RelationEntity) SelectListBy(query any, role int) ([]RelationEntity, int) {
	var res []RelationEntity

	db := utils.Db.Where(query)
	if err := utils.If(role == 1, db.Preload("FriendInfo").Find(&res), db.Preload("ProposerInfo").Preload("CommunityInfo").Find(&res)).Error; err != nil {
		return res, status.ERROR
	}
	return res, status.SUCCESS
}

// AddFriend 添加好友
func (relation RelationEntity) AddFriend() int {
	err := utils.Db.Transaction(func(tx *gorm.DB) error {
		updateFields := map[string]any{
			"user_id":       relation.UserId,
			"friend_id":     relation.FriendId,
			"relation_type": relation.RelationType,
		}

		if err := tx.Model(&relation).Where("user_id = ? AND friend_id = ?", relation.UserId, relation.FriendId).Updates(&updateFields).Error; err != nil {
			return err
		}

		relation.Id = utils.UUID()
		temp := relation.UserId
		relation.UserId = relation.FriendId
		relation.FriendId = temp
		if err := tx.Create(&relation).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// RemoveFriend 删除好友
func (relation RelationEntity) RemoveFriend() int {
	err := utils.Db.Transaction(func(tx *gorm.DB) error {
		updateFields := map[string]any{
			"user_id":       relation.UserId,
			"friend_id":     relation.FriendId,
			"relation_type": relation.RelationType,
		}

		if err := tx.Model(&relation).Where("user_id = ? AND friend_id = ?", relation.UserId, relation.FriendId).Updates(&updateFields).Error; err != nil {
			return err
		}

		if err := tx.Where("user_id = ? AND friend_id = ?", relation.FriendId, relation.UserId).Delete(&relation).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}
