/* Copyright, Anas Khan © 2017 */

package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	n        uint
	m        uint
	expected uint
}

/* Example test pairs */
var tests = []testpair{{3, 4, 12}, {12, 80, 960}, {845, 93781, 79244945}, {12345678, 32165487, 397104745215186}, {123678, 745921, 92254017438}}

/* Test function */
func TestRussianPeasant(t *testing.T) {
	for _, pair := range tests {
		actual := russian_peasant(pair.n, pair.m)
		if !reflect.DeepEqual(actual, pair.expected) {
			t.Error("For", pair.n, pair.m, "expected", pair.expected, "got", actual)
		}
	}
}
