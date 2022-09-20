package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	Id        string         `gorm:"primarykey;type:varchar(32);not null" json:"id"`
	CreatedAt time.Time      `json:"createAt"`
	UpdatedAt time.Time      `json:"updateAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
