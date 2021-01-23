// Take an integer n (n >= 0) and a digit d (0 <= d <= 9) as an integer. Square all numbers k (0 <= k <= n) between 0 and n. Count the numbers of digits d used in the writing of all the k**2. Call nb_dig (or nbDig or ...) the function taking n and d as parameters and returning this count.

package main

import (
	"fmt"
	"strconv"
)


func main() {
	num, digit := 25, 1
	numberOfEnters := 0
	for i := 0; i < num; i++ {
		numberOfEnters += getNumberOfEnters((i + 1) * (i + 1), digit)
	}

	fmt.Print(numberOfEnters)
}

func getNumberOfEnters(num, digit int) int {
	numOfEnters := 0
	digitToString := strconv.Itoa(digit)
	for _, r := range strconv.Itoa(num) {
		if string(r) == digitToString {
			numOfEnters++
		}
	}

	return numOfEnters
}

// Easier way :D
// for i := 0; i < num; i++ {
// 	numOfEnters += strings.Count(strconv.Itoa(num), strconv.Itoa(digit))
// }