package main

import (
	"fmt"
	"shihua_algo/queue"
)

func main() {
	fmt.Println("初始化队列")
	// 初始化操作
	q := queue.CreateNewLinkedListQueue()
	// 入队
	queue.Pushqueue(q, 1)
	queue.Pushqueue(q, 2)
	queue.Pushqueue(q, 3)
	// 出队
	element := queue.Popqueue(q)
	fmt.Println(element)
	// 访问队首元素
	element_peek := queue.Peekqueue(q)
	fmt.Println(element_peek)
}