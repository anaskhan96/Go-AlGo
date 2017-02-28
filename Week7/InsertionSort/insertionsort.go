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
	sorted := insertionSort(arr)
	fmt.Println(sorted)
}

func insertionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		v := a[i]
		j := i - 1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v
	}
	return a
}
