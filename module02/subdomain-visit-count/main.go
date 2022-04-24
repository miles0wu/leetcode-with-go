package main

import (
	"log"
	"strconv"
)

func subdomainVisits(cpdomains []string) []string {
	counts := map[string]int{}

	for _, cpdomain := range cpdomains {
		strCount, domains := parseCpdomain(cpdomain)

		count, err := strconv.Atoi(strCount)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, domain := range domains {
			if _, ok := counts[domain]; !ok {
				counts[domain] = 0
			}
			counts[domain] += count
		}
	}

	ans := []string{}
	for doamin, count := range counts {
		ans = append(ans, strconv.Itoa(count)+" "+doamin)
	}
	return ans
}

func parseCpdomain(input string) (string, []string) {
	count := ""
	domains := []string{}

	for i := len(input); i > 0; i-- {
		if input[i-1:i] == "." {
			domains = append(domains, input[i:len(input)])
		} else if input[i-1:i] == " " {
			domains = append(domains, input[i:len(input)])
			count = input[0 : i-1]
			break
		}
	}

	return count, domains
}
