package common

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"strconv"
)

//设置注册中心
func GetConsulRegistry(host string, port int) registry.Registry {
	reg := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{host + ":" + strconv.Itoa(port)}
		},
	)
	return reg
}
