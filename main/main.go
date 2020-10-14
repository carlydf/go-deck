package main

import (
	"fmt"
	"github.com/gophercises/deck"
)

func main() {
	//deck := ordered()
	//deck := shuffled()
	deck := jokers(3)
	for _, c := range deck {
		fmt.Println(c)
	}
}

func ordered() deck.Deck {
	return deck.New()
}

func shuffled() deck.Deck {
	return deck.New(deck.WithShuffle())
}

func jokers(n int) deck.Deck {
	return deck.New(deck.WithJokers(n))
}