package main

// 根據題目：-3 * 10^4 <= xi, yi <= 3 * 10^4
const totalRows = 60001
const offset = 30000

func robotSim(commands []int, obstacles [][]int) int {
	// 1. 透過unordered set記錄障礙物
	ans := 0
	set := NewSet()
	for _, obs := range obstacles {
		set.add(obs[0], obs[1])
	}

	x, y := 0, 0
	// 2. 記錄方向 N=0, E=1, S=2, W=3
	dir := 0
	// 3. 通過兩個方向陣列分別表示x與y在對應的方向應該移動的偏移量
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for _, cmd := range commands {
		// 左轉: (dir-1+4)%4
		// 右轉: (dir+1)%4
		if cmd == -2 {
			dir = (dir + 3) % 4
		} else if cmd == -1 {
			dir = (dir + 1) % 4
		} else {
			for i := 0; i < cmd; i++ {
				nx, ny := x+dx[dir], y+dy[dir]
				if set.find(nx, ny) {
					break
				}
				x, y = nx, ny
				dist := x*x + y*y
				if dist > ans {
					ans = dist
				}
			}
		}
	}

	return ans
}

type set struct {
	values map[int]struct{}
}

func (s *set) add(x, y int) {
	s.values[hash(x, y)] = struct{}{}
}

func (s *set) find(x, y int) bool {
	_, exist := s.values[hash(x, y)]
	return exist
}

func NewSet() set {
	return set{map[int]struct{}{}}
}

func hash(x, y int) int {
	return totalRows*(x+offset) + (y + offset)
}
