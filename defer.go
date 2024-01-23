package main

import (
	"fmt"
	"strings"
)

func defer_demo() {
	fmt.Println(strings.Repeat("-", 50))
	// last-in first out
	defer fmt.Println("1st line of code with defer")
	defer fmt.Println("2nd line of code with defer")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("In loop i = %d\n", i+1)
	}

	fmt.Println("Last line of code")
}
