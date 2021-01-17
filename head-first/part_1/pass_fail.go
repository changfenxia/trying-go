package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter grade")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal((err))
	}

	input = strings.TrimSpace(input)
	res, _ := strconv.Atoi(input)
	if res >= 60 {
		fmt.Println("Your grade: ", res, " - You've passed!")
	} else {
		fmt.Println("Your grade: ", res, " - You've failed....")
	}
	
}

