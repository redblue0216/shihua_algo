package adapter

import (
	"fmt"
)

// 定义被适配的类
type Adaptee struct {
}

// 目标接口
type Target interface {
	Execute()
}

//定义了用于执行的方法SpecificExecute()
func (a *Adaptee) SpecificExecute() {
	fmt.Println("最终执行的方法")
}

// Adapter 是新接口 Target 的适配器，继承了 Adaptee 类
type Adapter struct {
	*Adaptee
}

// 实现 Target 接口，同时继承了 Adaptee 类
func (a *Adapter) Execute() {
	a.SpecificExecute()
}