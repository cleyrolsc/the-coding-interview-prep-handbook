package completed

import (
	"fmt"
	"sort"
	"strconv"
)

/*

You are given an integer array score of size n, where score[i] is the score of the ith athlete in a competition.
All the scores are guaranteed to be unique.

The athletes are placed based on their scores, where the 1st place athlete has the highest score, the 2nd place athlete
has the 2nd highest score, and so on. The placement of each athlete determines their rank:

The 1st place athlete's rank is "Gold Medal".
The 2nd place athlete's rank is "Silver Medal".
The 3rd place athlete's rank is "Bronze Medal".
For the 4th place to the nth place athlete, their rank is their placement number (i.e., the xth place athlete's rank is "x").
Return an array answer of size n where answer[i] is the rank of the ith athlete.

*/

func main() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
}

func findRelativeRanks(score []int) []string {
	var result []string
	mySlice := make([]int, len(score))
	copy(mySlice, score)

	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	myMap := make(map[int]string)

	for i, v := range score {
		if i >= 3 {
			myMap[v] = strconv.Itoa(i + 1)
		}

		if i == 0 {
			myMap[v] = "Gold Medal"
		} else if i == 1 {
			myMap[v] = "Silver Medal"
		} else if i == 2 {
			myMap[v] = "Bronze Medal"
		}
	}

	//slices.Sort(score)

	//for i, v := range mySlice{
	//
	//}

	for _, v := range mySlice {
		result = append(result, myMap[v])
	}

	return result
}
