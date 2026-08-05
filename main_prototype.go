package main

import (
	"fmt"
	"shihua_algo/prototype"
)

func main() {
	origin := &prototype.ConcretePrototype{Name: "原始原型对象"}
	cloneObj := origin.Clone()

	fmt.Printf("原始对象名称：%s\n", origin.GetName())
	fmt.Printf("克隆对象名称：%s\n", cloneObj.GetName())

	fmt.Println("原型模式演示完成")
}