package deck

type Card struct {
	Suit string // spade, diamond, club, heart, "" if joker
	Rank string // A, 2, ..., 10, J, Q, K, Joker
}