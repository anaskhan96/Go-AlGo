/* Copyright, Anas Khan © 2017 */

package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	n        int64
	a        int64
	b        int64
	expected int64
}

/* Example test pairs */
var tests = []testpair{{3, 456, 231, 105336}, {8, 12345678, 32165487, 397104745215186}, {6, 123678, 745921, 92254017438}}

/* Test Function */
func TestKaratsuba(t *testing.T) {
	for _, pair := range tests {
		actual := karatsuba(pair.a, pair.b, pair.n)
		if !reflect.DeepEqual(actual, pair.expected) {
			t.Error("For", pair.a, pair.b, pair.n, "expected", pair.expected, "got", actual)
		}
	}
}
