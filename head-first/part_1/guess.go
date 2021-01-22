package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputNum, _ := strconv.Atoi(input)
	println("from func", inputNum)
	return inputNum
}

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(seed)
	num := generator.Intn(10)
	println(num)

	println("Guess number: ")
	inputNum := getInput()

	for ;num != inputNum; {
		println("you were wrong :( Not ", inputNum)
		println("Another guess?")
		inputNum = getInput()
	}

	println("You were right!")
	println("the number is ", num)
}
