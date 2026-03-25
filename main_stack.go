package main

import (
	"fmt"
	"shihua_algo/stack"
)

func main() {
	fmt.Println("初始化栈")
	// 初始化操作
	s := stack.CreateNewLinkedListStack()
	// 入栈
	stack.PushStack(s, 1)
	stack.PushStack(s, 2)
	stack.PushStack(s, 3)
	// 出栈
	element := stack.PopStack(s)
	fmt.Println(element)
	// 访问栈顶元素
	element_peek := stack.PeekStack(s)
	fmt.Println(element_peek)
}