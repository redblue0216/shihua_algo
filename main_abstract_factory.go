package main

import (
	"fmt"
	"shihua_algo/abstract_factory"
)

func main() {
	factory := abstract_factory.NewConcreteFactory()
	product := factory.CreateProduct()
	product.GetName()

	fmt.Println("抽象工厂模式演示完成")
}