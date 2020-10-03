package main

import (
	"fmt"
	"math"
)

var heap []int

func main() {
	heap = append(heap, 27, 37, 48, 39, 71, 56, 104, 89, 66)

	i := 0
	for i == 0 {
		fmt.Println("\n1. INSERT")
		fmt.Println("2. EXTRACT MIN")
		fmt.Println("3. DISPLAY")
		fmt.Println("4. EXIT")
		var ch int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &ch)
		switch ch {
		case 1:
			fmt.Print("Enter the num to insert: ")
			var n int
			fmt.Scanf("%d\n", &n)
			insert(n)
		case 2:
			min, _ := extractMin()
			fmt.Printf("min = %+v\n", min)
		case 3:
			fmt.Printf("heap = %+v\n", heap)
		case 4:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func insert(x int) {
	heap = append(heap, x)
	minHeapify()
	// maxHeapify()
}

// parent(i) = i / 2
func minHeapify() {
	idx := len(heap) - 1
	x := heap[idx]
	parentIdx := int(math.Floor(float64((idx - 1) / 2)))
	for idx > 0 && x < heap[parentIdx] {
		heap[idx] = heap[parentIdx]
		idx = parentIdx
		parentIdx = int(math.Floor(float64((idx - 1) / 2)))
	}
	heap[idx] = x
}

func maxHeapify() {
	idx := len(heap) - 1
	x := heap[idx]
	parentIdx := int(math.Floor(float64((idx - 1) / 2)))
	for idx > 0 && x > heap[parentIdx] {
		heap[idx] = heap[parentIdx]
		idx = parentIdx
		parentIdx = int(math.Floor(float64((idx - 1) / 2)))
	}
	heap[idx] = x
}

func extractMin() (int, bool) {
	if len(heap)-1 == 0 {
		return 0, false
	}

	min := heap[0] // root
	heap[0] = heap[len(heap)-1]
	heap = heap[:len(heap)-1]
	bubbleDown(0)
	return min, true

}

func bubbleDown(index int) {
	min := index
	leftIndex := left(index)

	for i := 0; i < 2; i++ {
		if (leftIndex + i) <= len(heap)-1 {
			if heap[min] > heap[leftIndex+i] {
				min = leftIndex + i
			}
		}
	}

	if min != index {
		heap[index], heap[min] = heap[min], heap[index]
		bubbleDown(min)
	}
}

func left(k int) int {
	return 2*k + 1
}
