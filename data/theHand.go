package data

import (
	"fmt"
	"os"
	"io"
	"log"
	"strconv"
	"encoding/csv"
)

/* CHECKING FOR SUITED
func HandSuited() {
	fmt.Print("\nSUITED: You're 3% more likely to win this hand\n")
	fmt.Print("------------------------------------------------------------")
}
*/

type Suit string

type Card struct {
	suit Suit
	value int
}

func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("a card cant be over 13 idiot!")
	}
	return Card {
		suit: s,
		value: v,
	}
}

type Hand [5]Card 

const (
	card1 int = iota
	card2
	suited
	players2
	players3
	players4
	players5
	players6
	players7
	players8
)

var position = 100

func TheHand(c1 string, c2 string, p1 int) (int){
	//fmt.Print("you got to here")
	var finalNum int
	f, ferr := os.Open("Percentages.csv")
	if ferr != nil {
	panic(ferr)
	}
	reader := csv.NewReader(f)
	//discarding header
	_, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	rowNo := 0	
	for {
		rowNo++
		row, err := reader.Read()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if row[card1] == c1 && row[card2] == c2 {
			fmt.Print("\n------------------------------------------------------------")
			fmt.Print("\nYour percentage chance of WINNING is... " + row[p1+1] + "%\n")
			finalNum, err = strconv.Atoi(row[p1+1])
			fmt.Print("------------------------------------------------------------")
			if row[p1+1] > "60" {
				fmt.Print("ANSWER: You should RAISE... high chance of the best hand\n")
			} else if row[p1+1] > "30" {
				fmt.Print("\nANSWER: You should CALL... if your suited RAISE\n")
			} else if row[card1] < "9" && row[card2] < "10" && p1 > 2 {
				fmt.Print("ANSWER: Think about folding or low betting... LIKELYHOOD: LOSE")
			} else if row[p1+1] < "30" {
				fmt.Print("\nANSWER: This is a dangerous bet FOLD or bet VERY LOW\n")
			} else {
				fmt.Print("\nANSWER: play tight and see what happens\n")
			}
			fmt.Print("------------------------------------------------------------\n")
			return finalNum
		}
	}
	return finalNum
}
