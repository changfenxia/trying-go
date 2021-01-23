package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	// get random number - DONE
	s1 := rand.NewSource((time.Now().UnixNano()))
	r1 := rand.New(s1)
	num := r1.Intn(10)

	guess := getInput()
	tries := 1

	for ;guess != num; {
		if (guess > num) {
			fmt.Println("Wrong! Too high! Try again!")
		} else {
			fmt.Println("Yay! Too low... Try again")
		}
		guess = getInput()
		tries++
	}
	
	fmt.Println("You are right!")
	fmt.Println("The number was:", num)
	fmt.Println("Your guess:", guess)
	fmt.Println("It took", tries, "tries")
}

func getInput() int{
	fmt.Println("Enter number:")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputNum, _ := strconv.Atoi(strings.TrimSpace(input))
	return inputNum
}