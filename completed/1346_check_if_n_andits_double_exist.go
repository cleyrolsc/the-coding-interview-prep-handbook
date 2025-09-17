package completed

import "fmt"

/*
	Given an array arr of integers, check if there exist two indices i and j such that :
*/

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
	fmt.Println(checkIfExist([]int{3, 1, 7, 11}))
}

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[j] == arr[i]*2 {
				return true
			}
		}
	}

	return false
}
