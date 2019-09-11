package main

// RPN calculation engine struct
type RPN struct {
	//stack storage
	stack []string
}

// Pop from stack
func (m *RPN) Pop() string {
	if len(m.stack) <= 0 {
		return ""
	}
	popValue := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	return popValue
}

// Push to stack
func (m *RPN) Push(element string) string {
	m.stack = append(m.stack, element)
	return element
}

// GetStack Get actual stack
func (m *RPN) GetStack() []string {
	return m.stack
}

// StackLen get actual elements count in stack
func (m *RPN) StackLen() int {
	return len(m.stack)
}
