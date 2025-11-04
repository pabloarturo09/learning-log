package main

import (
	"errors"
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	sb := new(strings.Builder)
	words := new(Stack)

	for i := range s {
		ch := s[i]
		if ch == ' ' && sb.Len() > 0 {
			words.Push(sb.String())
			sb.Reset()
		} else if ch != ' ' {
			sb.WriteByte(ch)
		}
	}

	if sb.Len() > 0 {
		words.Push(sb.String())
	}

	sb.Reset()
	for words.top != nil {
		word, _ := words.Pop()
		sb.WriteString(word)

		if words.top != nil {
			sb.WriteByte(' ')
		}
	}

	return sb.String()
}

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}

type Node struct {
	value string
	next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(value string) {
	newNode := &Node{value: value}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) Pop() (string, error) {
	if s.top == nil {
		return "", errors.New("stack is empty")
	}

	value := s.top.value
	s.top = s.top.next
	return value, nil
}
