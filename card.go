package deck

import "fmt"

////go:generate stringer -type=Card
type Card struct {
	Rank string // A, 2, ..., 10, J, Q, K, Joker
	Suit string // spade, diamond, club, heart, "" if joker
}

func (c Card) String() string {
	if c.Rank == "Joker" {
		return fmt.Sprint(c.Rank)
	}
	return fmt.Sprint(c.Rank, " of ", c.Suit + "s")
}