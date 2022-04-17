package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		}
		digits[i] = 0
	}

	if digits[0] == 0 {
		tmp := make([]int, 0, len(digits)+1)
		tmp = append(tmp, 1)
		digits = append(tmp, digits...)
	}

	return digits
}
