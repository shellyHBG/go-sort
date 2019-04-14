package main

import (
	"./deck"
)

func main() {

	// 0. In order to use "sort" function supported by standard package, three methods are neccessarily required to implement.
	// https://golang.org/pkg/sort/#Interface
	// Len() int  				Len is the number of elements in the collection.
	// Less(i, j int) bool		Less reports whether the element with index i should sort before the element with index j.
	// Swap(i, j int)			Swap swaps the elements with indexes i and j.

	// 1.1 Generate a sort of Cards
	Deck := deck.NewDeck()
	//Deck.PrintDeck()

	// 1.2 Shuffle the Deck
	Deck.Shuffle()
	//Deck.PrintDeck()

	// 2.1 Sort Cards DESC
	Deck.SortDeckByDESC()
	Deck.PrintDeck()

	// 2.2 Sort Cards ASC
	Deck.SortDeckByASC()
	//Deck.PrintDeck()
}
