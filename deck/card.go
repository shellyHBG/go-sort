package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	suit int // 1: spade, 2: heart, 3: diamond, 4: club
	face int // 1-10, 11: Jack, 12: Queen, 13: King
}

type Deck []Card

func (c Card) printCard() {
	fmt.Printf("%+v\n", c)
}

func (d Deck) PrintDeck() {
	for i := 0; i < 52; i++ {
		fmt.Printf("%d: %+v\n", i, d[i])
	}
}

func (d Deck) Shuffle() {
	// Swap the values of two index randomly
	// =>O(n)

	// Generate a random seed with its random source built by current time
	// https://golang.org/pkg/math/rand/
	// https://golang.org/pkg/time/#Time.Unix
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	// Swap all of the Cards in Deck
	// https://golang.org/pkg/math/rand/#Intn
	// Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
	for i := 0; i < len(d); i++ {
		randIndex := ran.Intn(len(d))
		d[i], d[randIndex] = d[randIndex], d[i]
	}

}

func (d Deck) SortDeckByDESC() {
	sort.Sort(sortDeckDESC(d))
}

func (d Deck) SortDeckByASC() {
	sort.Sort(sortDeckASC(d))
}

func NewDeck() Deck {
	cards := make([]Card, 52)

	for s := 1; s <= 4; s++ {
		for f := 1; f <= 13; f++ {
			cards[(s-1)*13+(f-1)] = Card{suit: s, face: f}
			//cards[(s-1)*4+f].printCard()
		}
	}

	return Deck(cards)
}
