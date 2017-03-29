/* Copyright, Anas Khan © 2017 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, num, x int // n terms for degree n-1, x for value of x
	fmt.Scanf("%d", &n)
	arr := make([]int, 0, 10)
	for i := n - 1; i >= 0; i-- {
		fmt.Scanf("%d", &num) // each coefficient of x starting with x=n-1
		arr = append(arr, num)
	}
	fmt.Scanf("%d", &x)
	res := horners(arr, x)
	fmt.Println(res)
	// Calculate power using horners method
	var a, p int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &p)
	res = hornersPower(a, p) // a^p
	fmt.Println(res)
}

func horners(a []int, x int) int {
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = result*x + a[i]
	}
	return result
}

func hornersPower(a, n int) int {
	result := 1
	p := strconv.FormatInt(int64(n), 2) // converting it to binary
	for _, digit := range p {
		result = result * result
		if digit == '1' {
			result = result * a
		}
	}
	return result
}
