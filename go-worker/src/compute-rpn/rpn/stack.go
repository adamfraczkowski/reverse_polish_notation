package rpn

import (
	"fmt"
	"strconv"
	"strings"
)

// RPN calculation engine struct
type RPN struct {
	//stack storage
	stack []int
}

// Pop from stack
func (m *RPN) Pop() (int, error) {
	if len(m.stack) <= 0 {
		return -1, fmt.Errorf("Stack empty")
	}
	popValue := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	return popValue, nil
}

// Push to stack
func (m *RPN) Push(element int) int {
	m.stack = append(m.stack, element)
	return element
}

// GetStack Get actual stack
func (m *RPN) GetStack() []int {
	return m.stack
}

// StackLen get actual elements count in stack
func (m *RPN) StackLen() int {
	return len(m.stack)
}

// CalcRPN calculate RPN expression
func (m *RPN) CalcRPN(expression string) (int, error) {
	// split string into slice
	expressionSlice := strings.Split(expression, " ")
	for i := 0; i < len(expressionSlice); i++ {
		switch expressionSlice[i] {
		case "*":
			valA, _ := m.Pop()
			valB, _ := m.Pop()
			result := valB * valA
			m.Push(result)
			break

		case "+":
			valA, _ := m.Pop()
			valB, _ := m.Pop()
			result := valB + valA
			m.Push(result)
			break

		case "-":
			valA, _ := m.Pop()
			valB, _ := m.Pop()
			result := valB - valA
			m.Push(result)
			break

		case "/":
			valA, _ := m.Pop()
			valB, _ := m.Pop()
			result := valB / valA
			m.Push(result)
			break

		default:
			// simple validation of phrase
			_, err := strconv.Atoi(expressionSlice[i])
			if err != nil {
				return -1, fmt.Errorf("Incorrect phrase in expression %s", expressionSlice[i])
			}
			//add value to stack
			val, _ := strconv.Atoi(expressionSlice[i])
			m.Push(val)
			break
		}
	}
	finalResult, _ := m.Pop()
	return finalResult, nil
}
