package main

import "fmt"

/*
Given an array arr, replace every element in that array with the greatest element among the elements to its right,
and replace the last element with -1.

After doing so, return the array.
*/

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
	fmt.Println(replaceElements([]int{400}))
}

func replaceElements(arr []int) []int {
	//var myArr []int
	var maxRight int

	if len(arr) == 1 {
		return []int{-1}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		temp := arr[i]
		if i == len(arr)-1 {
			arr[i] = -1
			maxRight = temp
			continue
		}

		arr[i] = maxRight
		if maxRight < temp {
			maxRight = temp
		}

	}

	return arr
}
