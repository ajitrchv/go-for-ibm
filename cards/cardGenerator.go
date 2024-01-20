package main

import "fmt"

func newCardSet() []string {

	var cardSet deck

	items := []string{"Ace", "Queen", "King", "Jack", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	symbols := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	for _, symbol := range symbols {
		for _, item := range items {
			card := fmt.Sprintf("%s of %s", item, symbol)
			cardSet = append(cardSet, card)
		}
	}

	return cardSet

}