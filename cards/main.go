package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	// also ```card := "Ace of spades"```
	// card := newCard()
	// fmt.Println(card)

	var cardSet deck


	cardSet = newCardSet()

	for i,card := range cardSet{
		fmt.Printf("[%d] : %s\n", i, card)
	}

}


