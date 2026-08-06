package bridge

import "fmt"

//实现接口
type Implementor interface {
	Implementation(str string)
}

//具体实现
type ConcreteImplementor struct{}

func (*ConcreteImplementor) Implementation(str string) {
	fmt.Printf("打印信息：[%v]", str)
}

//初始化具体实现对象
func NewConcreteImplementor() *ConcreteImplementor {
	return &ConcreteImplementor{}
}

//抽象接口
type Abstraction interface {
	Execute(str string)
}

//扩充抽象
type RefinedAbstraction struct {
	method Implementor
}

//扩充抽象方法
func (c *RefinedAbstraction) Execute(str string) {
	c.method.Implementation(str)
}

//初始化扩充抽象对象
func NewRefinedAbstraction(im Implementor) *RefinedAbstraction {
	return &RefinedAbstraction{method: im}
}