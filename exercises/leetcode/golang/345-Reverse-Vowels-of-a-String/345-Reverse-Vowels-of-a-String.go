package main

// Given a string s, reverse only all the vowels in the string and return it.

// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

// Example 1:

// Input: s = "IceCreAm"

// Output: "AceCreIm"

// Explanation:

// The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".

// Example 2:

// Input: s = "leetcode"

// Output: "leotcede"

// Constraints:

// 1 <= s.length <= 3 * 105
// s consist of printable ASCII characters.

import (
	"fmt"
)

func main() {
	s := "IceCreAm"
	fmt.Println(reverseVowels(s))
}

// SOLUTION WITH TWO POINTERS

func reverseVowels(s string) string {
	characters := []byte(s)
	left := 0
	right := len(s) - 1

	for left < right {
		leftChar := s[left]
		rightChar := s[right]

		if isVowel(leftChar) && isVowel(rightChar) {
			characters[left], characters[right] = characters[right], characters[left]
			left++
			right--
			continue
		}

		if !isVowel(leftChar) {
			left++
		}
		if !isVowel(rightChar) {
			right--
		}
	}

	return string(characters)
}

// SOLUTION WITH STACK

// func reverseVowels(s string) string {
// 	var sb = new(strings.Builder)
// 	var stack = new(Stack)

// 	for i := range s {
// 		if isVowel(s[i]) {
// 			stack.Push(s[i])
// 		}
// 	}

// 	for i := range s {
// 		if isVowel(s[i]) {
// 			vowel, _ := stack.Pop()
// 			sb.WriteByte(vowel)
// 		} else {
// 			sb.WriteByte(s[i])
// 		}
// 	}

// 	return sb.String()
// }

func isVowel(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

// type Node struct {
// 	value byte
// 	next  *Node
// }

// type Stack struct {
// 	top *Node
// }

// func (s *Stack) Push(value byte) {
// 	newNode := &Node{value: value}
// 	newNode.next = s.top
// 	s.top = newNode
// }

// func (s *Stack) Pop() (byte, error) {
// 	if s.top == nil {
// 		return 0, errors.New("stack is empty")
// 	}

// 	value := s.top.value
// 	s.top = s.top.next
// 	return value, nil
// }
