/*
这是一个插入排序算法相关模块
*/
package sorting


/*
载入依赖包
*/
// InsertionSort 插入排序（对外暴露）
func InsertionSort(nums []int) {
	// 外循环：已排序区间为 [0, i-1]
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		// 内循环：将 base 插入到已排序区间 [0, i-1] 中的正确位置
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j] // 将 nums[j] 向右移动一位
			j--
		}
		nums[j+1] = base // 将 base 赋值到正确位置
	}
}
