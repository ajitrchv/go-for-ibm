package main

import "io/ioutil"

func (d deck) toFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}