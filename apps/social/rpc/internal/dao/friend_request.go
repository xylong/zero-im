package dao

import (
	"errors"
	"gorm.io/gorm"
	"zero-im/models"
)

type FriendRequestDao interface {
	Create(request *models.FriendRequest) error
	Update(request *models.FriendRequest) error
	FindOne(id int64) (*models.FriendRequest, error)
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return reqs, err
}

func (d *friendRequestDao) Create(request *models.FriendRequest) error {
	return d.db.Create(request).Error
}

func (d *friendRequestDao) Update(request *models.FriendRequest) error {
	return d.db.Model(request).Updates(request).Error
}

func (d *friendRequestDao) FindOne(id int64) (*models.FriendRequest, error) {
	var (
		err error
		req models.FriendRequest
	)

	err = d.db.Where("id = ?", id).First(&req).Error
	if err != nil {
		return nil, err
	}

	return &req, nil
}
