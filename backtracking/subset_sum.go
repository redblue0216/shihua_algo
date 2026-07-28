/*
这是一个回溯算法相关模块：子集和
*/
package backtracking

import "sort"

// backtrackSubsetSumI 私有递归，子集和I：元素可重复选取，数组无重复
func backtrackSubsetSumI(start, target int, state, choices *[]int, res *[][]int) {
	if target == 0 {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
		return
	}

	for i := start; i < len(*choices); i++ {
		if target-(*choices)[i] < 0 {
			break
		}

		*state = append(*state, (*choices)[i])
		backtrackSubsetSumI(i, target-(*choices)[i], state, choices, res)
		*state = (*state)[:len(*state)-1]
	}
}

// SubsetSumI 对外导出：可重复选元素，找出和为target的所有子集（允许复用元素）
func SubsetSumI(nums []int, target int) [][]int {
	state := make([]int, 0)
	sort.Ints(nums)
	start := 0
	res := make([][]int, 0)
	backtrackSubsetSumI(start, target, &state, &nums, &res)
	return res
}

// backtrackSubsetSumII 私有递归，子集和II：元素不可重复选，数组存在重复数字，结果去重
func backtrackSubsetSumII(start, target int, state, choices *[]int, res *[][]int) {
	if target == 0 {
		newState := append([]int{}, *state...)
		*res = append(*res, newState)
		return
	}

	for i := start; i < len(*choices); i++ {
		if target-(*choices)[i] < 0 {
			break
		}

		if i > start && (*choices)[i] == (*choices)[i-1] {
			continue
		}

		*state = append(*state, (*choices)[i])
		backtrackSubsetSumII(i+1, target-(*choices)[i], state, choices, res)
		*state = (*state)[:len(*state)-1]
	}
}

// SubsetSumII 对外导出：每个元素只能选1次，数组含重复，返回无重复子集
func SubsetSumII(nums []int, target int) [][]int {
	state := make([]int, 0)
	sort.Ints(nums)
	start := 0
	res := make([][]int, 0)
	backtrackSubsetSumII(start, target, &state, &nums, &res)
	return res
}