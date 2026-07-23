/*
这是一个基数排序算法相关模块
*/
package sorting


/*
载入依赖包
*/
import "math"


// digit 获取num第k位，exp=10^(k-1)，内部私有
func digit(num, exp int) int {
	return (num / exp) % 10
}

// countingSortDigit 按指定位计数排序，内部私有
func countingSortDigit(nums []int, exp int) {
	counter := make([]int, 10)
	n := len(nums)
	for i := 0; i < n; i++ {
		d := digit(nums[i], exp)
		counter[d]++
	}
	// 前缀和
	for i := 1; i < 10; i++ {
		counter[i] += counter[i-1]
	}
	// 倒序回填保证稳定
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		d := digit(nums[i], exp)
		j := counter[d] - 1
		res[j] = nums[i]
		counter[d]--
	}
	// 覆盖原数组
	for i := 0; i < n; i++ {
		nums[i] = res[i]
	}
}

// RadixSort 对外暴露基数排序入口
func RadixSort(nums []int) {
	max := math.MinInt
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	// 低位到高位逐位排序
	for exp := 1; max >= exp; exp *= 10 {
		countingSortDigit(nums, exp)
	}
}
