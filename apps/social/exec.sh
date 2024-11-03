# 生成social rpc代码（放项目跟目录下执行）
goctl rpc protoc apps/social/rpc/social.proto --go_out=./apps/social/rpc --go-grpc_out=./apps/social/rpc --zrpc_out=./apps/social/rpc -style go_zero
