package main

import (
	"fmt"
)

/*

Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that
they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1
< index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.

*/

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		currentSum := nums[left] + nums[right]

		if currentSum == target {
			return []int{left + 1, right + 1}
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}
