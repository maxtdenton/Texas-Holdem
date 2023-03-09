package deck 

import (
	"fmt"
	"strconv"
)

// creating a cards deck (52 cards, 4 suits)
// each card needs a number and a suit

//initialising a Suit, 'type' base interface for all data types
type Suit int

//turning iota (from struct) into string to text for suits
func (s Suit) String() string {
	switch s {
	case Spades:
		return "spades" 
	case Hearts: 
		return "hearts"
	case Diamonds:
		return "diamonds"
	case Clubs:
		return "clubs"
	default:
		panic("invalid card suit... what did you do?")
	}
} 

//creating suits, iota is an incrementing data type therefor
// spades = 0, heart = 1, ... and so on... and so on...
const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

//initialising struct ... similar to object
type Card struct {
	suit Suit
	value int
	image string
}

func (c Card) String() string {
	value := strconv.Itoa(c.value)
	if c.value == 1 {
		value = "ace"
	} else if c.value == 11 {
		value = "jack"
	} else if c.value == 12 {
		value = "queen"
	} else if c.value == 13 {
		value = "king"
	}
	return fmt.Sprintf("%s of %s %s", value, c.suit, c.image)
}

//initialises new card by taking the params - number on card, suit of card
func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("a card cant be over 13 idiot!")
	}
	return Card {
		suit: s,
		value: v,
		image: fmt.Sprintf("%d_of_%s.png", v, s),
	}
}

type Deck [52]Card 

func New() Deck {
	var (
		nSuits	= 4
		nCards 	= 13	
		d 		= [52]Card{}
	)

	x := 0
  
	for i := 0; i < nSuits; i++ {
		for j := 0; j < nCards; j++ {
			 d[x] = NewCard(Suit(i), j+1)
			 x++  
		}
	}
	return d
}

func suitToUnicode(s Suit) string {  
	switch s {
	case Spades:
		return "♠" 
	case Hearts: 
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		panic("invalid card suit... what did you do?")
	}
}