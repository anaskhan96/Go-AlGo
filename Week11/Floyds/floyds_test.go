/* Copyright, Anas Khan © 2017 */

package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	adj      [][]int
	expected [][]int
}

/* Example test pairs */
var tests = []testpair{{[][]int{[]int{0, 1000, 3, 1000}, []int{2, 0, 1000, 1000}, []int{1000, 7, 0, 1}, []int{6, 1000, 1000, 0}},
	[][]int{[]int{0, 10, 3, 4}, []int{2, 0, 5, 6}, []int{7, 7, 0, 1}, []int{6, 16, 9, 0}}}}

/* Test Function */
func TestFloyds(t *testing.T) {
	for _, pair := range tests {
		actual := floyds(pair.adj)
		if !reflect.DeepEqual(actual, pair.expected) {
			t.Error("For", pair.adj, "expected", pair.expected, "got", actual)
		}
	}
}
