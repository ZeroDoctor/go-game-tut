package util

type Stack struct {
	data []interface{}
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Push(element interface{}) {
	s.data = append(s.data, element)
}

func (s *Stack) Pop() interface{} {
	result := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return result
}

func (s *Stack) Peek() interface{} {
	return s.data[len(s.data)-1]
}
