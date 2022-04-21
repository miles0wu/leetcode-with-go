package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, str := range strs {
		key := hash(str)
		if _, exist := groups[key]; !exist {
			groups[key] = []string{}
		}
		groups[key] = append(groups[key], str)
	}

	ans := [][]string{}
	for _, group := range groups {
		ans = append(ans, group)
	}

	return ans
}

func hash(word string) string {
	charArray := []rune(word)
	sort.Slice(charArray, func(i int, j int) bool {
		return charArray[i] < charArray[j]
	})
	return string(charArray)
}
