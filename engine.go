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

	length := len(cards)

	shuffled := make([]int, length)

	for i := 0; i < length; i++ {

		rand.Seed(time.Now().UnixNano())
		roll := rand.Intn(len(cards))

		fmt.Println("roll: ", roll)

		shuffled[i] = cards[roll]

		cards = RemoveIndex(cards, roll)

		fmt.Println("Remaining Cards: ", cards)
		fmt.Println("Shuffled: ", shuffled)
		fmt.Println("i: ", i)

	}

}

func RemoveIndex(i []int, index int) []int {
	return append(i[:index], i[index+1:]...)
}
