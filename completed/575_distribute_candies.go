package completed

import "fmt"

/*

Alice has n candies, where the ith candy is of type candyType[i]. Alice noticed that she started to gain weight,
so she visited a doctor.

The doctor advised Alice to only eat n / 2 of the candies she has (n is always even). Alice likes her candies very much,
and she wants to eat the maximum number of different types of candies while still following the doctor's advice.

Given the integer array candyType of length n, return the maximum number of different types of candies she can eat if she
only eats n / 2 of them.


*/

func main() {
	//fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	//fmt.Println(distributeCandies([]int{1, 1, 2, 3}))
	//fmt.Println(distributeCandies([]int{6, 6, 6, 6}))
	fmt.Println(distributeCandies([]int{1000, 1, 1, 1}))
}

func distributeCandies(candyType []int) int {
	types := make(map[int]int)
	//count := 0

	for i := 0; i < len(candyType); i++ {
		if types[candyType[i]] == 0 {
			types[candyType[i]]++
		}
	}

	if len(types) == 1 {
		return 1
	}

	if len(types) < (len(candyType) / 2) {
		return len(types)
	}

	return len(candyType) / 2
}
