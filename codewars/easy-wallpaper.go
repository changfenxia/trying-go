package main

import (
	"fmt"
	"math"
)


func main() {
	// your code
l, w, h := 0.0, 3.5, 3.0
  numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve","thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
  if (l == 0.0 || w == 0.0 || h == 0) {
	  fmt.Println("zero")
  }
  wallsSquare := (l + w) * 2 * h * 1.15
  
  singleRollSquare := 0.52 * 10
  
  numRolls := uint(math.Ceil(wallsSquare / singleRollSquare))
  fmt.Print(numbers[numRolls]) 
}