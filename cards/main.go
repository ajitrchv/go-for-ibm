package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	// also ```card := "Ace of spades"```
	// card := newCard()
	// fmt.Println(card)

	var cardSet deck

	cardSet = newCardSet()

	// cardSet.print()
	hand, remainingDeck := deal(cardSet, 5)
	fmt.Println("Current in hand:\n")
	hand.print()
	fmt.Println("===========================\n")
	fmt.Println("Current remaining:\n")
	remainingDeck.print()
	fmt.Println("===========================\n")


}
