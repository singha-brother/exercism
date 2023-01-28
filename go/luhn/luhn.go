package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	var values []int
	for _, char := range id {
		i, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		values = append(values, i)
	}

	sum := 0
	values = reverseSlice(values)
	for i, v := range values {
		if i%2 != 0 {
			v *= 2
			if v > 9 {
				v -= 9
			}
			sum += v
		} else {
			sum += v
		}
	}
	return sum%10 == 0
}

// as I read the description carelessly,
// I thought that the order is from left to right
// and implemented that order
// so quick solution is to reverse the slice
func reverseSlice(slice []int) []int {
	revSlice := []int{}
	for i := range slice {
		revSlice = append(revSlice, slice[len(slice)-1-i])
	}
	return revSlice
}
