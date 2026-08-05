package main

import (
	"fmt"
	"shihua_algo/object_pool"
)

func main() {
	// 定义对象创建函数
	pool := object_pool.NewPool(func() interface{} {
		return fmt.Sprintf("池对象-%d", 0)
	})

	// 获取对象
	obj1 := pool.Acquire()
	obj2 := pool.Acquire()
	fmt.Printf("获取对象1: %v\n", obj1)
	fmt.Printf("获取对象2: %v\n", obj2)

	// 归还对象
	pool.Release(obj1)
	// 再次获取，复用归还的对象
	obj3 := pool.Acquire()
	fmt.Printf("再次获取对象: %v\n", obj3)

	fmt.Println("对象池模式演示完成")
}