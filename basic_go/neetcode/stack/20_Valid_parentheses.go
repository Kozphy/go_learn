package stack

import (
	"errors"
	"fmt"
)

type Stack interface {
	Push(interface{})
	Pop() (interface{}, error)
	isEmpty() bool
}

type stack struct {
	data []interface{}
}

func (s *stack) Push(data interface{}) {
	s.data = append(s.data, data)
}

// remove and return top elem
func (s *stack) Pop() (interface{}, error) {
	if s.isEmpty() {
		return nil, errors.New("stack is empty, nothing could delete")
	}
	last_elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last_elem, nil
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func NewStack() Stack {
	stack := &stack{}
	return stack
}

func isValid(s string) bool {
	stack := NewStack()

	for _, str := range s {
		if str == '(' || str == '[' || str == '{' {
			stack.Push(str)
		} else {
			if stack.isEmpty() {
				return false
			}
			char, err := stack.Pop()
			if err != nil {
				return false
			}
			if (char == '(' && str == ')') || (char == '[' && str == ']') || (char == '{' && str == '}') {
				continue
			} else {
				return false
			}
		}
	}
	return stack.isEmpty()
}

func Execute_isValid() {
	s := "()[]{}"
	ans := isValid(s)
	fmt.Println(ans)
}
