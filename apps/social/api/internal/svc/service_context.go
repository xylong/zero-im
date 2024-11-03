package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-im/apps/social/api/internal/config"
	"zero-im/apps/social/rpc/social_client"
	"zero-im/apps/user/rpc/user_client"
)

type ServiceContext struct {
	Config config.Config

	social_client.Social
	user_client.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Social: social_client.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
		User:   user_client.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
