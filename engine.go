package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 11 - J
// 12 - Q
// 13 - K

var diamonds, clubs, hearts, spades [13]int

var arrayLength int

func startGame() {

	for i := 0; i < 13; i++ {
		diamonds[i] = i + 1
		clubs[i] = i + 1
		hearts[i] = i + 1
		spades[i] = i + 1
	}

	cards := []int{}

	for i := 0; i < 13; i++ {
		cards = append(cards, diamonds[i])
	}

	for i := 0; i < 13; i++ {
		cards = append(cards, clubs[i])
	}

	for i := 0; i < 13; i++ {
		cards = append(cards, hearts[i])
	}

	for i := 0; i < 13; i++ {
		cards = append(cards, spades[i])
	}

	fmt.Println("Cards: ", cards)

	cards = shuffle(cards)

	fmt.Print("Final: ", cards)

}

func shuffle(deck []int) []int {

	shuffled := make([]int, len(deck))

	fmt.Println("Length of Shuffled: ", len(shuffled))

	for i := 0; i < len(shuffled); i++ {

		rand.Seed(time.Now().UnixNano())
		roll := rand.Intn(len(deck))

		shuffled[i] = deck[roll]

		deck = RemoveIndex(deck, roll)
	}

	return shuffled
}

func RemoveIndex(i []int, index int) []int {
	return append(i[:index], i[index+1:]...)
}

func deal() {

}
