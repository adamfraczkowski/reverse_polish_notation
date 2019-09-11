package main

import "testing"

func TestStack(t *testing.T) {
	// initialize stack
	testStack := RPN{}
	// push values to stack
	_ = testStack.Push("1")
	_ = testStack.Push("2")
	_ = testStack.Push("3")
	// check length of stack
	stackLength := testStack.StackLen()
	if stackLength != 3 {
		t.Errorf("Length was incorrect, got %d, want %d", stackLength, 3)
	}
	// check stack order and GetStack method
	stackList := testStack.GetStack()
	if len(stackList) != 3 {
		t.Errorf("GetStack method failed got %d, want %d", len(stackList), 3)
	}
	if stackList[0] != "1" || stackList[1] != "2" || stackList[2] != "3" {
		t.Errorf("Incorrect order of stack")
	}
	// pop stack and check value and stack length after pop
	testValue := testStack.Pop()
	if testValue != "3" {
		t.Errorf("Incorrect value from POP got %s want %s", testValue, "3")
	}
	//checking length after pop
	actualLen := testStack.StackLen()
	if actualLen != 2 {
		t.Errorf("Incorrect stack length got %d want %d", actualLen, 2)
	}
	// pop two times to have zero length
	_ = testStack.Pop()
	_ = testStack.Pop()
	stackLen := testStack.StackLen()
	if stackLen != 0 {
		t.Errorf("Incorrect stack length got %d want %d", stackLen, 0)
	}
	actualZeroVal := testStack.Pop()
	if actualZeroVal != "" {
		t.Errorf("Incorrect stack value got %s want %s", actualZeroVal, "")
	}
}

func TestRPN(t *testing.T) {

}
