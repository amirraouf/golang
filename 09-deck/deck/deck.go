package deck

import (
	"math/rand"
	"sort"
)

type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

func (s Suit) String() string {
	switch s {
	case Spades:
		return "♠️"
	case Hearts:
		return "❤"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣️"
	default:
		return ""

	}
}

type Value int

const (
	_ Value = iota
	_
	ValueTwo
	ValueThree
	ValueFour
	ValueFive
	ValueSix
	ValueSeven
	ValueEight
	ValueNine
	ValueTen
	ValueJoker
	ValueQueen
	ValueKing
	ValueAce
)

type Card struct {
	Suit  Suit // enum
	Value Value
}

type Option func([]Card) []Card

func New(options ...Option) []Card {
	var deck []Card
	for suit := Spades; suit <= Clubs; suit++ {
		for value := ValueTwo; value <= ValueAce; value++ {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}
	for _, option := range options {
		deck = option(deck)
	}
	return deck
}

// OptionSort can sort a deck based on the sorting function fn
func OptionSort(fn func(Card, Card) bool) Option {
	return func(deck []Card) []Card {
		sort.Slice(deck, func(i, j int) bool {
			return fn(deck[i], deck[j])
		})
		return deck
	}
}

// OptionShuffle shuffles a deck
func OptionShuffle() Option {
	return func(deck []Card) []Card {
		rand.Shuffle(len(deck), func(i, j int) {
			deck[i], deck[j] = deck[j], deck[i]
		})
		return deck
	}
}

// OptionExclude uses fn to know which cards to exludes from a deck
func OptionExclude(fn func(Card) bool) Option {
	return func(deck []Card) []Card {
		var newDeck []Card
		for _, c := range deck {
			if fn(c) {
				continue
			}
			newDeck = append(newDeck, c)
		}
		return newDeck
	}
}

// OptionCompose composes a bigger deck by adding other decks to a deck
func OptionCompose(decks ...[]Card) Option {
	return func(deck []Card) []Card {
		for _, d := range decks {
			deck = append(deck, d...)
		}
		return deck
	}
}
