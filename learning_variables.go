package main

import "fmt"

func learning_variables(){
	fmt.Println("Variables")
	
	var tada, _ = "okie", false
	var t bool
	hours := 10 // variables shorthand
	var minutes, seconds, miliseconds = 32, 00, 001
	fmt.Printf("%d:%d:%0.2d:%0.3d", hours, minutes, seconds, miliseconds)
	fmt.Println(tada, t)
}

