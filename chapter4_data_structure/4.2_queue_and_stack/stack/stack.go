package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []int
}

func (stack *Stack) Push(i int) {
	stack.data = append(stack.data, i)
}

func (stack *Stack) Pop() (int, error) {
	if len(stack.data) <= 0 {
		return 0, errors.New("stack is empty")
	}
	n := stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	return n, nil
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	n, err := stack.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n)

	n, err = stack.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n)

	n, err = stack.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n)
}
