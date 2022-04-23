package main

func findSubstring(s string, words []string) []int {
	wordslength := 0
	singleWordLength := len(words[0])
	wordsCnt := map[string]int{}
	for _, w := range words {
		wordslength += len(w)
		if _, exist := wordsCnt[w]; !exist {
			wordsCnt[w] = 0
		}
		wordsCnt[w]++
	}

	ans := []int{}
	for i := 0; i <= len(s)-wordslength; i++ {
		splitWord := s[i : i+wordslength]
		if valid(splitWord, singleWordLength, wordsCnt) {
			ans = append(ans, i)
		}
	}

	return ans
}

func CompareMap(map1, map2 map[string]int) bool {
	for k, v := range map1 {
		if val, exist := map2[k]; !exist || v != val {
			return false
		}
	}

	for k, v := range map2 {
		if val, exist := map1[k]; !exist || v != val {
			return false
		}
	}

	return true
}

func valid(str string, wordLength int, words map[string]int) bool {
	cnt := map[string]int{}
	for i := 0; i < len(str); i = i + wordLength {
		splitWord := str[i : i+wordLength]
		if _, exist := cnt[splitWord]; !exist {
			cnt[splitWord] = 0
		}
		cnt[splitWord]++

	}

	return CompareMap(cnt, words)
}
