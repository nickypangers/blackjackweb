package main

import "fmt"

func main() {

	var input string

	for {
		game()
		fmt.Println("End game, press 1 to continue, 2 to quit")
		fmt.Scanln(&input)

		if input == "1" {
			continue
		} else {
			break
		}
	}
}
