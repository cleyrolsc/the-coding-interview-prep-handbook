package completed

import "fmt"

/*
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
*/

func main() {
	fmt.Println(validMountainArray([]int{2, 1}))
	fmt.Println(validMountainArray([]int{3, 5, 5}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
	fmt.Println(validMountainArray([]int{2, 0, 2}))
}

func validMountainArray(arr []int) bool {

	if len(arr) < 3 {
		return false
	}

	decreasing := false
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			return true
		}

		if arr[i] == arr[i+1] {
			return false
		}

		if decreasing == false {
			if arr[i] > arr[i+1] {
				decreasing = true
			}
		} else {
			if arr[i] < arr[i+1] {
				return false
			}
		}

	}

	return true
}
