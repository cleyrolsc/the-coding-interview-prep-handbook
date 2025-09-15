package main

import "fmt"

/*

Given a string s and an integer k, reverse the first k characters for every 2k characters counting from the start
of the string.

If there are fewer than k characters left, reverse all of them. If there are less than 2k but greater than or equal
to k characters, then reverse the first k characters and leave the other as original.

*/

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
	//fmt.Println(reverseStr("abcd", 2))
}

func reverseStr(s string, k int) string {

	arrBytes := []byte(s)
	var result []byte

	if len(arrBytes) < k {
		newArrBytes := reverse(arrBytes)
		return string(newArrBytes)
	}

	counter := k

	for i := 0; i < len(arrBytes); i += k * 2 {
		toSlice := arrBytes[i:counter]
		fmt.Printf("%c\n", toSlice)
		toSlice = reverse(toSlice)

		//fmt.Println(i)
		result = append(result, toSlice...)
		if len(arrBytes[i:]) >= k*2 {
			restSlice := arrBytes[len(toSlice) : len(toSlice)+k]
			result = append(result, restSlice...)
		} else {

		}
		//result = append(result, restSlice...)
		fmt.Printf("%c", result)
		counter += k
	}

	return ""
}

func reverse(b []byte) []byte {
	left := 0
	right := len(b) - 1

	for left < right {
		temp := b[left]
		b[left] = b[right]
		b[right] = temp
		left++
		right--
	}

	return b
}
