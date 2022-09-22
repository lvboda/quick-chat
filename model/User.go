package model

import (
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
)

type UserEntity struct {
	BaseEntity
	NickName  string `gorm:"type:varchar(32);not null;INDEX" json:"nickName" binding:"required,min=1,max=12" label:"昵称"`
	UserId    string `gorm:"type:varchar(32);not null " json:"userId" binding:"required,min=6,max=15" label:"用户id"`
	Password  string `gorm:"type:varchar(100);not null " json:"password" binding:"required,min=6,max=15" label:"密码"`
	UserRole  int    `gorm:"type:int;DEFAULT:1" json:"userRole" binding:"required" label:"角色类型"`
	Gender    int    `gorm:"type:int;DEFAULT:1" json:"gender" binding:"required" label:"性别"`
	Signature string `gorm:"type:varchar(255) " json:"signature" label:"个人签名"`
	Mobile    string `gorm:"type:varchar(16) " json:"mobile" label:"电话"`
	Face      string `gorm:"type:varchar(255) " json:"face" label:"头像"`
}

func (UserEntity) TableName() string {
	return "user_base"
}

// Insert 增
func (user UserEntity) Insert() int {
	user.Id = utils.UUID()

	if err := utils.Db.Create(&user).Error; err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// Delete 删
func (user UserEntity) Delete(uid string) int {
	if err := utils.Db.Where("user_id = ? ", uid).Delete(&user).Error; err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// Update 改
func (user UserEntity) Update(uid string) int {
	updateFields := map[string]any{
		"nick_name": user.NickName,
		"password":  user.Password,
		"gender":    user.Gender,
		"signature": user.Signature,
		"mobile":    user.Mobile,
		"face":      user.Face,
	}

	if err := utils.Db.Model(&user).Where("user_id = ?", uid).Updates(&updateFields).Error; err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// SelectBy 查单条
func (user UserEntity) SelectBy(query any) (UserEntity, int) {
	var res UserEntity

	if err := utils.Db.Where(query).First(&res).Error; err != nil {
		return res, status.ERROR
	}
	return res, status.SUCCESS
}
