package deck

import "fmt"

const (
	Club    = "Club"
	Diamond = "Diamond"
	Heat    = "Heat"
	Spade   = "Spade"
)

type Card struct {
	Suit  string
	Value int
}

func (c Card) String() string {
	return fmt.Sprintf("%s%02d", c.Suit, c.Value)
}
