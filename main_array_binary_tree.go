package main

import (
	"fmt"
	"shihua_algo/tree"
)

func main() {
	fmt.Println("===== 数组实现二叉树 测试示例 =====")
	// 构建数组存储的二叉树数组，nil 代表空节点
	// 树结构：
	//        1
	//      /   \
	//     2     3
	//    / \   /
	//   4  nil 5
	arr := []any{1, 2, 3, 4, nil, 5}
	// 初始化数组二叉树
	abt := tree.NewArrayBinaryTree(arr)

	// 1. 获取树节点数量
	fmt.Println("\n1. 获取二叉树节点数组容量")
	fmt.Printf("数组长度 size = %d\n", abt.Size())

	// 2. 获取指定索引节点值、左右孩子、父节点索引
	fmt.Println("\n2. 索引0（根节点）相关信息")
	rootIdx := 0
	fmt.Printf("索引%d 节点值: %v\n", rootIdx, abt.Val(rootIdx))
	fmt.Printf("左子节点索引: %d\n", abt.Left(rootIdx))
	fmt.Printf("右子节点索引: %d\n", abt.Right(rootIdx))

	fmt.Println("\n索引1（值=2）相关信息")
	node1 := 1
	fmt.Printf("索引%d 节点值: %v\n", node1, abt.Val(node1))
	fmt.Printf("父节点索引: %d\n", abt.Parent(node1))
	fmt.Printf("左子节点索引: %d, 值=%v\n", abt.Left(node1), abt.Val(abt.Left(node1)))
	fmt.Printf("右子节点索引: %d, 值=%v\n", abt.Right(node1), abt.Val(abt.Right(node1)))

	// 测试越界索引取值
	fmt.Println("\n测试越界索引 i=99: val = %v", abt.Val(99))

	// 3. 层序遍历
	fmt.Println("\n3. 层序遍历 levelOrder")
	levelRes := abt.LevelOrder()
	fmt.Printf("层序结果: %v\n", levelRes)

	// 4. 前序、中序、后序深度优先遍历
	fmt.Println("\n4. 深度优先遍历")
	preRes := abt.PreOrder()
	fmt.Printf("前序遍历 preOrder: %v\n", preRes)

	inRes := abt.InOrder()
	fmt.Printf("中序遍历 inOrder: %v\n", inRes)

	postRes := abt.PostOrder()
	fmt.Printf("后序遍历 postOrder: %v\n", postRes)

	// 演示另一棵带更多空节点的二叉树
	fmt.Println("\n===== 新建一棵复杂二叉树测试 =====")
	arr2 := []any{10, 20, nil, 30, 40, nil, nil, nil, 50}
	tree2 := tree.NewArrayBinaryTree(arr2)
	fmt.Printf("数组容量: %d\n", tree2.Size())
	fmt.Printf("层序遍历: %v\n", tree2.LevelOrder())
	fmt.Printf("后序遍历: %v\n", tree2.PostOrder())
}
