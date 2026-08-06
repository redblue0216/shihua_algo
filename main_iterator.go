package main

import (
	"shihua_algo/iterator"
)

func main() {

	//声明具体集合对象
	concreteCollection := &iterator.ConcreteCollection{}

	//声明具体迭代器对象
	iterator_instance := concreteCollection.CreateIterator()

	//执行具体方法
	for iterator_instance.HasMore() {
		iterator_instance.GetNext()
	}
}