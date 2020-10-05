package main

import "fmt"

func main() {

	var input string

	playerScore := 0
	dealerScore := 0

	for {
		playerScore, dealerScore = game(playerScore, dealerScore)
		fmt.Printf("Your score: %d\nDealer's score: %d\n", playerScore, dealerScore)
		// fmt.Println("Your score: ", playerScore, "")
		fmt.Println("End game, press anything to continue, Q to quit")
		fmt.Scanln(&input)

		if input != "q" {
			continue
		} else {
			break
		}
	}
}
