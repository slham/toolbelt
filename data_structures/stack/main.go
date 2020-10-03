package main

import "log"

type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(in int) {
	s.items = append(s.items, in)
}

// Pop
func (s *Stack) Pop() int {
	loc := len(s.items) - 1
	out := s.items[loc]
	s.items = s.items[:loc]
	return out
}

func main() {
	s := Stack{}
	log.Println(s)
	s.Push(1)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	log.Println(s)
	out := s.Pop()
	log.Println(s, out)
}
