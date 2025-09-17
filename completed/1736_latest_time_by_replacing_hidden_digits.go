package completed

import "fmt"

/*
You are given a string time in the form of  hh:mm, where some of the digits in the string are hidden (represented by ?).

The valid times are those inclusively between 00:00 and 23:59.

Return the latest valid time you can get from time by replacing the hidden digits.
*/

func main() {
	fmt.Println(maximumTime("2?:?0"))
	fmt.Println(maximumTime("0?:3?"))
	fmt.Println(maximumTime("1?:22"))
}

func maximumTime(time string) string {

	timeString := []byte(time)

	if timeString[0] == '?' {
		timeString[0] = '1'
	}
	if timeString[1] == '?' {
		if timeString[0] == '2' {
			timeString[1] = '3'
		} else {
			timeString[1] = '9'
		}
	}

	if timeString[3] == '?' {
		timeString[3] = '5'
	}
	if timeString[4] == '?' {
		timeString[4] = '9'
	}

	return string(timeString)
}
