package isogram

import "strings"

func IsIsogram(word string) bool {
	wordMap := make(map[rune]int, len(word))
	word = strings.ToLower(word)
	for _, w := range word {
		if w >= 97 && w <= 122 {
			if _, ok := wordMap[w]; !ok {
				wordMap[w] = 1
			} else {
				return false
			}
		}
	}
	return true
}
