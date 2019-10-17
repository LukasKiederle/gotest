package stack

//Stack struct
type Stack struct {
	data []interface{}
}

// NewStack constructor
func NewStack() *Stack {
	return new(Stack)
}

// Push new Element to Stack
func (s *Stack) Push(element interface{}) {
	s.data = append(s.data, element)
}

// Pop last element
func (s *Stack) Pop() interface{} {
	if s.Size() > 0 {
		result := s.data[s.Size()-1]
		s.data = s.data[0 : s.Size()-1]
		return result
	}

	panic("Stack size is 0!!!")
}

// Size of objects in stack
func (s *Stack) Size() int {
	return len(s.data)
}

// GetAt index
func (s *Stack) GetAt(index int) interface{} {
	if s.Size() > 0 {
		return s.data[index]
	}

	panic("Stack size is 0!!!")
}
