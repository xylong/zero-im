package svc

import (
	"gorm.io/gorm"
	"zero-im/apps/social/rpc/internal/config"
	"zero-im/apps/social/rpc/internal/dao"
	"zero-im/pkg/core/store/gormc"
)

type ServiceContext struct {
	Config config.Config

	DB *gorm.DB

	FriendDao        dao.FriendDao
	FriendRequestDao dao.FriendRequestDao
	GroupDao         dao.GroupDao
	GroupMemberDao   dao.GroupMemberDao
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := gormc.NewEngine(c.Mysql)

	return &ServiceContext{
		Config: c,

		DB: db,

		FriendDao:        dao.NewFriendDao(db),
		FriendRequestDao: dao.NewFriendRequestDao(db),
		GroupDao:         dao.NewGroupDao(db),
		GroupMemberDao:   dao.NewGroupMemberDao(db),
	}
}
