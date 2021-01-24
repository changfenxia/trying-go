package main

import "fmt"

func main() {
	trueish := true
	bigish := 999999999999.9999
	stringish := "hello worldish"
	fmt.Printf("trueish %T, bigish %T, stringish %T", trueish, bigish, stringish)
}
