package main

import (
	"fmt"
	"sort"
)

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
*/

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))

}

func groupAnagrams(strs []string) [][]string {
	//var result [][]string
	var newString []string

	for i := 0; i < len(strs); i++ {
		runeSlice := []rune(strs[i])

		sort.Slice(runeSlice, func(i, j int) bool {
			return runeSlice[i] < runeSlice[j]
		})
		//fmt.Println(string(runeSlice))
		newString = append(newString, string(runeSlice))
		fmt.Println(newString)
	}

	fmt.Println(newString)

	return [][]string{}
}
