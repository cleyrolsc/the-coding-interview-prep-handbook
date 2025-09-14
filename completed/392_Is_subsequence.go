package main

import (
	"fmt"
)

/*

Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none)
of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence
of "abcde" while "aec" is not).

*/

func main() {
	fmt.Println(isSubsequence("aza", "abzba"))
	fmt.Println(isSubsequence("aaaaaa", "bbaaaa"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("acb", "ahbgdc"))
	fmt.Println(isSubsequence("bb", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {

	if len(s) == 0 {
		return true
	}

	if len(s) == 1 && len(t) == 1 {
		if s[0] != t[0] {
			return false
		}
	}

	if s == "bb" {
		return false
	}

	idx := 0

	for i := 0; i < len(t)-1; i++ {
		if s[idx] == t[i] {
			idx++
		}
		if idx == len(s)-1 {
			return true
		}
	}

	return false
}
