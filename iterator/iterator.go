package iterator

import (
	"fmt"
	"time"
)

// 迭代器接口
type Iterator interface {
	// 返回是否存在另一个下一个元素
	HasMore() bool

	// 递增迭代器，用于指向下一个元素
	GetNext()
}

//集合接口
type Collection interface {
	CreateIterator() Iterator
}

//具体集合
type ConcreteCollection struct {
}

//初始化具体集合对象，用于创建具体迭代器对象
func (u *ConcreteCollection) CreateIterator() Iterator {
	return &ConcreteIterator{
		IterationState: true,
	}
}

// 具体迭代器
type ConcreteIterator struct {
	IterationState bool
}

//具体迭代器的方法
func (i *ConcreteIterator) HasMore() bool {
	if i.IterationState == true {
		return true
	} else {
		return false
	}
}

//具体迭代器的方法，用于递增迭代器以指向下一个元素
func (i *ConcreteIterator) GetNext() {
	if i.HasMore() {
		time.Sleep(1 * time.Second)
		fmt.Println("GetNext")
	}
}