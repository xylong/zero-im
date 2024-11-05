package main

import (
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
)

// TestSt 配置结构定义
type TestSt struct {
	Name string `json:"name"`
}

func main() {
	ss := subscriber.MustNewEtcdSubscriber(subscriber.EtcdConf{
		Hosts: []string{"localhost:2379"}, // etcd 地址
		Key:   "test",                     // 配置key
	})

	cc := configurator.MustNewConfigCenter[TestSt](configurator.Config{
		Type: "json", // 配置值类型：json,yaml,toml
	}, ss)

	v, err := cc.GetConfig()
	if err != nil {
		panic(err)
	}
	println(v.Name)

	// 如果想监听配置变化，可以添加 listener
	cc.AddListener(func() {
		v, err := cc.GetConfig()
		if err != nil {
			panic(err)
		}
		println(v.Name)
	})

	select {}
}
