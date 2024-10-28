package model

import (
	"gorm.io/gorm"
	"time"
)

// User 用户
type User struct {
	ID        string         `gorm:"column:id;type:varchar(24);NOT NULL" json:"id"`
	Avatar    string         `gorm:"column:avatar;type:varchar(191);NOT NULL;default:''" json:"avatar"`
	Nickname  string         `gorm:"column:nickname;type:varchar(24);NOT NULL;default:''" json:"nickname"`
	Phone     string         `gorm:"column:phone;type:varchar(20);NOT NULL;default:''" json:"phone"`
	Password  string         `gorm:"column:password;type:varchar(191);NOT NULL;default:''" json:"password"`
	Status    int64          `gorm:"column:status;type:tinyint(1);NOT NULL;default:0" json:"status"`
	Sex       int64          `gorm:"column:sex;type:tinyint(1);NOT NULL;default:0" json:"sex"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;not null;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;not null;comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;comment:删除时间" json:"deleted_at"`
}
