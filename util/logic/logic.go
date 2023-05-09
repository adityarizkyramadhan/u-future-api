package logic

import "strings"

func MostFrequentElements(arr []string) string {
	counts := make(map[string]int)
	for _, item := range arr {
		counts[item]++
	}

	maxCount := 0
	mostFrequent := ""
	for item, count := range counts {
		if count > maxCount {
			maxCount = count
			mostFrequent = item
		}
	}
	return mostFrequent
}

func TrimAndChangeStringToArray(s string) []string {
	subs := strings.Split(s, ",")
	for i, sub := range subs {
		subs[i] = strings.TrimSpace(sub)
	}
	return subs
}
