package main

import "shihua_algo/decorator"

func main() {
	concreteComponent := &decorator.ConcreteComponent{}
	decoratorA := &decorator.DecoratorA{}
	decoratorB := &decorator.DecoratorB{}
	decoratorA.SetComponent(concreteComponent)
	decoratorB.SetComponent(decoratorA)
	decoratorB.Operation()
}

//$ go run main.go
//具体的对象开始操作...
//装饰A扩展的方法~
//装饰B扩展的方法~