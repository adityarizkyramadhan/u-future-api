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

func CompareArray(arr1, arr2 []string) int {
	count := 0
	for _, val1 := range arr1 {
		// loop array kedua
		for _, val2 := range arr2 {
			// jika ditemukan kesamaan data
			if val1 == val2 {
				// tambahkan nilai count
				count++
				// keluar dari loop array kedua
				break
			}
		}
	}
	return count
}
