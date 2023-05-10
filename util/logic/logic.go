package logic

import (
	"sort"
	"strings"
)

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

func CompareJurusanAndHistory(jurusanRAISEC, studentRAISEC string) int {
	arrJurusan := TrimAndChangeStringToArray(jurusanRAISEC)
	arrStudent := TrimAndChangeStringToArray(studentRAISEC)
	return CompareArray(arrJurusan, arrStudent)
}

func CalculateResult(data string, jawaban []string) float64 {
	totalR, totalI, totalA, totalS, totalE, totalC := 0, 0, 0, 0, 0, 0

	for _, jawab := range jawaban {
		switch jawab {
		case "R":
			totalR++
		case "I":
			totalI++
		case "A":
			totalA++
		case "S":
			totalS++
		case "E":
			totalE++
		case "C":
			totalC++
		}
	}

	switch data {
	case "r":
		return float64(totalR) / 4 * 100
	case "I":
		return float64(totalI) / 4 * 100
	case "A":
		return float64(totalA) / 3 * 100
	case "S":
		return float64(totalS) / 3 * 100
	case "E":
		return float64(totalE) / 3 * 100
	case "C":
		return float64(totalC) / 3 * 100
	default:
		return 0
	}
}

func GetMostFrequentItems(dataArray []string) string {
	frequency := make(map[string]int)
	for _, item := range dataArray {
		frequency[item]++
	}
	sortedItems := make([]string, 0, len(frequency))
	for item := range frequency {
		sortedItems = append(sortedItems, item)
	}
	sort.Slice(sortedItems, func(i, j int) bool {
		return frequency[sortedItems[i]] > frequency[sortedItems[j]]
	})
	result := make([]string, 0, 3)
	for i := 0; i < 3 && i < len(sortedItems); i++ {
		result = append(result, sortedItems[i])
	}

	str := strings.Join(result, ",")
	return str
}
