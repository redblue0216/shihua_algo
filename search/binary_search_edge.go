/*
这是一个二分查找边界相关模块
*/
package search

/*
载入依赖包
*/

/*
基本算法--二分查找边界
技术实现--有序数组
*/

// BinarySearchLeftEdge 二分查找最左一个 target
func BinarySearchLeftEdge(nums []int, target int) int {
	// 等价于查找 target 的插入点
	i := BinarySearchInsertion(nums, target)
	// 未找到 target ，返回 -1
	if i == len(nums) || nums[i] != target {
		return -1
	}
	// 找到 target ，返回索引 i
	return i
}

// BinarySearchRightEdge 二分查找最右一个 target
func BinarySearchRightEdge(nums []int, target int) int {
	// 转化为查找最左一个 target + 1
	i := BinarySearchInsertion(nums, target+1)
	// j 指向最右一个 target ，i 指向首个大于 target 的元素
	j := i - 1
	// 未找到 target ，返回 -1
	if j == -1 || nums[j] != target {
		return -1
	}
	// 找到 target ，返回索引 j
	return j
}