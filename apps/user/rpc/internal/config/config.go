package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-im/pkg/core/store/gormc"
	"zero-im/pkg/ctxdata"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql gormc.Config
	Jwt   ctxdata.Jwt
}
