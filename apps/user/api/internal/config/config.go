package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	UserRpc zrpc.RpcClientConf // user rpc 客户端配置

	JwtAuth struct {
		AccessSecret string
		//AccessExpire int64
	}
}
