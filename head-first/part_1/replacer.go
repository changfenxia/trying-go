package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "F#cking *d*ots"
	replacer := strings.NewReplacer("#", "u")
	replacerTwo := strings.NewReplacer("*", "i")

	fixed := replacer.Replace(broken)
	fixed = replacerTwo.Replace(fixed)

	fmt.Print(fixed)
}

