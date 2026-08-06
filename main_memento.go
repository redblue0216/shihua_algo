package main

import (
	"fmt"
	"shihua_algo/memento"
)

func main() {
	//声明负责人对象
	Caretaker := &memento.Caretaker{
		History: make([]*memento.Memento, 0),
	}

	//声明原发器对象
	n := memento.NewOriginator(100)

	//添加备忘录
	Caretaker.AddMemento(n.CreateMemento())
	n.TenTimes()
	fmt.Printf("Originator 当前的值: %d\n", n.Value())

	//添加备忘录
	Caretaker.AddMemento(n.CreateMemento())
	n.TenTimes()
	fmt.Printf("Originator 当前的值: %d\n", n.Value())

	//恢复原发器对象的值
	n.RestoreMemento(Caretaker.GetMemento(0))
	fmt.Printf("恢复备忘录后 Originator 当前的值: %d\n", n.Value())
}

//$ go run main.go
//Originator 当前的值: 1000
//Originator 当前的值: 10000
//恢复备忘录后 Originator 当前的值: 100