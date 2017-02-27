/* Copyright, Anas Khan © 2017 */

package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	arr      []int
	key      int
	expected int
}

/* Example test pairs */
var tests = []testpair{{[]int{11, 17, 5, 3, 43, 22, 90, 65, 45, 32, 14, 89, 86, 75, 47}, 65, 7},
	{[]int{123, -111, 32, 1234, 32}, 32, 2}, {[]int{45, 43, 31, 67, 87, 54, 56, 23, 90, 134, 2, 5}, 85, -1}}

/* Test function */
func TestSeqSearch(t *testing.T) {
	for _, pair := range tests {
		actual := seqSearch(pair.arr, pair.key)
		if !reflect.DeepEqual(actual, pair.expected) {
			t.Error("For", pair.arr, "expected", pair.expected, "got", actual)
		}
	}
}
