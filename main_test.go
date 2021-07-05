package main

import (
	"github.com/mennerich/go-exercises-deck/deck"
	"testing"
)

func TestDeck(t *testing.T) {

	t.Run("Test Value Creation", func(t *testing.T) {
		hand := []deck.Card{{1, 1}, {1, 2}}
		want := 12
		got := getValue(hand)
		if got != want {
			t.Errorf("Wanted %d, got %d", want, got)
		}
	})
}
