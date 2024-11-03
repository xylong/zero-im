package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-im/pkg/core/store/gormc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql gormc.Config
}
