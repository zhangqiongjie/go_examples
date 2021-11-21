package main

import (
	"fmt"
	"strings"
)

type deck []string

func newCard() deck {
	cards := deck{}
	deckSuit := []string{"Spades", "Hearts", "Diamonds", "Clubes"}
	deckValue := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range deckSuit {
		for _, value := range deckValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
