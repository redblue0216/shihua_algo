/*
这是一个计数排序算法相关模块
*/
package sorting


/*
载入依赖包
*/
// CountingSort 简易计数排序（不稳定，仅整数）
func CountingSortNaive(nums []int) {
	// 1. 统计数组最大元素 m
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	// 2. 统计各数字的出现次数
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}
	// 3. 遍历 counter 回填原数组
	for i, num := 0, 0; num < m+1; num++ {
		for j := 0; j < counter[num]; j++ {
			nums[i] = num
			i++
		}
	}
}

// CountingSort 稳定完整版计数排序
func CountingSort(nums []int) {
	// 1. 统计数组最大元素 m
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	// 2. 统计频次
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}
	// 3. 前缀和
	for i := 0; i < m; i++ {
		counter[i+1] += counter[i]
	}
	// 4. 倒序遍历保证稳定
	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		res[counter[num]-1] = num
		counter[num]--
	}
	copy(nums, res)
}
