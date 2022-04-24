package main

const MaxDist = 50000

func findShortestSubArray(nums []int) int {
	candidates := []int{}
	numMap := map[int]*Num{}
	maxCount := 0
	for i, num := range nums {
		if _, ok := numMap[num]; !ok {
			numMap[num] = NewNum(num)
			numMap[num].Start = i
		}
		numMap[num].End = i
		numMap[num].Count++

		if len(candidates) == 0 || numMap[num].Count > maxCount {
			candidates = []int{num}
			maxCount = numMap[num].Count
		} else if numMap[num].Count == maxCount {
			candidates = append(candidates, num)
		}
	}

	dist := MaxDist

	for _, num := range candidates {
		if candidatesDist := numMap[num].End - numMap[num].Start; candidatesDist < dist {
			dist = candidatesDist
		}
	}

	return dist + 1
}

type Num struct {
	value int
	Start int
	End   int
	Count int
}

func NewNum(num int) *Num {
	return &Num{value: num}
}
