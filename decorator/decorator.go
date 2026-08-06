package decorator

import "fmt"

//组件接口
type Component interface {
	Operation()
}

//具体组件
type ConcreteComponent struct {
}

//具体组件方法
func (c *ConcreteComponent) Operation() {
	fmt.Println("具体的对象开始操作...")
}

//装饰
type Decorator struct {
	component Component
}

//装饰设置组件方法
func (d *Decorator) SetComponent(c Component) {
	d.component = c
}

//装饰方法
func (d *Decorator) Operation() {
	if d.component != nil {
		d.component.Operation()
	}
}

//具体装饰器A
type DecoratorA struct {
	Decorator
}

//具体装饰器A的方法
func (d *DecoratorA) Operation() {
	d.component.Operation()
	d.IndependentMethod()
}

func (d *DecoratorA) IndependentMethod() {
	fmt.Println("装饰A扩展的方法~")
}

//具体装饰器B
type DecoratorB struct {
	Decorator
}

//具体装饰器B的方法
func (d *DecoratorB) Operation() {
	d.component.Operation()
	fmt.Println(d.String())
}

//具体装饰器B的拓展方法
func (d *DecoratorB) String() string {
	return "装饰B扩展的方法~"
}