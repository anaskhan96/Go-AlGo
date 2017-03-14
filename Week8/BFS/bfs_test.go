/* Copyright, Anas Khan © 2017 */

package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	visit      []int
	adj        [][]int
	components int
}

/* Example test pairs */
var tests = []testpair{{[]int{0, 0, 0, 0, 0}, [][]int{[]int{0, 1, 1, 1, 0}, []int{1, 0, 0, 0, 1}, []int{1, 0, 0, 0, 1}, []int{1, 0, 0, 0, 1}, []int{0, 1, 1, 1, 0}}, 1},
	{[]int{0, 0, 0, 0, 0}, [][]int{[]int{0, 1, 1, 1, 0}, []int{1, 0, 0, 0, 0}, []int{1, 0, 0, 0, 0}, []int{1, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}}, 2}}

/* Test Function */
func TestBfs(t *testing.T) {
	for _, pair := range tests {
		actual := callBfs(pair.adj, pair.visit)
		if !reflect.DeepEqual(actual, pair.components) {
			t.Error("For", pair.adj, "expected", pair.components, "got", actual)
		}
	}
}
