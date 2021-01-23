package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 1234
	fmt.Println(reflect.TypeOf(fmt.Sprint(num)))
}