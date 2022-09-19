package model

import (
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
)

type User struct {
	Base
	NickName  string `gorm:"type:varchar(32);not null " binding:"required,min=2,max=12" label:"昵称"`
	UserId    string `gorm:"type:varchar(32);not null " binding:"required,min=6,max=15" label:"用户id"`
	Password  string `gorm:"type:varchar(100);not null " binding:"required,min=6,max=15" label:"密码"`
	UserRole  int    `gorm:"type:int;DEFAULT:0" binding:"required,gte=0" label:"角色类型"`
	Gender    int    `gorm:"type:int;DEFAULT:0" binding:"required,gte=0" label:"性别"`
	Signature string `gorm:"type:varchar(255) " label:"个人签名"`
	Mobile    string `gorm:"type:varchar(16) " label:"电话"`
	Face      string `gorm:"type:varchar(255) " label:"头像"`
}

func (User) TableName() string {
	return "user_base"
}

// 增
func (user User) Insert() int {
	user.ID = utils.UUID()

	err := utils.Db.Create(&user).Error

	if err != nil {
		return status.ERROR
	}

	return status.SUCCESS
}

// 删
func (user User) Delete(id string) int {
	err := utils.Db.Where("id = ? ", id).Delete(&user).Error

	if err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// 改
func (user User) Update(id string) int {
	query := map[string]any{
		"nick_name": user.NickName,
		"password":  user.Password,
		"gender":    user.Gender,
		"signature": user.Signature,
		"mobile":    user.Mobile,
		"face":      user.Face,
	}
	err := utils.Db.Debug().Model(&user).Where("id = ?", id).Updates(&query).Error

	if err != nil {
		return status.ERROR
	}
	return status.SUCCESS
}

// 查单条
func (user User) SelectBy(query any) (User, int) {
	var res User

	err := utils.Db.Debug().Where(query).First(&res).Error

	if err != nil {
		return res, status.ERROR
	}
	return res, status.SUCCESS
}
