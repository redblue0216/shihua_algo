package factory

import "fmt"

// 工厂接口
type Factory interface {
	FactoryMethod(owner string) Product
}

// 具体工厂
type ConcreteFactory struct {
}

// 具体工厂的工厂方法
func (cf *ConcreteFactory) FactoryMethod(owner string) Product {
	p := &ConcreteProduct{}
	return p
}

// 产品
type Product interface {
	Use()
}

//具体产品
type ConcreteProduct struct {
}

//具体产品的方法
func (p *ConcreteProduct) Use() {
	fmt.Println("This is a concrete product")
}