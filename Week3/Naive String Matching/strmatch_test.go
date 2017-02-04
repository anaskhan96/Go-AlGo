/* Copyright, Anas Khan © 2017 */

package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	text    string
	pattern string
	index   int
}

/* Example test pairs */
var tests = []testpair{{"hello world", "wo", 6},
	{"The pattern is not supposed to be here", "si", -1},
	{"You gotta test your program yourself", "ur", 17}}

/* Test Function */
func TestStringMatching(t *testing.T) {
	for _, pair := range tests {
		actual := string_matching(pair.text, pair.pattern)
		if !reflect.DeepEqual(actual, pair.index) {
			t.Error("For", pair.text, "with the pattern", pair.pattern,
				"expected", pair.index, "got", actual)
		}
	}
}
