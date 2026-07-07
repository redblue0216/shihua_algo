/*
这是一个队列相关模块
*/
package queue

/*
载入依赖包
*/
import "container/list"


/*
基本数据结构--队列
技术实现--链表
*/
// 使用内置包list实现队列
type LinkedListQueue struct {
	data *list.List
}

// 初始化队列
func CreateNewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		data: list.New(),
	}
}

// 入队
func Pushqueue(queue *LinkedListQueue, value int) {
	queue.data.PushBack(value) // 使用list的PushBack函数，往链表的尾部添加一个新元素
}

// 出队
func Popqueue(queue *LinkedListQueue) int {
	if queue.data.Len() == 0 { // 判断队列是否为空
		return -1
	}
	first_element := queue.data.Front() // 使用list的Front函数，获取链表第一个元素
	queue.data.Remove(first_element) // 使用list的Remove函数，删除链表中的某个节点
	return first_element.Value.(int) // 强制类型转换
}

// 访问队首元素
func Peekqueue(queue *LinkedListQueue) int {
	if queue.data.Len() == 0 { // 判断队列是否为空
		return -1
	}
	element := queue.data.Front() // 使用list的Front函数，获取链表第一个元素
	return element.Value.(int)
}
