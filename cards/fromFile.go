package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ", ")
	return deck(s)
	
}