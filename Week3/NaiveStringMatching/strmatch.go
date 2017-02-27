/* Copyright, Anas Khan © 2017 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text_, _ := bufio.NewReader(os.Stdin).ReadString('\n')            // input text
	pattern_, _ := bufio.NewReader(os.Stdin).ReadString('\n')         // string to search for
	text, pattern := text_[:len(text_)-1], pattern_[:len(pattern_)-1] // removing '\n' from both strings
	index := stringMatching(text, pattern)
	fmt.Println(index)
}

func stringMatching(t, p string) int {
outer:
	for i := 0; i < len(t)-len(p)+1; i++ {
		k := i
		for j := 0; j < len(p); j++ {
			if t[k] != p[j] {
				continue outer
			}
			k++
		}
		/* Returning index if matched */
		return i
	}
	/* Returning -1 if not found */
	return -1
}
