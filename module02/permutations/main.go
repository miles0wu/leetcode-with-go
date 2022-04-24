package main

func permute(nums []int) [][]int {
	var ans [][]int

	var recur func([]int, int, []int)
	recur = func(nums []int, numsLen int, chosen []int) {
		if len(nums) == 0 {
			ans = append(ans, append([]int{}, chosen...))
			return
		}

		for i := 0; i < numsLen; i++ {
			// 當前迭代的數值
			c := nums[i]
			// 選當前數值
			chosen = append(chosen, c)
			// 從未選取陣列移除已選取數值
			nums = append(nums[:i], nums[i+1:]...)
			// 重複子問題
			recur(nums, len(nums), chosen)
			// 還原
			nums = append(nums[:i], append([]int{c}, nums[i:]...)...)
			chosen = chosen[:len(chosen)-1]
		}
	}

	recur(nums, len(nums), []int{})

	return ans
}
