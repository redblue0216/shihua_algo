package main

import (
	"fmt"
	"shihua_algo/builder"
)

func main() {
	b := builder.NewConcreteBuilder()
	director := builder.NewDirector(&b)
	director.Construct()

	product := b.GetResult()
	fmt.Printf("产品构建状态: Built=%t\n", product.Built)

	fmt.Println("建造者模式演示完成")
}