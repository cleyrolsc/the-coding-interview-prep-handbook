package completed

import "fmt"

/*

You are given an array of characters letters that is sorted in non-decreasing order, and a character target.
There are at least two different characters in letters.

Return the smallest character in letters that is lexicographically greater than target. If such a character
does not exist, return the first character in letters.
*/

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
	fmt.Println(nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z'))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	smallest := letters[0]

	for _, v := range letters {

		if smallest <= target {
			if v > target {
				smallest = v
			}
		} else if v < smallest && v > target {
			smallest = v
		}
	}

	//fmt.Printf("%c", smallest)
	return smallest
}
