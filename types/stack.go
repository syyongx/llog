package types

import "errors"

type Stack []interface{}

// push
func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

// Pop
func (s *Stack) Pop() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("You tried to pop from an empty stack.")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

// length
func (s *Stack) Len() int {
	return len(*s)
}
