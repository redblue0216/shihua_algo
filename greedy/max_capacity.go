/*
这是一个贪心-盛水容器相关模块
*/
package greedy

/*
载入依赖包
*/
import "math"

// MaxCapacity 最大盛水容量：贪心双指针（首字母大写导出）
func MaxCapacity(ht []int) int {
	// 初始化 i, j，使其分列数组两端
	i, j := 0, len(ht)-1
	// 初始最大容量为 0
	res := 0
	// 循环贪心选择，直至两板相遇
	for i < j {
		// 更新最大容量
		capacity := int(math.Min(float64(ht[i]), float64(ht[j]))) * (j - i)
		res = int(math.Max(float64(res), float64(capacity)))
		// 向内移动短板
		if ht[i] < ht[j] {
			i++
		} else {
			j--
		}
	}
	return res
}