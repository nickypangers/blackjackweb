package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var deck []int

func game() {

	// dealer face cards
	dealerHand := []string{}
	playerHand := []string{}

	// dealer face values
	dealer := []int{}
	player := []int{}

	dealerScore := 0
	playerScore := 0

	game := 0

	var input string

	i := 0

	if game == 0 {
		setDeck()
		deck = shuffle()

		game++
	}
	// deal
	dealerHand, playerHand, dealer, player, dealerScore, playerScore = deal(dealerHand, playerHand, dealer, player, dealerScore, playerScore)

	fmt.Println("Player: ", playerHand, ", ", playerScore)
	fmt.Println("Dealer: ", dealerHand, ", ", dealerScore)

	for {
		input = getInput(input)
		if input == "H" || input == "h" {

			// hit
			playerHand, player, playerScore = hit(playerHand, player, playerScore)
			fmt.Println("Player: ", playerHand, ", ", playerScore)
			fmt.Println("Dealer: ", dealerHand, ", ", dealerScore)
			i++
			if playerScore > 21 {
				fmt.Println("Player Busts")
				fmt.Println("Dealer Wins")
				break
			}
		}
		if input == "S" || input == "s" {
			fmt.Println("Player Stands")
			fmt.Println("Player: ", playerHand, ", ", playerScore)
			// stand -> dealer plays
			dealerHand, dealer, dealerScore = dealerPlay(dealerHand, dealer, dealerScore)
			fmt.Println("Player: ", playerHand, ", ", playerScore)
			fmt.Println("Dealer: ", dealerHand, ", ", dealerScore)
			break
		}
	}

	checkScore(dealerScore, playerScore)

}

func getInput(input string) string {
	fmt.Println("Hit (H) or Stand (S)")
	fmt.Scanln(&input)

	return input
}

func setDeck() {

	// set suits
	diamonds := make([]int, 13)
	clubs := make([]int, 13)
	hearts := make([]int, 13)
	spades := make([]int, 13)

	// assign cards to each suit
	for i := 0; i < 13; i++ {
		diamonds[i] = i + 1
		clubs[i] = i + 1
		hearts[i] = i + 1
		spades[i] = i + 1
	}

	// combine all cards into one deck
	for ii := 0; ii < 4; ii++ {
		for i := 0; i < len(diamonds); i++ {
			deck = append(deck, diamonds[i])
		}

		for i := 0; i < len(clubs); i++ {
			deck = append(deck, clubs[i])
		}

		for i := 0; i < len(hearts); i++ {
			deck = append(deck, hearts[i])
		}

		for i := 0; i < len(spades); i++ {
			deck = append(deck, spades[i])
		}
	}

	fmt.Println("deck: ", deck, ", length: ", len(deck))

	// return combined deck
	return

}

// FUNC shuffle deck
func shuffle() []int {

	deckLength := len(deck)

	// create new array for shuffled deck
	shuffled := []int{}

	// fmt.Println("Length of shuffled: ", len(shuffled))

	// shuffle cards
	for i := 0; i < deckLength; i++ {

		// create random int
		rand.Seed(time.Now().UnixNano())
		roll := rand.Intn(len(deck))

		// append randomly selected card in deck to shuffle
		shuffled = append(shuffled, deck[roll])

		fmt.Println("Shuffled length: ", len(shuffled))

		//fmt.Println("Shuffled: ", shuffled[i], ", roll: ", roll)

		// remove selected card from deck
		deck = removeIndex(deck, roll)
	}

	fmt.Println("Shuffled: ", shuffled)

	fmt.Println("Check: ", shuffled[1])
	// return shuffled deck
	return shuffled
}

// FUNC remove selected card from deck
func removeIndex(i []int, index int) []int {
	return append(i[:index], i[index+1:]...)
}

//

func appendHand(hand []string, side []int, sideScore int) (newHand []string, newSide []int, newSideScore int) {

	newHand = hand
	newSide = side

	if deck[0] == 11 {
		newHand = append(newHand, "J")
		newSide = append(newSide, 10)
	} else if deck[0] == 12 {
		newHand = append(newHand, "Q")
		newSide = append(newSide, 10)
	} else if deck[0] == 13 {
		newHand = append(newHand, "K")
		newSide = append(newSide, 10)
	} else if deck[0] == 1 {
		newHand = append(newHand, "A")
		newSide = append(newSide, 11)
	} else {
		newHand = append(newHand, strconv.Itoa(deck[0]))
		newSide = append(newSide, deck[0])
	}

	newSideScore = summation(newSide)

	deck = removeIndex(deck, 0)

	return

}

func deal(dealerHand, playerHand []string, dealer, player []int, dealerScore, playerScore int) (newDealerHand, newPlayerHand []string, newDealer, newPlayer []int, newDealerScore, newPlayerScore int) {

	newDealerHand = dealerHand
	newPlayerHand = playerHand
	dealer = newDealer
	player = newPlayer
	newDealerScore = dealerScore
	newPlayerScore = playerScore

	for i := 0; i < 3; i++ {
		// set cards to respective J, Q, K or A if selected
		if i == 1 {
			newDealerHand, newDealer, newDealerScore = appendHand(newDealerHand, newDealer, newDealerScore)
		} else {
			newPlayerHand, newPlayer, newPlayerScore = appendHand(newPlayerHand, newPlayer, newPlayerScore)
		}
	}

	return
}

func hit(playerHand []string, player []int, playerScore int) (newPlayerHand []string, newPlayer []int, newPlayerScore int) {

	newPlayerHand = playerHand
	newPlayer = player
	newPlayerScore = playerScore

	newPlayerHand, newPlayer, newPlayerScore = appendHand(newPlayerHand, newPlayer, newPlayerScore)

	return
}

func summation(hand []int) (sum int) {
	ace := 0
	for i := range hand {
		sum += hand[i]
		if hand[i] == 11 {
			ace++
		}
		if sum > 21 && ace > 0 {
			sum -= 10
			ace--
		}
	}

	return

}

func dealerPlay(dealerHand []string, dealer []int, dealerScore int) (newDealerHand []string, newDealer []int, newDealerScore int) {

	newDealer = dealer
	newDealerHand = dealerHand
	newDealerScore = dealerScore

	newDealerHand, newDealer, newDealerScore = appendHand(newDealerHand, newDealer, newDealerScore)
	fmt.Println("Dealer: ", newDealerHand, ", ", newDealerScore)

	for newDealerScore < 17 {
		fmt.Println("Dealer Hits")
		newDealerHand, newDealer, newDealerScore = appendHand(newDealerHand, newDealer, newDealerScore)
		fmt.Println("Dealer: ", newDealerHand, ", ", newDealerScore)
	}

	fmt.Println("Dealer Stands")

	// newDealerScore = summation(newDealer)

	return

}

func checkScore(dealerScore, playerScore int) {
	if playerScore > dealerScore && playerScore <= 21 {
		fmt.Println("Player Wins")
	}
	if dealerScore > playerScore && dealerScore <= 21 {
		fmt.Println("Dealer Wins")
	}
	if dealerScore == playerScore {
		fmt.Println("Push")
	}
}
