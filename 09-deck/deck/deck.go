package deck

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

func New() []Card {
	var deck []Card
	for suit := Spades; suit <= Clubs; suit++ {
		for value := ValueTen; value <= ValueAce; value++ {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}
	return deck
}
