package main

import "shihua_algo/visitor"

func main() {
	//声明具体元素A
	concreteElementA := &visitor.ConcreteElementA{}
	//调用具体元素A的方法
	concreteElementA.FeatureA()
	//具体元素A接受具体访问者A
	concreteElementA.Accept(&visitor.ConcreteVisitorA{})

	//声明具体元素B
	concreteElementB := &visitor.ConcreteElementB{}
	//调用具体元素B的方法
	concreteElementB.FeatureB()
	//具体元素B接受具体访问者B
	concreteElementB.Accept(&visitor.ConcreteVisitorB{})
}

//$ go run main.go
//具体访问者A ConcreteElementA
//具体访问者B ConcreteElementB