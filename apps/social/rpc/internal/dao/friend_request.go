package dao

import (
	"gorm.io/gorm"
	"zero-im/models"
)

type FriendRequestDao interface {
	Create(request *models.FriendRequest) error
	FindByReqUidAndUserId(rid, uid string) ([]*models.FriendRequest, error)
}

type friendRequestDao struct {
	db *gorm.DB
}

func NewFriendRequestDao(db *gorm.DB) FriendRequestDao {
	return &friendRequestDao{db: db}
}

func (d *friendRequestDao) FindByReqUidAndUserId(rid, uid string) ([]*models.FriendRequest, error) {
	var (
		err  error
		reqs []*models.FriendRequest
	)

	err = d.db.Where("req_uid = ?", rid).Where("user_id = ?", uid).Find(&reqs).Error
	return reqs, err
}

func (d *friendRequestDao) Create(request *models.FriendRequest) error {
	return d.db.Create(request).Error
}
