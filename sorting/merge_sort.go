/*
这是一个归并排序算法相关模块
*/
package sorting


/*
载入依赖包
*/
// merge 内部合并逻辑（小写私有）
func merge(nums []int, left, mid, right int) {
	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}

	for idx := 0; idx < len(tmp); idx++ {
		nums[left+idx] = tmp[idx]
	}
}

// mergeSort 内部递归
func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

// MergeSort 对外暴露入口函数
func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}
