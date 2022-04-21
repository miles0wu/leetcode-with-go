package main

func twoSum(nums []int, target int) []int {
	set := map[int]int{}
	for i, v := range nums {
		expect := target - v
		if j, exist := set[expect]; exist {
			return []int{j, i}
		}
		set[v] = i
	}
	return []int{}
}
