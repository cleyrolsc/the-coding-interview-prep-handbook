package main

import (
	"fmt"
	"sort"
	"strings"
)

/*

Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

You can use each character in text at most once. Return the maximum number of instances that can be formed.

*/

func main() {

	fmt.Println(maxNumberOfBalloons("nllbblooon"))

}

func maxNumberOfBalloons(text string) int {
	var myMap = map[byte]int{}
	myBytes := []byte(text)

	sort.Slice(myBytes, func(i, j int) bool {
		return myBytes[i] < myBytes[j] // Simple byte comparison
	})

	if strings.Contains(text, "a") == false {
		return 0
	}

	for _, v := range myBytes {
		if v == 'b' || v == 'a' || v == 'l' || v == 'o' || v == 'n' {
			myMap[v]++
		}
	}

	var double bool

	fmt.Println(myMap)
	var minimum int = myMap['b']
	var maximum int = myMap['b']

	for _, v := range myMap {
		if v < minimum {
			minimum = v
		}

		if v > maximum {
			maximum = v
		}
	}

	var useLOrO int = 0

	if myMap['l'] > myMap['o'] {
		useLOrO = myMap['o']
	} else {
		useLOrO = myMap['l']
	}

	fmt.Println(useLOrO)
	fmt.Println(minimum)

	for i := 1; i <= minimum; i++ {

		if minimum*2 == useLOrO {
			return minimum
		} else if minimum*2 < useLOrO {
			useLOrO--
		} else {
			minimum--
		}
	}

	if myMap['l'] >= (minimum*2) && myMap['o'] >= (minimum*2) {
		double = true
	}

	if double {
		return minimum
	}

	return 0

}
