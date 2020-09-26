package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(mergeSort(nums))
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[mid:])
	right := mergeSort(arr[:mid])

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

// func mergeSort(a []int, fst, lst int) {
//   var i, j, k, p, middle int
//   var w = make([]int, len(a)/2+1)
//   if fst < lst {
//     middle = (fst + lst) / 2
//     mergeSort(a, fst, middle)
//     mergeSort(a, middle+1, lst)
//     p = 0
//     for i = fst; i <= middle; i++ {
//       w[p] = a[i]
//       p++
//     }
//     i = middle + 1
//     j = 0
//     k = fst
//     for i <= lst && j < p {
//       if w[j] <= a[i] {
//         a[k] = w[j]
//         k++
//         j++
//       } else {
//         a[k] = a[i]
//         k++
//         i++
//       }
//     }
//     for j < p {
//       a[k] = w[j]
//       k++
//       j++
//     }
//   }
//   fmt.Println(a)
// }
