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
	sorted := bubble_sort(arr)
	fmt.Println(sorted)
}

func bubble_sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				/* Swapping the XOR way */
				a[j] = a[j] ^ a[j+1]
				a[j+1] = a[j] ^ a[j+1]
				a[j] = a[j] ^ a[j+1]
			}
		}
	}
	return a
}
