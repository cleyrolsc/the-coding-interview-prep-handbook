package completed

import (
	"fmt"
)

/*

496. Next Greater Element I
Easy
Topics
premium lock icon
Companies
The next greater element of some element x in an array is the first greater element that is to the right of x in the same array.

You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.

For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the next greater element of nums2[j] in nums2. If there is no next greater element, then the answer for this query is -1.

Return an array ans of length nums1.length such that ans[i] is the next greater element as described above.

*/

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
	fmt.Println(nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)

	for _, v := range nums1 {
		for j, k := range nums2 {
			if k == v {

				if j == (len(nums2) - 1) {
					ans = append(ans, -1)
					continue
				}

				for x := j + 1; x < len(nums2); x++ {
					fmt.Println(x)
					if nums2[x] > k {
						ans = append(ans, nums2[x])
						break
					}

					if x == (len(nums2) - 1) {
						ans = append(ans, -1)
					}
				}
			}
		}
	}

	return ans
}
