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
	num := r1.Intn(20)

	tries := 3
	fmt.Println("Guess the number! You have", tries, "guesses!")
	guess := getInput()

	for ;guess != num && tries > 0; {
		tries--
		if (guess > num) {
			fmt.Println("Wrong! Too high! Try again!", tries, "guesses left!")
		} else {
			fmt.Println("Yay! Too low... Try again", tries, "guesses left!")
		}
		guess = getInput()
	}

	if tries > 0 {
		fmt.Println("You are right!")
		fmt.Println("The number was:", num)
		fmt.Println("Your guess:", guess)
		fmt.Println("It took", tries, "tries")
	} else {
		fmt.Println("You lost :( The number was: ", num)
	}
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