package deck

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
	Index int
}

func New() *Deck {
	cards := make([]Card, 0)
	for _, s := range []string{Club, Diamond, Heat, Spade} {
		for j := 1; j < 14; j++ {
			cards = append(cards, Card{
				Suit:  s,
				Value: j,
			})
		}
	}

	return &Deck{
		Cards: cards,
		Index: 0,
	}
}

func (d *Deck) Deal() Card {
	c := d.Cards[d.Index]
	d.Index++
	return c
}

func (d *Deck) DealWith(n int) []Card {
	c := d.Cards[d.Index : d.Index+n]
	d.Index = d.Index + n
	return c
}

func (d *Deck) Size() int {
	return len(d.Cards)
}

func (d *Deck) Remain() int {
	return d.Size() - d.Index
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().Unix())

	for i := 0; i < len(d.Cards); i++ {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
