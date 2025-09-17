package completed

/*
Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written. Do the above modifications to the input
array in place and do not return anything.


*/

func main() {
	duplicateZeros([]int{1, 2, 3})
	duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0})
}

func duplicateZeros(arr []int) {
	var myArr []int

	for _, v := range arr {
		myArr = append(myArr, v)
		if v == 0 {
			myArr = append(myArr, 0)
		}
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = myArr[i]
	}

	//fmt.Println(arr)
}
