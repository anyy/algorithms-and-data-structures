package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	data []int
}

func (stack *Queue) Enqueue(i int) {
	stack.data = append(stack.data, i)
}

func (stack *Queue) Dequeue() (int, error) {
	if len(stack.data) <= 0 {
		return 0, errors.New("queue is empty")
	}
	n := stack.data[0]
	stack.data = stack.data[1:]
	return n, nil
}

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	n, err := queue.Dequeue()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n)

	n, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n)

	n, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n)
}
