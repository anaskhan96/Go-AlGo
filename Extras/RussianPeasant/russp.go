/* Copyright, Anas Khan © 2017 */

package main

import "fmt"

func main() {
	var n, m uint              // the algorithm being implemented works on positive integers, so unsigned ints
	fmt.Scanf("%d %d", &n, &m) // the numbers to be multiplied
	res := russianPeasant(n, m)
	fmt.Println(res)
}

func russianPeasant(n, m uint) uint {
	if n == 1 {
		return m
	}
	if n&1 == 1 {
		return m + russianPeasant(n>>1, m<<1)
	}
	return russianPeasant(n>>1, m<<1)
}
