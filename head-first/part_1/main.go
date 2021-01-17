package main

import (
	"fmt"
	"reflect"
)

func main() {
	babe := "Carrie"
	var quantity int
	var length, width = 1.2, 4.23
	quantity = 54
	var customerName string
	customerName = "ЖОПА"
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(length, width)
	fmt.Println(quantity * int (length))
	fmt.Println(customerName + babe)
}