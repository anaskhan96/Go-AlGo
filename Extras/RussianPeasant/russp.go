/* Copyright, Anas Khan © 2017 */

package main

import "fmt"

func main() {
	var n, m uint
	fmt.Scanf("%d %d", &n, &m) // the numbers to be multiplied
	res := russian_peasant(n, m)
	fmt.Println(res)
}

func russian_peasant(n, m uint) uint {
	if n == 1 {
		return m
	}
	if n&1 == 1 {
		return m + russian_peasant(n>>1, m<<1)
	}
	return russian_peasant(n>>1, m<<1)
}
