package main

import (
	"errors"
	"fmt"
)

var errEmpty = errors.New("empty error")

type MyStack struct {
	list []int
}

func (m *MyStack) Pop() (int, error) {
	if m.list == nil || len(m.list) == 0 {
		return 0, errEmpty
	}

	result := m.list[len(m.list)-1]
	m.list = m.list[:len(m.list)-1]
	return result, nil
}

func (m *MyStack) Push(value int) {
	if m.list == nil {
		m.list = make([]int, 0)
	}
	m.list = append(m.list, value)
}

func (m *MyStack) Len() int {
	return len(m.list)
}

func (m *MyStack) Peek() (int, error) {
	if m.list == nil || len(m.list) == 0 {
		return 0, errEmpty
	}

	return m.list[0], nil
}

type MyQueue struct {
	in  *MyStack
	out *MyStack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	queue := MyQueue{
		in:  &MyStack{},
		out: &MyStack{},
	}
	return queue
}

/** Push element x to the back of queue. */
func (m *MyQueue) Push(x int) {
	m.in.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (m *MyQueue) Pop() int {
	if m.out.Len() == 0 {
		for m.in.Len() != 0 {
			value, err := m.in.Pop()
			if err != nil {
				panic(err)
			}
			m.out.Push(value)
		}
	}

	result, err := m.out.Pop()
	if err != nil {
		return 0
	}
	return result
}

/** Get the front element. */
func (m *MyQueue) Peek() int {
	if m.in.Len() == 0 && m.out.Len() == 0 {
		return 0
	}

	if m.in.Len() != 0 {
		result, err := m.in.Peek()
		if err != nil {
			result, err = m.out.Peek()
			if err != nil {
				return 0
			}
			return result
		}
		return result
	}
	return 0
}

/** Returns whether the queue is empty. */
func (m *MyQueue) Empty() bool {
	if m.in.Len() == 0 && m.out.Len() == 0 {
		return true
	}
	return false
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}
