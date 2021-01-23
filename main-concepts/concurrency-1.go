package main

import (
	"time"
)

func printer(thing string, c chan string) {
	for i := 1; i < 100; i++ {
		c <- thing
		
		time.Sleep(time.Millisecond * 700)
	}
}

func main() {
	c := make(chan string)
	go printer("teeeeeemo", c)

	msg := <- c
	println(msg)

}