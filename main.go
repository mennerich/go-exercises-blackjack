package main

import (
	"fmt"
	"github.com/mennerich/go-exercises-deck/deck"
)

var (
	d = deck.Shuffle(deck.GetStandardDeck())
	playerHand = []deck.Card{}
	dealerHand = []deck.Card{}
)

func main() {
	fmt.Println("+=============+")
	fmt.Println("|  BlackJack  |")
	fmt.Println("+=============+\n")
	fmt.Println("+===============+")
	fmt.Println("| Dealing Cards |")
	fmt.Println("+===============+")
	dealCards()
	showHand()
	tallyScore()
}

func dealCards() {
	for i := 0; i < 2; i++ {
		var c deck.Card
		c,d = deck.DealCard(d)
		playerHand = append(playerHand,c)
		c,d = deck.DealCard(d)
		dealerHand = append(dealerHand,c)
	}
}

func showHand() {
	fmt.Println("\nYour hand:")
	for _,c := range playerHand {
		fmt.Println("  ", c)
	}
	fmt.Println("\nDealer show a", dealerHand[len(dealerHand) - 1])
}

func tallyScore() {
	pValue := getValue(playerHand)
	dValue := getValue(dealerHand)
	fmt.Println("Player's hand value:", pValue)
	fmt.Println("Dealer's hand value:", dValue)

	if dValue > pValue { fmt.Println("Dealer Wins") }

	if pValue > dValue { fmt.Println("You Win") }

	if pValue == dValue{ fmt.Println("Draw, Dealer Winds") }
}

func getValue(hand []deck.Card) int {
	value := 0
	for _,i := range hand {
		if i.Value == 1 {
			value = value + 11
		}
		if i.Value > 1 && i.Value < 10 {
			value = value + i.Value
		}
		if i.Value >= 10 {
			value = value + 10
		}
	}
	return value
}
