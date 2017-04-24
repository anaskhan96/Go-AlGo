/* Copyright, Anas Khan © 2017 */

package main

import "fmt"

func main() {
	var n, num int // n is the size of the adjacency matrix
	fmt.Scanf("%d", &n)
	adj := make([][]int, 0, 10) // the adjacency matrix
	row := make([]int, 0, 10)   // each row of the matrix
	// Input for the adjacency matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &num)
			row = append(row, num)
		}
		adj = append(adj, row)
		row = make([]int, 0, 10)
	}
	res := warshalls(adj)
	fmt.Println(res)
}

func warshalls(a [][]int) [][]int {
	for k := 0; k < len(a); k++ {
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a); j++ {
				a[i][j] = a[i][j] | (a[i][k] & a[k][j])
			}
		}
	}
	return a
}
