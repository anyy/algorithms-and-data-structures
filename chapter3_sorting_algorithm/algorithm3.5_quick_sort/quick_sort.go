package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(a []int, fst, lst int) {
	var i, j, pivot int
	pivot = a[(fst+lst)/2]
	i = fst
	j = lst
	for true {
		for a[i] < pivot {
			i++
		}
		for pivot < a[j] {
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	if fst < i-1 {
		quickSort(a, fst, i-1)
	}
	if j+1 < lst {
		quickSort(a, j+1, lst)
	}
}
