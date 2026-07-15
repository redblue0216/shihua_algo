package main

import (
	"fmt"
	"shihua_algo/tree"
)

func main() {
	fmt.Println("===== 链表实现二叉搜索树 BST 测试示例 =====")
	// 初始化空二叉搜索树
	bst := tree.NewBinarySearchTree()
	fmt.Println("初始化空树，打印当前树结构：")
	bst.Print()

	// 1. 批量插入节点
	fmt.Println("\n=== 执行Insert批量添加数据 ===")
	nums := []int{8, 4, 12, 2, 6, 10, 14}
	for _, v := range nums {
		bst.Insert(v)
	}
	fmt.Println("批量插入 8,4,12,2,6,10,14 后树层序结构：")
	bst.Print()
	fmt.Printf("BST中序遍历(升序): %v\n", bst.InOrder())

	// 2. 插入重复值（不会生效）
	fmt.Println("\n=== 插入重复数字 8，无变化 ===")
	bst.Insert(8)
	bst.Print()

	// 3. 插入新节点
	fmt.Println("\n=== 插入新节点 5 ===")
	bst.Insert(5)
	fmt.Println("插入5后树结构：")
	bst.Print()
	fmt.Printf("更新后中序遍历: %v\n", bst.InOrder())

	// 4. Search 查询节点
	fmt.Println("\n=== 查询节点 ===")
	searchVal := 6
	findNode := bst.Search(searchVal)
	if findNode != nil {
		fmt.Printf("找到值为 %d 的节点\n", findNode.Val)
	} else {
		fmt.Printf("未找到值为 %d 的节点\n", searchVal)
	}

	searchVal2 := 99
	findNode2 := bst.Search(searchVal2)
	if findNode2 != nil {
		fmt.Printf("找到值为 %d 的节点\n", findNode2.Val)
	} else {
		fmt.Printf("未找到值为 %d 的节点\n", searchVal2)
	}

	// 5. 删除节点演示
	fmt.Println("\n=== 删除节点4（拥有左右子树） ===")
	bst.Remove(4)
	fmt.Println("删除4后树结构：")
	bst.Print()
	fmt.Printf("删除后中序遍历: %v\n", bst.InOrder())

	fmt.Println("\n=== 删除叶子节点2 ===")
	bst.Remove(2)
	bst.Print()
	fmt.Printf("删除后中序遍历: %v\n", bst.InOrder())

	// 6. 删除不存在节点，无报错
	fmt.Println("\n=== 删除不存在数字99，无任何变化 ===")
	bst.Remove(99)
	bst.Print()

	// 7. 清空树：循环删除所有根节点
	fmt.Println("\n=== 逐步删除所有节点，清空树 ===")
	for bst.GetRoot() != nil {
		bst.Remove(bst.GetRoot().Val)
	}
	fmt.Println("清空后：")
	bst.Print()
}
