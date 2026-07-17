/*
这是一个堆相关模块
*/
package heap


import (
	"fmt"
	"math"
)

/*
基本数据结构--堆
技术实现--数组大顶堆
*/
// 数组表示下的大顶堆
type MaxHeap struct {
	// 使用切片而非数组，这样无须考虑扩容问题
	Data []any
}

// NewHeap 构造函数，建立空堆（对外暴露）
func NewHeap() *MaxHeap {
	return &MaxHeap{
		Data: make([]any, 0),
	}
}

// NewMaxHeap 构造函数，根据切片建堆（对外暴露）
func NewMaxHeap(nums []any) *MaxHeap {
	// 将列表元素原封不动添加进堆
	h := &MaxHeap{Data: nums}
	for i := h.Parent(len(h.Data) - 1); i >= 0; i-- {
		// 堆化除叶节点以外的其他所有节点
		h.SiftDown(i)
	}
	return h
}

// Left 获取左子节点的索引
func (h *MaxHeap) Left(i int) int {
	return 2*i + 1
}

// Right 获取右子节点的索引
func (h *MaxHeap) Right(i int) int {
	return 2*i + 2
}

// Parent 获取父节点的索引
func (h *MaxHeap) Parent(i int) int {
	// 向下整除
	return (i - 1) / 2
}

// Swap 交换元素
func (h *MaxHeap) Swap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

// Size 获取堆大小
func (h *MaxHeap) Size() int {
	return len(h.Data)
}

// IsEmpty 判断堆是否为空
func (h *MaxHeap) IsEmpty() bool {
	return len(h.Data) == 0
}

// Peek 访问堆顶元素
func (h *MaxHeap) Peek() any {
	return h.Data[0]
}

// Push 元素入堆
func (h *MaxHeap) Push(val any) {
	// 添加节点
	h.Data = append(h.Data, val)
	// 从底至顶堆化
	h.SiftUp(len(h.Data) - 1)
}

// SiftUp 从节点 i 开始，从底至顶堆化
func (h *MaxHeap) SiftUp(i int) {
	for true {
		// 获取节点 i 的父节点
		p := h.Parent(i)
		// 当“越过根节点”或“节点无须修复”时，结束堆化
		if p < 0 || h.Data[i].(int) <= h.Data[p].(int) {
			break
		}
		// 交换两节点
		h.Swap(i, p)
		// 循环向上堆化
		i = p
	}
}

// Pop 元素出堆
func (h *MaxHeap) Pop() any {
	// 判空处理
	if h.IsEmpty() {
		fmt.Println("error: heap is empty")
		return nil
	}
	// 交换根节点与最右叶节点（交换首元素与尾元素）
	h.Swap(0, h.Size()-1)
	// 删除节点
	val := h.Data[len(h.Data)-1]
	h.Data = h.Data[:len(h.Data)-1]
	// 从顶至底堆化
	h.SiftDown(0)
	// 返回堆顶元素
	return val
}

// SiftDown 从节点 i 开始，从顶至底堆化
func (h *MaxHeap) SiftDown(i int) {
	for true {
		// 判断节点 i, l, r 中值最大的节点，记为 max
		l, r, max := h.Left(i), h.Right(i), i
		if l < h.Size() && h.Data[l].(int) > h.Data[max].(int) {
			max = l
		}
		if r < h.Size() && h.Data[r].(int) > h.Data[max].(int) {
			max = r
		}
		// 若节点 i 最大或索引 l, r 越界，则无须继续堆化，跳出
		if max == i {
			break
		}
		// 交换两节点
		h.Swap(i, max)
		// 循环向下堆化
		i = max
	}
}

// Print 打印堆（二叉树格式）
func (h *MaxHeap) Print() {
	PrintHeap(h.Data)
}

// PrintHeap 内置堆打印函数，替代外部pkg依赖
func PrintHeap(h []any) {
	if len(h) == 0 {
		fmt.Println("Heap is empty")
		return
	}
	n := len(h)
	// 计算堆的层数
	level := int(math.Log2(float64(n))) + 1
	start := 0
	for l := 0; l < level; l++ {
		// 当前层节点数量
		count := int(math.Pow(2, float64(l)))
		end := start + count
		if end > n {
			end = n
		}
		// 每层前置空格
		space := int(math.Pow(2, float64(level-l))) - 1
		for i := 0; i < space; i++ {
			fmt.Print(" ")
		}
		// 打印当前层所有节点
		for i := start; i < end; i++ {
			fmt.Printf("%d ", h[i].(int))
			// 节点间隔空格
			interval := int(math.Pow(2, float64(level-l+1))) - 1
			for j := 0; j < interval; j++ {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		start = end
	}
}
