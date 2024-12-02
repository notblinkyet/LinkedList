package linkedlist

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) String() string {
	res := make([]int, l.Size())
	if len(res) != 0 {
		cur := l.Head
		for i := 0; i < l.Size(); i++ {
			res[i] = cur.Val
			cur = cur.Next
		}
	}
	return fmt.Sprint(res)
}

func (l *LinkedList) Add(val int) {
	if l.Head == nil {
		l.Head = &Node{Val: val, Next: nil}
	} else {
		cur := l.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = &Node{
			Val:  val,
			Next: nil}
	}
}

func (l *LinkedList) Pop() *int {
	cur := l.Head
	pred := &Node{
		Val:  0,
		Next: cur,
	}
	for cur.Next != nil {
		cur = cur.Next
		pred = pred.Next
	}
	pred.Next = nil
	return &cur.Val
}

func (l *LinkedList) Size() int {
	if l.Head == nil {
		return 0
	}
	cur := l.Head
	res := 0
	for cur != nil {
		cur = cur.Next
		res += 1
	}
	return res
}

func (l *LinkedList) DeleteFrom(pos int) {
	if pos == 0 {
		l.Head = l.Head.Next
	} else {
		cur := l.Head
		pred := &Node{
			Val:  0,
			Next: cur,
		}
		ind := 0
		for cur != nil {
			if ind == pos {
				pred.Next = cur.Next
				break
			}
			ind++
			cur = cur.Next
			pred = pred.Next
		}
	}
}

func (l *LinkedList) UpdateAt(pos, val int) {
	cur := l.Head
	ind := 0
	for cur != nil {
		if ind == pos {
			cur.Val = val
			break
		}
		ind++
		cur = cur.Next
	}
}

func New(arr []int) *LinkedList {
	res := LinkedList{Head: nil}
	if len(arr) == 0 {
		return &res
	}
	for i := 0; i < len(arr); i++ {
		res.Add(arr[i])
	}
	return &res
}
