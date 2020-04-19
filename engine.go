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

var arrayLength, gamePlayed int

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

	fmt.Println("Game Played: ", gamePlayed)

	player := [10]int{}
	dealer := [10]int{}
	playerHand := [10]string{}
	dealerHand := [10]string{}
	dealerAce := 0
	playerAce := 0

	round := 0

	playableDeck, player, dealer, playerHand, dealerHand = deal(playableDeck, player, dealer, playerHand, dealerHand)

	pHand := 0
	dHand := 0

	pHand, playerAce = summation(player, playerAce)

	dHand, dealerAce = summation(dealer, dealerAce)

	fmt.Println("Player Hand: ", playerHand, ", ", pHand)
	fmt.Println("Dealer Hand: ", dealerHand, ", ", dHand)

	for {
		var input string
		fmt.Println("H - Hit, S = Stand: ")
		fmt.Scanln(&input)

		if input == "H" {
			playableDeck, player, playerHand, round, pHand, playerAce = hit(playableDeck, player, playerHand, round, playerAce)

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
			playableDeck, dealer, dealerHand, dHand, dealerAce = stand(playableDeck, dealer, dealerHand, dHand, dealerAce)
			break
		}
	}

	results := score(pHand, dHand)

	fmt.Println(results)

}

func checkAce(ace, hand int) (newHand int) {
	if hand > 21 {
		if ace > 0 {
			hand -= 10
			ace--
		}
	}
	return
}

func score(pHand int, dHand int) (result string) {
	if pHand > dHand && pHand <= 21 {
		result = "Player Wins"
	}
	if pHand == dHand {
		result = "Push"
	}
	if pHand < dHand && dHand <= 21 {
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
			if newDealer[di] == 1 {
				newDealerHand[di] = "A"
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
			if newPlayer[pi] == 1 {
				newPlayerHand[pi] = "A"
			}
			pi++
		}

		newDeck = RemoveIndex(newDeck, 0)
	}

	fmt.Println("\nUndealt: ", newDeck)
	return
}

func playerAdd(deck []int, side [10]int, sideHand [10]string, round int) (newSide [10]int, newSideHand [10]string) {

	newSide = side
	newSideHand = sideHand

	if deck[0] == 11 {
		newSideHand[round+1] = "J"
		newSide[round+1] = 10
	}
	if deck[0] == 12 {
		newSideHand[round+1] = "Q"
		newSide[round+1] = 10
	}
	if deck[0] == 13 {
		newSideHand[round+1] = "K"
		newSide[round+1] = 10
	}
	if deck[0] < 11 {
		newSideHand[round+1] = strconv.Itoa(deck[0])
		newSide[round+1] = deck[0]
	}
	if newSide[round+1] == 1 {
		newSideHand[round+1] = "A"
	}

	return
}

func hit(deck []int, player [10]int, playerHand [10]string, round, ace int) (newDeck []int, newPlayer [10]int, newPlayerHand [10]string, newRound int, pHand, newAce int) {

	fmt.Println("Player Hits")

	newRound = round + 1

	newPlayer = player

	newPlayerHand = playerHand

	newDeck = deck

	newPlayer, newPlayerHand = playerAdd(deck, player, playerHand, newRound)

	newDeck = RemoveIndex(deck, 0)

	pHand, newAce = summation(newPlayer, ace)

	return

}

func summation(hand [10]int, ace int) (sum, newAce int) {
	for i := range hand {
		if hand[i] > 10 {
			sum += 10
		} else if hand[i] == 1 {
			sum += 11
			if sum > 21 {
				sum = sum - 10
			}
		} else {
			sum += hand[i]
		}
	}

	newAce = checkAce(ace, sum)

	return
}

func stand(deck []int, dealer [10]int, dealerHand [10]string, dHand, dealerAce int) (newDeck []int, newDealer [10]int, newDealerHand [10]string, newdHand, newdAce int) {

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
	if newDealer[i] == 1 {
		newDealerHand[i] = "A"
	}

	newdHand, newdAce = summation(newDealer, dealerAce)

	fmt.Println("Dealer Hand: ", newDealerHand, ", ", newdHand)

	newDeck = RemoveIndex(newDeck, 0)

	i++

	newdHand, newdAce = summation(newDealer, dealerAce)

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
		if newDeck[0] == 1 {
			newDealerHand[i] = "A"
			newDealer[i] = 11
		}

		newDeck = RemoveIndex(newDeck, 0)

		i++

		newdHand, newdAce = summation(newDealer, dealerAce)

		fmt.Println("Dealer Hand: ", newDealerHand, ", ", newdHand)

	}

	//newdHand, newdAce = summation(newDealer, dealerAce)

	if newdHand <= 21 {
		fmt.Println("Dealer Stands")
	} else {
		fmt.Println("Dealer Busts")
		fmt.Println("Player Wins")
	}

	return
}
