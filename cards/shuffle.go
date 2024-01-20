package main

import "math/rand"

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition],d[i]
	}
}