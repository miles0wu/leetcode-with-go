package main

import "fmt"

func subsets(nums []int) [][]int {
	var chosen = []int{}
	var ans = [][]int{}

	var recur func(int)
	recur = func(i int) {
		if i == len(nums) {
			fmt.Println(chosen)
			ans = append(ans, append([]int(nil), chosen...))
			return
		}

		// 當前處理層不選
		recur(i + 1)

		// 當前處理層選
		chosen = append(chosen, nums[i])
		recur(i + 1)

		// 還原此層操作
		chosen = chosen[:len(chosen)-1]
	}
	recur(0)

	return ans
}
