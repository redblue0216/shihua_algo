/*
这是一个分治算法相关模块
*/
package divide_and_conquer

/*
载入依赖包
*/
import "container/list"

// move 内部私有方法，移动单个圆盘
func move(src, tar *list.List) {
	// 从 src 顶部拿出一个圆盘
	pan := src.Back()
	// 将圆盘放入 tar 顶部
	tar.PushBack(pan.Value)
	// 移除 src 顶部圆盘
	src.Remove(pan)
}

// DfsHanota 对外暴露的递归汉诺塔DFS
func DfsHanota(i int, src, buf, tar *list.List) {
	// 若 src 只剩下一个圆盘，则直接将其移到 tar
	if i == 1 {
		move(src, tar)
		return
	}
	// 子问题 f(i-1) ：将 src 顶部 i-1 个圆盘借助 tar 移到 buf
	DfsHanota(i-1, src, tar, buf)
	// 子问题 f(1) ：将 src 剩余一个圆盘移到 tar
	move(src, tar)
	// 子问题 f(i-1) ：将 buf 顶部 i-1 个圆盘借助 src 移到 tar
	DfsHanota(i-1, buf, src, tar)
}

// SolveHanota 对外暴露完整汉诺塔求解入口
func SolveHanota(A, B, C *list.List) {
	n := A.Len()
	// 将 A 顶部 n 个圆盘借助 B 移到 C
	DfsHanota(n, A, B, C)
}
