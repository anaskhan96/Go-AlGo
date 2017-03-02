/* Copyright, Anas Khan © 2017 */

package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	n        int
	arr      []int
	x        int
	expected int
}

/* Example test pairs */
var tests = []testpair{{4, []int{2, -6, 2, -1}, 3, 5}, {4, []int{2, 0, 3, 1}, 2, 23}}

/* Test function */
func TestHorners(t *testing.T) {
	for _, pair := range tests {
		actual := horners(pair.arr, pair.x)
		if !reflect.DeepEqual(actual, pair.expected) {
			t.Error("For", pair.arr, "expected", pair.expected, "got", actual)
		}
	}
}
