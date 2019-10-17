package stack

import (
	"testing"

	"fmt"

	"github.com/LukasKiederle/gotest/src/exercises/2.1/rational"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	element1 := 10
	element2 := "Test with string"

	if stack.Size() != 0 {
		t.Error("A new generated Stack should have size 0")
	}

	stack.Push(element1)

	if stack.Size() != 1 {
		t.Error("Should have size 1")
	}

	stack.Push(element2)

	if stack.Size() != 2 {
		t.Error("Should have size 2")
	}

	returnElement1 := stack.Pop()
	returnElement2 := stack.Pop()

	if returnElement1 != element2 {
		t.Error("Should be the string object")
	}

	if returnElement2 != element1 {
		t.Error("Should be the int object")
	}

	r1 := rational.NewRational(1, 2)
	r2 := rational.NewRational(2, 4)

	stack.Push(r1)
	stack.Push(r2)

	if stack.Pop() != r2 {
		t.Error("Pop() did not return r2")
	}

}

func TestStackWithRationals(t *testing.T) {
	s := NewStack()

	r1 := rational.NewRational(1, 2)
	r2 := rational.NewRational(2, 4)
	r3 := rational.NewRational(2, 4)

	resultRational := rational.NewRational(5, 2)
	s.Push(r1)
	s.Push(r2)
	s.Push(r3)

	sum := rational.NewRational(1, 1)

	for i := 0; i < s.Size(); i++ {
		sum = sum.Add(s.GetAt(i).(rational.Rational))
	}

	if sum != resultRational {
		fmt.Printf("Result should be: %v but was %v", resultRational.String(), sum.String())
		t.Error("Result not matching!!!")
	}

}
