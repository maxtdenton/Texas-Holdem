package main

import  (
	"fmt"
	"strconv"
	"github.com/maxtdenton/texas-holdem/data"
	"github.com/maxtdenton/texas-holdem/lib"
	//"github.com/maxtdenton/texas-holdem/deck"
)

func PlayerCount() {
	strpl := strconv.Itoa(players)
	fmt.Print("Amount of players is now... " + strpl + "\n")
}


func PlayerCheck(num int) {
	players = players + num
}

func Reset() {
	var again string
	fmt.Print("\ncontinue?? (y/n) (mp/ap)\n")
	fmt.Print("------------------------------------------------------------\n")
	fmt.Scan(&again)
	send4 := lib.Up(again)
	lib.ClearTerm()
	if send4 == "Y" {
		main()
	} else if send4 == "MP" && players > 2 {
		PlayerCheck(-1)
		main()
	} else if send4 == "AP" && players < 7 {
		PlayerCheck(1)
		main()
	}
}

func NewRound() string {
	var again string
	fmt.Print("\ncontinue?? (y/c) (mp/ap)\n")
	fmt.Print("------------------------------------------------------------\n")
	fmt.Scan(&again)
	send4 := lib.Up(again)
	lib.ClearTerm()
	if send4 == "Y" {
		main()
	} else if send4 == "CP" && players > 2 {
		PlayerCheck(-1)
		main()
	} else if send4 == "AP" && players < 7 {
		PlayerCheck(1)
		main()
	}
	return send4
}

//HOW MANY PLAYERS ARE ON THE TABLE?
var players = 4
//POSITION ON TABLE
var position = 0

func main() {
	var c1, c2 string
	var finalNum, flushVal int
	PlayerCount()
	fmt.Print("whats you're hand? (best card first)\n")
	fmt.Scan(&c1, &c2, &flushVal)
	c1 = lib.Up(c1)
	c2 = lib.Up(c2)
	finalNum = data.TheHand(c1, c2, players)
	finalNum = finalNum + flushVal
	fmt.Print("Final Percent: ", finalNum, "%\n")
	//data.TheDraw(c1, c2, players, flushVal)
	Reset()
}


//printing a full deck
	//data := deck.New()
	//fmt.Println(data)
	//fmt.Println(d)