package main

import "fmt"

func print_board(boards map[int]string){
	fmt.Println(boards)
}

func learning_variables(){
	fmt.Println("Variables")
	
	var tada, _ = "okie", false
	var t bool
	hours := 10 // variables shorthand
	var minutes, seconds, miliseconds = 32, 00, 001
	fmt.Printf("%d:%d:%0.2d:%0.3d", hours, minutes, seconds, miliseconds)
	fmt.Println(tada, t)

	// array
	names := [3]string{"A", "B", "C"}
	names[0] = "X"
	names[1] = "Y"	
	// names[2] = "Z"
	fmt.Println("names[2] is nil ", names[2] == "")
	fmt.Println(names)

	countries := []string{}
	countries = append(countries, "Viet Nam")
	countries = append(countries, "Japan")
	countries = append(countries, "France")
	countries = append(countries, "China")
	countries = append(countries, "Russia")
	countries = append(countries, "United State")
	countries = append(countries, "United KingDom")
	fmt.Println(countries)

	// multi dimension arrays
	boards := map[int]string{
		0: "",
		1: "",
		3: "",
		4: "",
		5: "",
		6: "",
		7: "",
		8: "",
	}
	print_board(boards)
	boards[0] = "X"
	print_board(boards)
	boards[4] = "Y"
	print_board(boards)
	
	// another example about map
	settings := map[string]string{}
	settings["rows_per_page"] = "1"
	settings["max_rows_per_page"] = "100"
	fmt.Println(settings)
	delete(settings, "max_rows_per_page")
	fmt.Println(settings)

	colors := map[string]string{}
	colors["default"] = "\033[0m"
	colors["green"] = "\033[32m"
	colors["blue"] = "#0000FF"
	colors["red"] = "\033[31m"
	colors["yellow"] = "\033[33m"
	fmt.Println(colors["red"], "test[red]")
	fmt.Println(colors["default"], "test[defaut]")
}

