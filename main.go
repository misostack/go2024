package main

import (
	"fmt"
	"math"
)

/*
The main function for Go Tutorial 2024
*/
func main() {
	// say hello to the world!
	fmt.Println("Go Tutorial 2024")
	// string
	fmt.Println(`Multiple line string
	in Go
	`)
	// character
	fmt.Println('A')
	// number
	fmt.Println("Calculation Multiplication: 5 * 3 = ", 5*3)
	fmt.Println("Calculation Substraction: 5 - 3 = ", 5-3)
	fmt.Println("Calculation Division: 5 / 3 = ", 5/3)
	fmt.Println("Calculation Modular: 5 % 3 = ", 5%3)
	fmt.Println("Calculation Division with Float Number: 5.0 / 3 = ", 5.0/3)
	fmt.Println("Calculation Exponents: 20^3 = ", math.Pow(20.0, 3))
	// boolean
	fmt.Println("Is 1 greater than 2: ", 1 > 2)
	fmt.Println("Is 1 less than 2: ", 1 < 2)
	fmt.Println("Is 1 equal 2: ", 1 == 2)
	fmt.Println("Equivalent: 2 = 2.0", 2 == 2.0)
	learning_variables()
	conditionals_demo()
	defer_demo()
	read_user_input_demo()
}
