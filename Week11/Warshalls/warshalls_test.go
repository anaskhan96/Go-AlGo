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
var tests = []testpair{{[][]int{[]int{0, 1, 0, 0, 0, 0}, []int{0, 0, 1, 0, 0, 0}, []int{1, 0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 1, 0}, []int{0, 0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 1, 0}},
	[][]int{[]int{1, 1, 1, 1, 1, 0}, []int{1, 1, 1, 1, 1, 0}, []int{1, 1, 1, 1, 1, 0}, []int{0, 0, 0, 1, 1, 0}, []int{0, 0, 0, 1, 1, 0}, []int{0, 0, 0, 1, 1, 0}}}}

/* Test Function */
func TestWarshalls(t *testing.T) {
	for _, pair := range tests {
		actual := warshalls(pair.adj)
		if !reflect.DeepEqual(actual, pair.expected) {
			t.Error("For", pair.adj, "expected", pair.expected, "got", actual)
		}
	}
}
