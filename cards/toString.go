package main

import "strings"


func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}