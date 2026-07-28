/*
这是一个回溯算法相关模块：全排列
*/
package backtracking

// backtrackI 私有递归函数，全排列I，无重复数字
func backtrackI(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	// 当状态长度等于元素数量时，记录解
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
		return
	}
	// 遍历所有选择
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		// 剪枝：不允许重复选择元素
		if !(*selected)[i] {
			// 尝试：做出选择，更新状态
			(*selected)[i] = true
			*state = append(*state, choice)
			// 进行下一轮选择
			backtrackI(state, choices, selected, res)
			// 回退：撤销选择，恢复到之前的状态
			(*selected)[i] = false
			*state = (*state)[:len(*state)-1]
		}
	}
}

// PermutationsI 导出函数：无重复数组全排列
func PermutationsI(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, len(nums))
	backtrackI(&state, &nums, &selected, &res)
	return res
}

// backtrackII 私有递归函数，全排列II，处理含重复数字
func backtrackII(state *[]int, choices *[]int, selected *[]bool, res *[][]int) {
	// 当状态长度等于元素数量时，记录解
	if len(*state) == len(*choices) {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
		return
	}
	// 遍历所有选择
	duplicated := make(map[int]struct{})
	for i := 0; i < len(*choices); i++ {
		choice := (*choices)[i]
		// 剪枝：不允许重复选择元素 且 跳过同层等值元素避免重复排列
		if _, exists := duplicated[choice]; exists || (*selected)[i] {
			continue
		}
		// 标记当前层已经选过该数值，同层不再选相同值
		duplicated[choice] = struct{}{}
		// 尝试选择
		(*selected)[i] = true
		*state = append(*state, choice)
		backtrackII(state, choices, selected, res)
		// 回退
		(*selected)[i] = false
		*state = (*state)[:len(*state)-1]
	}
}

// PermutationsII 导出函数：含重复数组，返回去重全排列
func PermutationsII(nums []int) [][]int {
	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, len(nums))
	backtrackII(&state, &nums, &selected, &res)
	return res
}