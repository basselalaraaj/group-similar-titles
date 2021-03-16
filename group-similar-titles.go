package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertIntSliceToStringSlice(data []int) []string {
	wordMapKey := []string{}
	for _, v := range data {
		text := strconv.Itoa(v)
		wordMapKey = append(wordMapKey, text)
	}
	return wordMapKey
}

func contains(data []string, query string) bool {
	for _, v := range data {
		if query == v {
			return true
		}
	}

	return false
}

func groupTitles(titles []string) map[string][]string {
	groupedTitles := make(map[string][]string)

	for _, title := range titles {
		wordMap := make([]int, 26)
		for _, l := range []rune(title) {
			letterIndex := l - 'a'
			wordMap[letterIndex]++
		}

		hashKey := strings.Join(convertIntSliceToStringSlice(wordMap), "#")
		groupedTitles[hashKey] = append(groupedTitles[hashKey], title)
	}

	return groupedTitles
}

func main() {
	titles := []string{"duel", "dule", "speed", "spede", "deul", "cars"}
	query := "spede"

	for _, v := range groupTitles(titles) {
		if contains(v, query) {
			fmt.Println(v)
		}
	}

}
