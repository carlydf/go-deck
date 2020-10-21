package deck

import "math/rand"

type Deck []Card

func New(opts ...DeckOption) Deck {
	var deck Deck
	for _, s := range SuitsDefaultOrder() {
		for _, r := range RanksDefaultOrder() {
			deck = append(deck, Card{Rank: r, Suit: s})
		}
	}
	for _, opt := range opts {
		opt(&deck)
	}
	return deck
}

type DeckOption func(*Deck)

func WithJokers(numJokers int) DeckOption {
	return func(d *Deck) {
		for i := 0; i < numJokers; i++ {
			*d = append(*d, Card{Rank: "Joker", Suit: ""})
		}
	}
}

func WithShuffle() DeckOption {
	return func(d *Deck) {
		//rand.Shuffle(len(*d), func (i, j int) {
		//	(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
		//})
		perm := rand.Perm(len(*d))
		for i, j := range perm {
			(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
		}
	}
}

func SuitsDefaultOrder() [4]string {
	return [4]string{"Spade", "Diamond", "Club", "Heart"}
}

func RanksDefaultOrder() [13]string {
	return [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
}
