package main

import (
	// "fmt"
	// "bufio"
	// "os"
	// "strconv"
	"math/rand"
	"time"
)

func main() {
	// get random number
	s1 := rand.NewSource((time.Now().UnixNano()))
	r1 := rand.New(s1)
	num := r1.Intn(10)
	num = 5

	// get user input
	input := 5
	// print both
	println(num, input)
	// implement loop & conditions
}