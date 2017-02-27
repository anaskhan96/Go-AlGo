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
	sorted := mergeSort(arr)
	fmt.Println(sorted)
}

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2
	/* Dividing */
	x1 := mergeSort(a[:mid])
	x2 := mergeSort(a[mid:])
	/* Conquering and returning */
	return merge(x1, x2)
}

func merge(x, y []int) []int {
	z := make([]int, len(x)+len(y))
	var i, j int
	for i < len(x) && j < len(y) {
		if x[i] < y[j] {
			z[i+j] = x[i]
			i++
		} else {
			z[i+j] = y[j]
			j++
		}
	}
	/* When either of the passed arrays' traversal finishes */
	for i < len(x) {
		z[i+j] = x[i]
		i++
	}
	for j < len(y) {
		z[i+j] = y[j]
		j++
	}
	return z
}
