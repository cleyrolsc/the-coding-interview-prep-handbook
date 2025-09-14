package completed

import (
	"fmt"
)

/*

You are given a string s consisting of lowercase English letters. A duplicate removal consists of
choosing two adjacent and equal letters and removing them.

We repeatedly make duplicate removals on s until we no longer can.

Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.

*/

func main() {
	fmt.Println(removeDuplicates("aababaab"))

}

func removeDuplicates(s string) string {
	var found bool
	index := 0

	for {

		if len(s) <= 1 {
			return s
		}

		//fmt.Println("Entre al index: ", index)
		//fmt.Printf("%c, %c\n", s[index], s[index+1])
		if s[index] == s[index+1] {

			if len(s) == 2 {
				return ""
			}

			firstPart := s[:index]
			secondPart := s[index+2:]

			s = firstPart + secondPart

			index = 0
			found = true
			continue
		}

		if len(s) == 2 {
			return s
		}

		index++
		found = false

		if index == (len(s)-1) && !found {
			break
		}

	}

	return s
}
