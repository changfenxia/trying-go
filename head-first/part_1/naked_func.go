package main

import "fmt"

func join(a, b string) (c string) {
	c = a + b
	return
}

func main() {
	fmt.Println(join("Hell", "o, world"))
}