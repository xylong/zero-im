package dao

import (
	"errors"
	"gorm.io/gorm"
	"zero-im/models"
)

type FriendDao interface {
	FindByUidAndFid(uid, fid string) (*models.Friend, error)
	Insert(...*models.Friend) error
	ListByUserid(userid string) ([]*models.Friend, error)
}

type friendDao struct {
	db *gorm.DB
}

func NewFriendDao(db *gorm.DB) FriendDao {
	return &friendDao{db: db}
}

func (d *friendDao) FindByUidAndFid(uid, fid string) (*models.Friend, error) {
	var (
		err    error
		friend models.Friend
	)

	err = d.db.Where("user_id = ?", uid).Where("friend_uid = ?", fid).First(&friend).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &friend, err
}

func (d *friendDao) Insert(friends ...*models.Friend) error {
	if len(friends) == 0 {
		return nil
	}

	return d.db.CreateInBatches(friends, 100).Error
}

func (d *friendDao) ListByUserid(userid string) ([]*models.Friend, error) {
	var (
		err     error
		friends []*models.Friend
	)

	err = d.db.Model(&models.Friend{}).Where("user_id = ?", userid).Find(&friends).Error
	if err != nil {
		return nil, err
	}

	return friends, nil
}
