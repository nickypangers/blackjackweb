package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 11 - J
// 12 - Q
// 13 - K

var diamonds, clubs, hearts, spades [13]int

var arrayLength int

func startGame() []int {

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

	return cards

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

func game() {

	playableDeck := startGame()

	player := [10]int{}
	dealer := [10]int{}
	playerHand := [10]string{}
	dealerHand := [10]string{}

	round := 0

	playableDeck, player, dealer, playerHand, dealerHand = deal(playableDeck, player, dealer, playerHand, dealerHand)

	pHand := 0
	dHand := 0

	pHand = summation(player)

	dHand = summation(dealer)

	fmt.Println("Player Hand: ", playerHand, ", ", pHand)
	fmt.Println("Dealer Hand: ", dealerHand, ", ", dHand)

	for {
		var input string
		fmt.Println("H - Hit, S = Stand: ")
		fmt.Scanln(&input)

		if input == "H" {
			playableDeck, player, playerHand, round, pHand = hit(playableDeck, player, playerHand, round)

			fmt.Println("Player Hand: ", playerHand, ", ", pHand)
			fmt.Println("Dealer Hand: ", dealerHand, ", ", dHand)

			if pHand > 21 {
				fmt.Println("Player Busts")
				fmt.Println("Dealer Wins")
				break
			}
		}

		if input == "S" {
			fmt.Println("Player Stands")
			playableDeck, dealer, dealerHand, dHand = stand(playableDeck, dealer, dealerHand, dHand)
			break
		}
	}

	results := score(pHand, dHand)

	fmt.Println(results)

}

func score(pHand int, dHand int) (result string) {
	if pHand > dHand {
		result = "Player Wins"
	}
	if pHand == dHand {
		result = "Push"
	}
	if pHand < dHand {
		result = "Dealer Wins"
	}

	return result
}

func deal(deck []int, player, dealer [10]int, playerHand [10]string, dealerHand [10]string) (newDeck []int, newPlayer, newDealer [10]int, newPlayerHand, newDealerHand [10]string) {

	pi := 0
	di := 0

	newPlayer = player
	newDealer = dealer
	newPlayerHand = playerHand
	newDealerHand = dealerHand
	newDeck = deck

	for i := 0; i < 3; i++ {

		if i == 1 {
			if newDeck[0] == 11 {
				newDealerHand[di] = "J"
				newDealer[di] = 10
			}
			if newDeck[0] == 12 {
				newDealerHand[di] = "Q"
				newDealer[di] = 10
			}
			if newDeck[0] == 13 {
				newDealerHand[di] = "K"
				newDealer[di] = 10
			}
			if newDeck[0] < 11 {
				newDealerHand[di] = strconv.Itoa(newDeck[0])
				newDealer[di] = newDeck[0]

			}
			di++
		} else {
			if newDeck[0] == 11 {
				newPlayerHand[pi] = "J"
				newPlayer[pi] = 10
			}
			if newDeck[0] == 12 {
				newPlayerHand[pi] = "Q"
				newPlayer[pi] = 10
			}
			if newDeck[0] == 13 {
				newPlayerHand[pi] = "K"
				newPlayer[pi] = 10
			}
			if newDeck[0] < 11 {
				newPlayerHand[pi] = strconv.Itoa(newDeck[0])
				newPlayer[pi] = newDeck[0]
			}
			pi++
		}

		newDeck = RemoveIndex(newDeck, 0)
	}

	fmt.Println("\nUndealt: ", newDeck)
	return
}

func hit(deck []int, player [10]int, playerHand [10]string, round int) (newDeck []int, newPlayer [10]int, newPlayerHand [10]string, newRound int, pHand int) {

	fmt.Println("Player Hits")

	newRound = round + 1

	newPlayer = player

	newPlayerHand = playerHand

	newDeck = deck

	if newDeck[0] == 11 {
		newPlayerHand[newRound+1] = "J"
		newPlayer[newRound+1] = 10
	}
	if newDeck[0] == 12 {
		newPlayerHand[newRound+1] = "Q"
		newPlayer[newRound+1] = 10
	}
	if newDeck[0] == 13 {
		newPlayerHand[newRound+1] = "K"
		newPlayer[newRound+1] = 10
	}
	if newDeck[0] < 11 {
		newPlayerHand[newRound+1] = strconv.Itoa(newDeck[0])
		newPlayer[newRound+1] = newDeck[0]
	}

	//newPlayer[newRound+1] = deck[0]

	newDeck = RemoveIndex(deck, 0)

	pHand = summation(newPlayer)

	return

}

func summation(hand [10]int) (sum int) {
	for i := range hand {
		if hand[i] > 10 {
			sum += 10
		} else {
			sum += hand[i]
		}
	}
	return sum
}

func checkAce() {

}

func stand(deck []int, dealer [10]int, dealerHand [10]string, dHand int) (newDeck []int, newDealer [10]int, newDealerHand [10]string, newdHand int) {

	newdHand = dHand
	newDealer = dealer
	newDealerHand = dealerHand
	newDeck = deck
	i := 1

	newDealer[i] = newDeck[0]
	if newDeck[0] == 11 {
		newDealerHand[i] = "J"
		newDealer[i] = 10
	}
	if newDeck[0] == 12 {
		newDealerHand[i] = "Q"
		newDealer[i] = 10
	}
	if newDeck[0] == 13 {
		newDealerHand[i] = "K"
		newDealer[i] = 10
	}
	if newDeck[0] < 11 {
		newDealerHand[i] = strconv.Itoa(newDeck[0])
		newDealer[i] = newDeck[0]
	}

	newdHand = summation(newDealer)

	fmt.Println("Dealer Hand: ", newDealerHand, ", ", newdHand)

	newDeck = RemoveIndex(newDeck, 0)

	i++

	newdHand = summation(newDealer)

	for newdHand < 17 {
		fmt.Println("Dealer Hits")
		if newDeck[0] == 11 {
			newDealerHand[i] = "J"
			newDealer[i] = 10
		}
		if newDeck[0] == 12 {
			newDealerHand[i] = "Q"
			newDealer[i] = 10
		}
		if newDeck[0] == 13 {
			newDealerHand[i] = "K"
			newDealer[i] = 10
		}
		if newDeck[0] < 11 {
			newDealerHand[i] = strconv.Itoa(newDeck[0])
			newDealer[i] = newDeck[0]
		}

		newDeck = RemoveIndex(newDeck, 0)

		i++

		newdHand = summation(newDealer)

		fmt.Println("Dealer Hand: ", newDealerHand, ", ", newdHand)

	}

	newdHand = summation(newDealer)

	if newdHand <= 21 {
		fmt.Println("Dealer Stands")
	} else {
		fmt.Println("Dealer Busts")
	}

	return
}
