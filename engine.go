package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var deck []int

var gameInt int

func game(playerScore, dealerScore int) (newPlayerScore, newDealerScore int) {

	// dealer face cards
	dealerHand := []string{}
	playerHand := []string{}

	// dealer face values
	dealer := []int{}
	player := []int{}

	dealerPoints := 0
	playerPoints := 0

	var input string

	i := 0

	fmt.Println("gameInt: ", gameInt)

	if gameInt == 0 {
		setDeck()
		deck = shuffle()
		gameInt++
	}

	fmt.Println("Deck length: ", len(deck))
	// deal
	dealerHand, playerHand, dealer, player, dealerPoints, playerPoints = deal(dealerHand, playerHand, dealer, player, dealerPoints, playerPoints)

	fmt.Println("Player: ", playerHand, ", ", playerPoints)
	fmt.Println("Dealer: ", dealerHand, ", ", dealerPoints)

	for {
		if playerPoints == 21 {
			input = "S"
		} else {
			input = getInput(input)
		}
		if input == "H" || input == "h" {

			// hit
			playerHand, player, playerPoints = hit(playerHand, player, playerPoints)
			fmt.Println("Player: ", playerHand, ", ", playerPoints)
			fmt.Println("Dealer: ", dealerHand, ", ", dealerPoints)
			i++
			if playerPoints > 21 {
				fmt.Println("Player Busts")
				fmt.Println("Dealer Wins")
				break
			}
		}
		if input == "S" || input == "s" {
			fmt.Println("Player Stands")
			fmt.Println("Player: ", playerHand, ", ", playerPoints)
			// stand -> dealer plays
			dealerHand, dealer, dealerPoints = dealerPlay(dealerHand, dealer, dealerPoints)
			fmt.Println("Player: ", playerHand, ", ", playerPoints)
			fmt.Println("Dealer: ", dealerHand, ", ", dealerPoints)
			break
		}
	}

	result := checkScore(dealerPoints, playerPoints)

	if result == 1 {
		newPlayerScore = playerScore + 1
		newDealerScore = dealerScore
		return newPlayerScore, newDealerScore
	}
	newPlayerScore = playerScore
	newDealerScore = dealerScore + 1
	return newPlayerScore, newDealerScore
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

	// shuffle cards
	for i := 0; i < deckLength; i++ {

		// create random int
		rand.Seed(time.Now().UnixNano())
		roll := rand.Intn(len(deck))

		// append randomly selected card in deck to shuffle
		shuffled = append(shuffled, deck[roll])

		// fmt.Println("Shuffled length: ", len(shuffled))

		// fmt.Println("Shuffled: ", shuffled[i], ", roll: ", roll)

		// remove selected card from deck
		deck = removeIndex(deck, roll)
	}

	// fmt.Println("Shuffled: ", shuffled)

	// fmt.Println("Check: ", shuffled[1])

	// fmt.Println("Length of shuffled: ", len(shuffled))

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

func deal(dealerHand, playerHand []string, dealer, player []int, dealerPoints, playerPoints int) (newDealerHand, newPlayerHand []string, newDealer, newPlayer []int, newdealerPoints, newplayerPoints int) {

	newDealerHand = dealerHand
	newPlayerHand = playerHand
	dealer = newDealer
	player = newPlayer
	newdealerPoints = dealerPoints
	newplayerPoints = playerPoints

	for i := 0; i < 3; i++ {
		// set cards to respective J, Q, K or A if selected
		if i == 1 {
			newDealerHand, newDealer, newdealerPoints = appendHand(newDealerHand, newDealer, newdealerPoints)
		} else {
			newPlayerHand, newPlayer, newplayerPoints = appendHand(newPlayerHand, newPlayer, newplayerPoints)
		}
	}

	return
}

func hit(playerHand []string, player []int, playerPoints int) (newPlayerHand []string, newPlayer []int, newplayerPoints int) {

	newPlayerHand = playerHand
	newPlayer = player
	newplayerPoints = playerPoints

	newPlayerHand, newPlayer, newplayerPoints = appendHand(newPlayerHand, newPlayer, newplayerPoints)

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

func dealerPlay(dealerHand []string, dealer []int, dealerPoints int) (newDealerHand []string, newDealer []int, newdealerPoints int) {

	newDealer = dealer
	newDealerHand = dealerHand
	newdealerPoints = dealerPoints

	newDealerHand, newDealer, newdealerPoints = appendHand(newDealerHand, newDealer, newdealerPoints)
	fmt.Println("Dealer: ", newDealerHand, ", ", newdealerPoints)

	for newdealerPoints < 17 {
		fmt.Println("Dealer Hits")
		newDealerHand, newDealer, newdealerPoints = appendHand(newDealerHand, newDealer, newdealerPoints)
		fmt.Println("Dealer: ", newDealerHand, ", ", newdealerPoints)
		if newdealerPoints > 21 {
			fmt.Println("Dealer Busts")
			return
		}
	}

	fmt.Println("Dealer Stands")

	// newdealerPoints = summation(newDealer)

	return

}

func checkScore(dealerPoints, playerPoints int) int {
	result := 0
	if playerPoints <= 21 && dealerPoints <= 21 {
		if playerPoints > dealerPoints {
			fmt.Println("Player Wins")
			result = 1
		} else if playerPoints == dealerPoints {
			fmt.Println("Push")
		} else {
			fmt.Println("Dealer Wins")
		}
	} else if dealerPoints > 21 {
		fmt.Println("Player Wins")
		result = 1
	}
	return result
}
