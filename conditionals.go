package main

import (
	"fmt"
	"math/rand"
)

func conditionals_demo() {
	// random
	fmt.Println("Conditional Demo")
	// random an integer number less than 100
	game_level := rand.Intn(101)
	fmt.Printf("Your lucky number is %d\n", game_level)
	win_prize := 0
	if game_level >= 50 {
		switch {
		case game_level > 60 && game_level <= 70:
			win_prize = 20
		case game_level > 70 && game_level <= 80:
			win_prize = 40
		case game_level > 80 && game_level <= 90:
			win_prize = 80
		case game_level > 90 && game_level < 100:
			win_prize = 160
		case game_level == 100:
			win_prize = 1000
		default:
			win_prize = 10
		}
		fmt.Printf("You win %d USD\n", win_prize)
	} else {
		fmt.Println("Good luck and try next time!")
	}
	if win_prize > 0 {
		// daysOfWeek
		days_of_week := map[int]string{
			0: "Monday",
			1: "Tuesday",
			2: "Wednesday",
			3: "Thursday",
			4: "Friday",
			5: "Saturday",
			6: "Sunday",
		}

		switch random_day_of_week := days_of_week[rand.Intn(len(days_of_week))]; random_day_of_week {
		case "Saturday":
			fmt.Printf("The prize will be double, final prize is : %d", win_prize*2)
		case "Sunday":
			fmt.Printf("The prize will be tripple, final prize is : %d", win_prize*3)
		default:
			fmt.Printf("The prize will be sent on %s", random_day_of_week)
		}

	}
}
