/* Copyright, Anas Khan © 2017 */

package main

import "fmt"

func main() {
	var n, components, num int // size and number of compnents
	fmt.Scanf("%d", &n)
	visit := make([]int, 0, 10)
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
	// Initializing the visit matrix to zero
	for i := 0; i < n; i++ {
		visit = append(visit, 0)
	}
	// Calling dfs() on all functions
	for i := 0; i < n; i++ {
		if visit[i] == 0 {
			components++
			dfs(adj, i, visit)
		}
	}
	fmt.Println()
	// Printing the number of components
	fmt.Println(components)
}

func dfs(a [][]int, v int, visit []int) {
	visit[v] = 1
	fmt.Printf("%d->", v) // dfs traversal
	for w := 0; w < len(a); w++ {
		if a[v][w] == 1 && visit[w] == 0 {
			dfs(a, w, visit)
		}
	}
}
