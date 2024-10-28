package svc

import (
	"gorm.io/gorm"
	"zero-im/apps/user/rpc/internal/config"
	"zero-im/apps/user/rpc/internal/dao"
	"zero-im/pkg/core/store/gormc"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB

	UserDao dao.UserDao
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := gormc.NewEngine(c.Mysql)

	return &ServiceContext{
		Config: c,
		DB:     db,

		UserDao: dao.NewUserDao(db),
	}
}
