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
	sorted := quickSort(arr)
	fmt.Println(sorted)
}

func quickSort(a []int) []int {
	if len(a) > 1 {
		s := partition(a)
		quickSort(a[:s])
		quickSort(a[s+1:])
	}
	return a
}

func partition(a []int) int {
	pivot := a[0]
	i, j := 1, len(a)-1
	for i <= j {
		for i <= j && a[i] <= pivot {
			i++
		}
		for i <= j && a[j] >= pivot {
			j--
		}
		if i < j {
			/* Swapping when i and j come across elements
			which belong on opposite sides */
			temp := a[i]
			a[i] = a[j]
			a[j] = temp
		}
	}
	/* Moving the pivot element to its right position */
	t := a[0]
	a[0] = a[j]
	a[j] = t
	return j
}
