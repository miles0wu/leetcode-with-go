package main

func combine(n int, k int) [][]int {
	var chosen []int
	var ans [][]int

	var recur func(int)
	recur = func(i int) {
		// 邊界條件
		if len(chosen) == k {
			ans = append(ans, append([]int(nil), chosen...))
			return
		}

		// 提早判斷分支不滿足條件
		if i > n || len(chosen)+n-i+1 < k {
			return
		}

		// 不選
		recur(i + 1)

		// 選
		chosen = append(chosen, i)
		recur(i + 1)

		// 還原
		chosen = chosen[:len(chosen)-1]
	}

	recur(1)

	return ans
}
