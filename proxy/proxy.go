package proxy

import (
	"fmt"
)

// 服务接口
type ServiceInterface interface {
	Execute(access string)
}

// 服务实现了用于执行任务的 ServiceInterface 接口
type Service struct {
}

// 服务对象的方法
func (t *Service) Execute(access string) {
	fmt.Println("Proxy Service: " + access)
}

// 代理对象
type Proxy struct {
	realService *Service
}

// 创建代理对象
func NewProxy() *Proxy {
	return &Proxy{realService: &Service{}}
}

// 拦截 Execute 命令并将其重新路由到服务命令
func (t *Proxy) Execute(access string) {
	if access == "yes" {
		t.realService.Execute(access)
	}
}