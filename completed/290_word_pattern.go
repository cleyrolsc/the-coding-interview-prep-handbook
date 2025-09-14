package main

import (
	"fmt"
	"slices"
	"strings"
)

/*

Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s. Specifically:

Each letter in pattern maps to exactly one unique word in s.
Each unique word in s maps to exactly one letter in pattern.
No two letters map to the same word, and no two words map to the same letter.

*/

func main() {

	fmt.Println(wordPattern("aaa", "aa aa aa aa"))

}

func wordPattern(pattern string, s string) bool {
	myMap := map[int32]string{}
	var mySlice []string

	sSplitted := strings.Split(s, " ")
	pSplitted := strings.Split(pattern, "")

	if len(sSplitted) != len(pSplitted) {
		return false
	}

	for i, v := range pattern {
		if slices.Contains(mySlice, sSplitted[i]) {
			continue
		}
		if myMap[v] == "" {
			myMap[v] = sSplitted[i]
			mySlice = append(mySlice, sSplitted[i])
		} else {
			continue
		}
	}

	for i, v := range pattern {
		fmt.Println(myMap[v])
		if myMap[v] != sSplitted[i] {
			return false
		}
	}

	return true
}
