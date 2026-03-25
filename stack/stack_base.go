/*
这是一个栈相关模块
*/
package stack

/*
载入依赖包
*/
import "container/list"


/*
基本数据结构--栈
技术实现--链表
*/
// 使用内置包list实现栈
type LinkedListStack struct {
	data *list.List
}

// 初始化栈
func CreateNewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		data: list.New(),
	}
}

// 入栈
func PushStack(stack *LinkedListStack, value int) {
	stack.data.PushBack(value) // 使用list的PushBack函数，往链表的尾部添加一个新元素
}

// 出栈
func PopStack(stack *LinkedListStack) int {
	if stack.data.Len() == 0 { // 判断栈是否为空
		return -1
	}
	final_element := stack.data.Back() // 使用list的Back函数，获取链表最后一个元素
	stack.data.Remove(final_element) // 使用list的Remove函数，删除链表中的某个节点
	return final_element.Value.(int) // 强制类型转换
}

// 访问栈顶元素
func PeekStack(stack *LinkedListStack) int {
	if stack.data.Len() == 0 { // 判断栈是否为空
		return -1
	}
	element := stack.data.Back() // 使用list的Back函数，获取链表最后一个元素
	return element.Value.(int)
}