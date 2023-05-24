package main

import (
	"fmt"

	"github.com/amirraouf/golang/09-deck/deck"
)

func main() {
	deck := deck.New(deck.OptionSort(
		func(i, j deck.Card) bool {
			return i.Suit > j.Suit || i.Suit == j.Suit && i.Value > j.Va
		}),
	)

	fmt.Println(deck)
}
