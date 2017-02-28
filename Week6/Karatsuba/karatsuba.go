/* Copyright, Anas Khan © 2017 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b int64
	fmt.Scanf("%d", &n) // no. of digits
	fmt.Scanf("%d", &a) // first number
	fmt.Scanf("%d", &b) // second number
	res := karatsuba(a, b, n)
	fmt.Println(res)
}

func karatsuba(a, b int64, n int64) int64 {
	a_str, b_str := strconv.FormatInt(a, 10), strconv.FormatInt(b, 10)
	if int64(len(a_str)) > n || int64(len(b_str)) > n {
		n++
	}
	if n == 1 {
		return a * b
	}
	m := n / 2
	m_ := m
	m2 := n - m
	if len(a_str) > len(b_str) {
		m_--
	} else if len(a_str) < len(b_str) {
		m--
	}
	a1, _ := strconv.ParseInt(a_str[:m], 10, 64)
	a2, _ := strconv.ParseInt(a_str[m:], 10, 64)
	b1, _ := strconv.ParseInt(b_str[:m_], 10, 64)
	b2, _ := strconv.ParseInt(b_str[m_:], 10, 64)
	p0 := karatsuba(a2, b2, m2)
	p1 := karatsuba(a1+a2, b1+b2, m)
	p2 := karatsuba(a1, b1, m)
	return (p2 * pow(10, 2*m2)) + ((p1 - p2 - p0) * pow(10, m2)) + p0
}

func pow(n, m int64) int64 {
	if m == 1 {
		return n
	}
	return n * pow(n, m-1)
}
