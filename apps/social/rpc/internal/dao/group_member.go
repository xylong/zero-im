package dao

import "gorm.io/gorm"

type GroupMemberDao interface {
}

type groupMemberDao struct {
	db *gorm.DB
}

func NewGroupMemberDao(db *gorm.DB) GroupMemberDao {
	return &groupMemberDao{db: db}
}
