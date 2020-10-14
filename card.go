package deck

import "fmt"

////go:generate stringer -type=Card
type Card struct {
	Rank string // A, 2, ..., 10, J, Q, K, Joker
	Suit string // Spade, Diamond, Club, Heart, "" if joker
}

func (c Card) String() string {
	if c.Rank == "Joker" {
		return fmt.Sprint(c.Rank)
	} else if c.Rank == "10" {
		return fmt.Sprint(c.Rank, " of ", c.Suit + "s")
	}
	return fmt.Sprint(" ", c.Rank, " of ", c.Suit + "s")
}