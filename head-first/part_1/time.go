package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	// var year int = now.Minute()
	fmt.Println(now.Hour(), now.Minute(), now.Second())	
}