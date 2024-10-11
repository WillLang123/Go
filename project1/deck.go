package main

import "fmt"

// Create a new type of deck aka slice
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
