package main

import (
	"fmt"
	"shihua_algo/tree"
)

func main() {
	fmt.Println("===== AVL 平衡二叉树 测试示例 =====")
	// 初始化空AVL树
	avl := tree.NewAVLTree()
	fmt.Println("初始化空树，打印当前树结构：")
	avl.Print()

	// 1. 批量插入节点（触发自动旋转平衡）
	fmt.Println("\n=== 执行Insert批量添加数据 ===")
	nums := []int{8, 4, 12, 2, 6, 10, 14}
	for _, v := range nums {
		avl.Insert(v)
	}
	fmt.Println("批量插入 8,4,12,2,6,10,14 后树层序结构：")
	avl.Print()
	fmt.Printf("AVL中序遍历(升序): %v\n", avl.InOrder())

	// 2. 插入重复值（不会生效）
	fmt.Println("\n=== 插入重复数字 8，无变化 ===")
	avl.Insert(8)
	avl.Print()

	// 3. 插入新节点，触发平衡调整
	fmt.Println("\n=== 插入新节点 5 ===")
	avl.Insert(5)
	fmt.Println("插入5后树结构（自动平衡）：")
	avl.Print()
	fmt.Printf("更新后中序遍历: %v\n", avl.InOrder())

	// 4. Search 查询节点
	fmt.Println("\n=== 查询节点 ===")
	searchVal := 6
	findNode := avl.Search(searchVal)
	if findNode != nil {
		fmt.Printf("找到值为 %d 的节点\n", findNode.Val)
	} else {
		fmt.Printf("未找到值为 %d 的节点\n", searchVal)
	}

	searchVal2 := 99
	findNode2 := avl.Search(searchVal2)
	if findNode2 != nil {
		fmt.Printf("找到值为 %d 的节点\n", findNode2.Val)
	} else {
		fmt.Printf("未找到值为 %d 的节点\n", searchVal2)
	}

	// 5. 删除节点演示（删除后自动平衡）
	fmt.Println("\n=== 删除节点4（双子树，删除后自动平衡） ===")
	avl.Remove(4)
	fmt.Println("删除4后树结构：")
	avl.Print()
	fmt.Printf("删除后中序遍历: %v\n", avl.InOrder())

	fmt.Println("\n=== 删除叶子节点2 ===")
	avl.Remove(2)
	avl.Print()
	fmt.Printf("删除后中序遍历: %v\n", avl.InOrder())

	// 6. 删除不存在节点，无报错无变化
	fmt.Println("\n=== 删除不存在数字99，无任何变化 ===")
	avl.Remove(99)
	avl.Print()

	// 7. 批量插入倾斜数据，强制触发多次旋转（LL、LR、RR、RL）
	fmt.Println("\n===== 批量插入倾斜序列 [1,2,3,4,5,6,7] 验证自动平衡 =====")
	avl2 := tree.NewAVLTree()
	slopeNums := []int{1, 2, 3, 4, 5, 6, 7}
	for _, v := range slopeNums {
		avl2.Insert(v)
	}
	avl2.Print()
	fmt.Printf("倾斜序列中序遍历: %v\n", avl2.InOrder())

	// 8. 清空树：循环删除所有根节点
	fmt.Println("\n=== 逐步删除所有节点，清空树 ===")
	for avl.GetRoot() != nil {
		avl.Remove(avl.GetRoot().Val)
	}
	fmt.Println("清空后：")
	avl.Print()
}
