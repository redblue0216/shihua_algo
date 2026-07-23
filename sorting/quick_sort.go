/*
这是一个快速排序算法相关模块
*/
package sorting


/*
载入依赖包
*/
// 快速排序基础结构体
type quickSort struct{}

// 快速排序（中位基准数优化）
type quickSortMedian struct{}

// 快速排序（递归深度优化/尾递归优化）
type quickSortTailCall struct{}

// 哨兵划分 基础快排
func (q *quickSort) partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func (q *quickSort) quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := q.partition(nums, left, right)
	q.quickSort(nums, left, pivot-1)
	q.quickSort(nums, pivot+1, right)
}

// QuickSort 对外调用入口：基础快速排序
func QuickSort(nums []int) {
	q := &quickSort{}
	q.quickSort(nums, 0, len(nums)-1)
}

// 三数取中优化
func (q *quickSortMedian) medianThree(nums []int, left, mid, right int) int {
	l, m, r := nums[left], nums[mid], nums[right]
	if (l <= m && m <= r) || (r <= m && m <= l) {
		return mid
	}
	if (m <= l && l <= r) || (r <= l && l <= m) {
		return left
	}
	return right
}

func (q *quickSortMedian) partition(nums []int, left, right int) int {
	med := q.medianThree(nums, left, (left+right)/2, right)
	nums[left], nums[med] = nums[med], nums[left]

	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func (q *quickSortMedian) quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := q.partition(nums, left, right)
	q.quickSort(nums, left, pivot-1)
	q.quickSort(nums, pivot+1, right)
}

// QuickSortMedian 对外入口：三数取中快排
func QuickSortMedian(nums []int) {
	q := &quickSortMedian{}
	q.quickSort(nums, 0, len(nums)-1)
}

// 递归深度优化快排
func (q *quickSortTailCall) partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func (q *quickSortTailCall) quickSort(nums []int, left, right int) {
	for left < right {
		pivot := q.partition(nums, left, right)
		if pivot-left < right-pivot {
			q.quickSort(nums, left, pivot-1)
			left = pivot + 1
		} else {
			q.quickSort(nums, pivot+1, right)
			right = pivot - 1
		}
	}
}

// QuickSortTailCall 对外入口：短数组优先递归防栈溢出
func QuickSortTailCall(nums []int) {
	q := &quickSortTailCall{}
	q.quickSort(nums, 0, len(nums)-1)
}
