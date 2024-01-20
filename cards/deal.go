package main

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}