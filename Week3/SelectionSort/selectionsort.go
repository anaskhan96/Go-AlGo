/* Copyright, Anas Khan © 2017 */

package main

import "fmt"

func main() {
	var n, num int // size of the array
	fmt.Scanf("%d", &n)
	arr := make([]int, 0, 10)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &num) // each element of the array
		arr = append(arr, num)
	}
	sorted := selectionSort(arr)
	fmt.Println(sorted)
}

func selectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		small, pos := a[i], i
		for j := i + 1; j < len(a); j++ {
			if a[j] < small {
				small = a[j]
				pos = j
			}
		}
		/* Swapping */
		temp := a[i]
		a[i] = a[pos]
		a[pos] = temp
	}
	return a
}
