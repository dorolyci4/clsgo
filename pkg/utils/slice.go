package utils

import (
	"sort"
	"strings"
)

func TailToHead(a []any) []any {
	prefix := a[len(a)-1]
	for i := len(a) - 1; i >= 1; i-- {
		a[i] = a[i-1]
	}
	a[0] = prefix
	return a
}

func CaseFoldIn(target string, str_array []string) bool {
	for _, element := range str_array {
		if strings.EqualFold(target, element) {
			return true
		}
	}
	return false
}

func In(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)

	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}
