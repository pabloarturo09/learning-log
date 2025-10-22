package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		word1  = "ab"
		word2  = "pqr"
		merged = mergeAlternately(word1, word2)
	)

	fmt.Println(merged)
}

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder

	minLength := min(len(word1), len(word2))
	for i := range minLength {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[i])
	}

	if len(word1) > len(word2) {
		sb.WriteString(word1[minLength:])
	} else if len(word2) > len(word1) {
		sb.WriteString(word2[minLength:])
	}

	return sb.String()
}
