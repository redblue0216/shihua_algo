package main

import (
	"fmt"
	"shihua_algo/hash_map"
)

func main() {
	fmt.Println("初始化链式地址哈希表")
	// 初始化哈希表
	hm := hash_map.NewHashMapChaining()
	// 添加键值对 put
	fmt.Println("\n=== 执行put添加数据 ===")
	hm.Put(1, "张三")
	hm.Put(2, "李四")
	hm.Put(5, "王五")
	hm.Put(8, "赵六")
	fmt.Println("添加后哈希表结构：")
	hm.Print()
	// 更新已有key的值
	fmt.Println("\n=== 更新key=2的值为「小李」 ===")
	hm.Put(2, "小李")
	hm.Print()
	// 查询 get
	fmt.Println("\n=== 查询key=5 ===")
	val := hm.Get(5)
	fmt.Printf("key=5 对应值：%s\n", val)
	fmt.Println("=== 查询不存在的key=99 ===")
	emptyVal := hm.Get(99)
	fmt.Printf("key=99 对应值：%q\n", emptyVal)
	// 删除 remove
	fmt.Println("\n=== 删除key=2 ===")
	hm.Remove(2)
	fmt.Println("删除后哈希表结构：")
	hm.Print()
	// 多次插入触发扩容演示
	fmt.Println("\n=== 批量插入大量数据触发扩容 ===")
	hm.Put(10, "A")
	hm.Put(11, "B")
	hm.Put(12, "C")
	hm.Put(13, "D")
	hm.Put(14, "E")
	hm.Print()
}
