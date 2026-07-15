/*
这是一个数组二叉树相关模块
*/
package tree


/*
基本数据结构--数组二叉树
技术实现--数组二叉树
*/
// 数组表示下的二叉树类
type ArrayBinaryTree struct {
	tree []any
}

// 构造方法
func NewArrayBinaryTree(arr []any) *ArrayBinaryTree {
	return &ArrayBinaryTree{
		tree: arr,
	}
}

// 列表容量
func (abt *ArrayBinaryTree) Size() int {
	return len(abt.tree)
}

// 获取索引为 i 节点的值
func (abt *ArrayBinaryTree) Val(i int) any {
	// 若索引越界，则返回 null ，代表空位
	if i < 0 || i >= abt.Size() {
		return nil
	}
	return abt.tree[i]
}

// 获取索引为 i 节点的左子节点的索引
func (abt *ArrayBinaryTree) Left(i int) int {
	return 2*i + 1
}

// 获取索引为 i 节点的右子节点的索引
func (abt *ArrayBinaryTree) Right(i int) int {
	return 2*i + 2
}

// 获取索引为 i 节点的父节点的索引
func (abt *ArrayBinaryTree) Parent(i int) int {
	return (i - 1) / 2
}

// 层序遍历
func (abt *ArrayBinaryTree) LevelOrder() []any {
	var res []any
	// 直接遍历数组
	for i := 0; i < abt.Size(); i++ {
		if abt.Val(i) != nil {
			res = append(res, abt.Val(i))
		}
	}
	return res
}

// 深度优先遍历
func (abt *ArrayBinaryTree) dfs(i int, order string, res *[]any) {
	// 若为空位，则返回
	if abt.Val(i) == nil {
		return
	}
	// 前序遍历
	if order == "pre" {
		*res = append(*res, abt.Val(i))
	}
	abt.dfs(abt.Left(i), order, res)
	// 中序遍历
	if order == "in" {
		*res = append(*res, abt.Val(i))
	}
	abt.dfs(abt.Right(i), order, res)
	// 后序遍历
	if order == "post" {
		*res = append(*res, abt.Val(i))
	}
}

// 前序遍历
func (abt *ArrayBinaryTree) PreOrder() []any {
	var res []any
	abt.dfs(0, "pre", &res)
	return res
}

// 中序遍历
func (abt *ArrayBinaryTree) InOrder() []any {
	var res []any
	abt.dfs(0, "in", &res)
	return res
}

// 后序遍历
func (abt *ArrayBinaryTree) PostOrder() []any {
	var res []any
	abt.dfs(0, "post", &res)
	return res
}
