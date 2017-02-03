/* Copyright, Anas Khan © 2017 */

package main

import "fmt"

func main() {
	var n, num int
	fmt.Scanf("%d", &n) // size of the array
	arr := make([]int, 0, 10)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &num) // each element of the array
		arr = append(arr, num)
	}
	var key int
	fmt.Scanf("%d", &key) // key element to be searched for
	index := bin_search(arr, key)
	fmt.Println(index)
}

func bin_search(a []int, k int) int {
	l, h := 0, len(a)-1
	for l <= h {
		m := (l + h) / 2
		if a[m] == k {
			return m
		} else if a[m] > k {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	/* Returning -1 if not found */
	return -1
}
