package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type deck []string

/*
*

	Factory methods

*
*/
func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Heart", "Diamond", "Clubs"}
	cardNums := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	for _, suite := range cardSuites {
		for _, cardNum := range cardNums {
			cards = append(cards, cardNum+" of "+suite)
		}
	}

	return cards
}

func readFromFile(filename string) (deck, error) {
	deckBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	deckStr := string(deckBytes)
	deckStrSlice := strings.Split(deckStr, ",")
	return deck(deckStrSlice), nil
}

/**
	Receiver methods
 **/

func (d deck) print() {
	for _, str := range d {
		fmt.Println(str)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d *deck) deal(size int) deck {
	res := (*d)[:size]
	(*d) = (*d)[size:]
	return res
}

func (d *deck) shuffle() {
	for idx := range *d {
		position := rand.IntN(len(*d))
		(*d)[idx], (*d)[position] = (*d)[position], (*d)[idx]
		//	(*d)[swapWith] = card
	}
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
