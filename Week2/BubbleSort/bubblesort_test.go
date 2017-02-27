/* Copyright, Anas Khan © 2017 */

package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	arr      []int
	expected []int
}

/* Example test pairs */
var tests = []testpair{{[]int{11, 17, 5, 3, 43, 22, 90, 65, 45, 32, 14, 89, 86, 75, 47}, []int{3, 5, 11, 14, 17, 22, 32, 43, 45, 47, 65, 75, 86, 89, 90}},
	{[]int{123, -111, 32, 1234, 32}, []int{-111, 32, 32, 123, 1234}},
	{[]int{45, 43, 31, 67, 87, 54, 56, 23, 90, 134, 2, 5}, []int{2, 5, 23, 31, 43, 45, 54, 56, 67, 87, 90, 134}}}

/* Test Function */
func TestBubbleSort(t *testing.T) {
	for _, pair := range tests {
		actual := bubbleSort(pair.arr)
		if !reflect.DeepEqual(actual, pair.expected) {
			t.Error("For", pair.arr, "expected", pair.expected, "got", actual)
		}
	}
}
