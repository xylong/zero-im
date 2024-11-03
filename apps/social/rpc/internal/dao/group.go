package dao

import "gorm.io/gorm"

type GroupDao interface {
}

type groupDao struct {
	db *gorm.DB
}

func NewGroupDao(db *gorm.DB) GroupDao {
	return &groupDao{db: db}
}
