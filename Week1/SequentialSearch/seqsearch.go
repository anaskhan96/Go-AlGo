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
	index := seq_search(arr, key)
	fmt.Println(index)
}

func seq_search(a []int, k int) int {
	n, i := len(a), 0
	a = append(a, k)
	for a[i] != k {
		i++
	}
	if i < n {
		return i
	}
	/* Returning -1 if not found */
	return -1
}
