package rpn

import "testing"

func TestStack(t *testing.T) {
	// initialize stack
	testStack := RPN{}
	// push values to stack
	_ = testStack.Push(1)
	_ = testStack.Push(2)
	_ = testStack.Push(3)
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
	if stackList[0] != 1 || stackList[1] != 2 || stackList[2] != 3 {
		t.Errorf("Incorrect order of stack")
	}
	// pop stack and check value and stack length after pop
	testValue, _ := testStack.Pop()
	if testValue != 3 {
		t.Errorf("Incorrect value from POP got %d want %s", testValue, "3")
	}
	//checking length after pop
	actualLen := testStack.StackLen()
	if actualLen != 2 {
		t.Errorf("Incorrect stack length got %d want %d", actualLen, 2)
	}
	// pop two times to have zero length
	_, _ = testStack.Pop()
	_, _ = testStack.Pop()
	stackLen := testStack.StackLen()
	if stackLen != 0 {
		t.Errorf("Incorrect stack length got %d want %d", stackLen, 0)
	}
	actualZeroVal, err := testStack.Pop()
	if err == nil {
		t.Errorf("Incorrect stack value got %d want %s", actualZeroVal, "")
	}
	//check RPN Value
	result, err := testStack.CalcRPN("12 2 3 4 * 10 5 / + * +")
	if err != nil {
		t.Errorf("Error during calculation.%s", err)
	}
	if result != 40 {
		t.Errorf("Incorrect result value got %d want %s", result, "40")
	}

	resultSec, errSec := testStack.CalcRPN("3 4 +")
	if errSec != nil {
		t.Errorf("Error during calculation.%s", err)
	}
	if resultSec != 7 {
		t.Errorf("Incorrect result value got %d want %s", result, "40")
	}

	resultTh, errTh := testStack.CalcRPN("5 1 2 + 4 * + 3 -")
	if errTh != nil {
		t.Errorf("Error during calculation.%s", err)
	}
	if resultTh != 14 {
		t.Errorf("Incorrect result value got %d want %s", result, "14")
	}
}
