package main

import (
	"testing"
)

func TestNewCardSet(t *testing.T) {
	d := newCardSet()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 , but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("The card ordering is wrong")
	}

	if d[len(d)-1] != "Ten of Clubs" {
		t.Errorf("The card ordering is wrong")
	}

}