package linkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	wait := LinkedList{
		&Node{Val: 1, Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 3,
				Next: &Node{
					Val: 4,
					Next: &Node{
						Val:  5,
						Next: nil,
					},
				},
			},
		}},
	}
	fact := New(s)
	if wait.String() != fact.String() {
		t.Errorf("\nNew Error;\nWait: %s;\nHave: %s;\n", wait.String(), fact.String())
	}

	wait = LinkedList{Head: nil}
	fact = New([]int{})
	if wait.String() != fact.String() {
		t.Errorf("\nNew Error with empty Array;\nWait: %s;\nHave: %s;\n", wait.String(), fact.String())
	}
}

func TestAdd(t *testing.T) {
	wait := New([]int{1, 2, 3, 4, 5})
	fact := New([]int{1, 2, 3, 4})
	fact.Add(5)
	if wait.String() != fact.String() {
		t.Errorf("\nAdd Error;\nWait: %s;\nHave: %s;\n", wait.String(), fact.String())
	}

	wait = New([]int{1})
	fact = New([]int{})
	fact.Add(1)
	if wait.String() != fact.String() {
		t.Errorf("\nAdd Error with empty Array;\nWait: %s;\nHave: %s;\n", wait.String(), fact.String())
	}
}

func TestPop(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	wait := New([]int{1, 2, 3, 4, 5})
	fact := New(s)
	val := fact.Pop()
	if val != s[len(s)-1] {
		t.Errorf("\nPop Error with value;\nWait val: %d\nHave val: %d\n", s[len(s)-1], val)
	}
	if wait.String() != fact.String() {
		t.Errorf("Pop Error, wrong result List:\n Wait %s;\n Have %s;\n", wait.String(), fact.String())
	}
}

func TestSize(t *testing.T) {
	s := []int{1, 2, 3}
	res := New(s)
	if res.Size() != len(s) {
		t.Errorf("\nSize Error;\nWait: %d;\nHave: %d;\n", len(s), res.Size())
	}

	s = []int{}
	res = New(s)
	if len(s) != res.Size() {
		t.Errorf("\nSize Error with empty arr;\nWait: %d;\nHave: %d;\n", len(s), res.Size())
	}
}

func TestDeleteFrom(t *testing.T) {
	wait := New([]int{1, 2, 4, 5, 6})
	fact := New([]int{1, 2, 3, 4, 5, 6})
	fact.DeleteFrom(2)
	if wait.String() != fact.String() {
		t.Errorf("\nDeleteFrom Error;\nWait: %s;\nHave: %s;\n", wait.String(), fact.String())
	}
	wait = New([]int{1, 2, 3})
	fact = New([]int{1, 2, 3})
	fact.DeleteFrom(10)
	fact.DeleteFrom(3)
	if wait.String() != fact.String() {
		t.Errorf("\nDeleteFrom Error with out off list index;\nWait: %s;\nHave: %s;\n", wait.String(), fact.String())
	}
}

func TestUpdateAt(t *testing.T) {
	wait := New([]int{1, 2, 3, 4, 5})
	fact := New([]int{1, 2, 4, 4, 5})
	fact.UpdateAt(2, 3)
	if wait.String() != fact.String() {
		t.Errorf("\nUpdateAt Error;\nWait: %s;\nHave: %s;\n", wait.String(), fact.String())
	}
	wait = New([]int{1, 2, 3, 4, 5})
	fact = New([]int{1, 2, 3, 4, 5})
	fact.UpdateAt(5, 10)
	fact.UpdateAt(100, 100)
	if wait.String() != fact.String() {
		t.Errorf("\nUpdateAt Error with out off list index;\nWait: %s;\nHave: %s;\n", wait.String(), fact.String())
	}
}
