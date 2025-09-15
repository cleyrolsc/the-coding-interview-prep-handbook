package completed

/*

Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

*/

func main() {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	reverseString([]byte{'h', 'a', 'n', 'n', 'a', 'h'})
}

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}
