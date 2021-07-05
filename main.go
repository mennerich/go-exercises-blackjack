package main

import (
	"bufio"
	"fmt"
	"github.com/mennerich/go-exercises-deck/deck"
	"os"
	"strings"
	"time"
)

var (
	d          = deck.Shuffle(deck.GetStandardDeck())
	playerHand = []deck.Card{}
	dealerHand = []deck.Card{}
	reader     = bufio.NewReader(os.Stdin)
)

func main() {
	fmt.Println("+=============+")
	fmt.Println("|  BlackJack  |")
	fmt.Println("+=============+")
	fmt.Println("+===============+")
	fmt.Println("| Dealing Cards |")
	fmt.Println("+===============+")
	dealCards()
	fmt.Printf("\nYour hand (%d):\n", getValue(playerHand))
	showHand(playerHand)
	if checkForBlackJack(playerHand) == true {
		fmt.Println("\nBlackJack, you win.")
	} else {
		showDealerHand()
		playerActions()
		dealerActions()
		tallyScore()
	}
}

func dealCards() {
	for i := 0; i < 2; i++ {
		var c deck.Card
		c, d = deck.DealCard(d)
		playerHand = append(playerHand, c)
		c, d = deck.DealCard(d)
		dealerHand = append(dealerHand, c)
	}
}

func showHand(h []deck.Card) {
	for _, c := range h {
		fmt.Println("  ", c)
	}
}

func showDealerHand() {
	fmt.Println("\nDealer shows a", dealerHand[len(dealerHand)-1])
}

func tallyScore() {

	pValue := getValue(playerHand)
	dValue := getValue(dealerHand)
	fmt.Println("\nPlayer's hand value:", pValue)
	fmt.Println("Dealer's hand value:", dValue)

	if dValue > pValue {
		fmt.Println("\nDealer Wins")
	}

	if pValue > dValue {
		fmt.Println("\nYou Win")
	}

	if pValue == dValue {
		fmt.Println("\nDraw, Dealer Winds")
	}
}

func getValue(hand []deck.Card) int {
	value := 0
	numAces := 0
	for _, i := range hand {
		if i.Value == 1 {
			value = value + 11
			numAces = numAces + 1
		}
		if i.Value > 1 && i.Value < 10 {
			value = value + i.Value
		}
		if i.Value >= 10 {
			value = value + 10
		}
	}

	for i := 0; i < numAces; i++ {
		if value > 21 {
			value = value - 10
		}
	}

	return value
}

func checkForBlackJack(h []deck.Card) bool {
	if getValue(h) == 21 {
		return true
	}
	return false
}

func playerActions() {
	stay := false
	for stay != true {

		fmt.Println("\n(H)it or (S)tay?")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "S" {
			stay = true
		}
		if text == "H" {
			var c deck.Card
			c, d = deck.DealCard(d)
			fmt.Println("\nCard Dealt:", c)
			playerHand = append(playerHand, c)
			fmt.Printf("\nYour Hand (%d):\n", getValue(playerHand))
			showHand(playerHand)
			if getValue(playerHand) > 21 {
				fmt.Println("\nBusted, Dealer wins")
				os.Exit(0)
			}
		}
	}
}

func dealerActions() {
	value := getValue(dealerHand)
	fmt.Printf("\nDealers Hand (%d):\n", value)
	showHand(dealerHand)
	for value < 17 {
		time.Sleep(2 * time.Second)
		fmt.Println("\nDealer Hits")
		var c deck.Card
		c, d = deck.DealCard(d)
		fmt.Println("\nCard Dealt:", c)
		dealerHand = append(dealerHand, c)
		value = getValue(dealerHand)
		fmt.Printf("\nDealers Hand (%d):\n", value)
		showHand(dealerHand)
		if value > 21 {
			fmt.Println("\nDealer Busts, you win.")
			os.Exit(0)
		}
	}
	time.Sleep(1 * time.Second)
	fmt.Println("\nDealer Stays")

}
