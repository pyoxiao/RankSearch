package utils

import (
	"strings"
)
func GetNumber(s string, str strType) int {
	strNeed := smap[str]
	index := strings.Index(s, strNeed)
	if index == -1 {
		return -1
	} 
	number := 0
	for pos := index + len(strNeed); s[pos] >= '0' && s[pos] <= '9'; pos += 1 {
		number = number * 10 + int(s[pos] - '0')
	}
	return number
}