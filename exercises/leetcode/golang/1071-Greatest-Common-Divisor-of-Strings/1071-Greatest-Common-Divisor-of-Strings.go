package main

import (
	"fmt"
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	mcd := getMcd(len(str1), len(str2))
	pattern := str1[:mcd]

	if strings.Repeat(pattern, len(str1)/mcd) == str1 &&
		strings.Repeat(pattern, len(str2)/mcd) == str2 {
		return pattern
	}

	return ""
}

func main() {
	word1 := "BBCCBBCC"
	word2 := "BBCCBBCCBBCCBBCCBBCC"
	fmt.Println(gcdOfStrings(word1, word2))
}

func getMcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
