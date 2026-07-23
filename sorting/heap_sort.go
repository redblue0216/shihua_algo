/*
这是一个堆排序算法相关模块
*/
package sorting


/*
载入依赖包
*/
// siftDown 内部堆化逻辑
func siftDown(nums *[]int, n, i int) {
	for {
		l := 2*i + 1
		r := 2*i + 2
		ma := i
		if l < n && (*nums)[l] > (*nums)[ma] {
			ma = l
		}
		if r < n && (*nums)[r] > (*nums)[ma] {
			ma = r
		}
		if ma == i {
			break
		}
		(*nums)[i], (*nums)[ma] = (*nums)[ma], (*nums)[i]
		i = ma
	}
}

// heapSort 内部逻辑
func heapSort(nums *[]int) {
	n := len(*nums)
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(nums, n, i)
	}
	for i := n - 1; i > 0; i-- {
		(*nums)[0], (*nums)[i] = (*nums)[i], (*nums)[0]
		siftDown(nums, i, 0)
	}
}

// HeapSort 对外入口，接收切片值，内部取地址
func HeapSort(nums []int) {
	heapSort(&nums)
}
