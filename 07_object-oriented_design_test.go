package main_test

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/deck"
	"github.com/itsubaki/cracking-the-coding-interview/hashtable"
)

func TestDeckOfCards(t *testing.T) {
	d := deck.New()
	if d.Size() != 52 {
		t.Errorf("size: %v", d.Size())
	}

	cases := []struct {
		num   int
		value []deck.Card
	}{
		{
			3,
			[]deck.Card{
				{Suit: "Club", Value: 1},
				{Suit: "Club", Value: 2},
				{Suit: "Club", Value: 3},
			},
		},
		{
			4,
			[]deck.Card{
				{Suit: "Club", Value: 4},
				{Suit: "Club", Value: 5},
				{Suit: "Club", Value: 6},
				{Suit: "Club", Value: 7},
			},
		},
	}

	for _, c := range cases {
		cards := d.DealWith(c.num)
		if len(c.value) != len(cards) {
			t.Errorf("cards=%v", cards)
		}

		for i := 0; i < len(cards); i++ {
			if c.value[i].String() != cards[i].String() {
				t.Errorf("card=%v", cards[i])
			}
		}
	}

}

func TestCallCenter(t *testing.T) {

}

func TestJukebox(t *testing.T) {

}

func TestParkingLot(t *testing.T) {

}

func TestOnlineBookReader(t *testing.T) {

}

func TestJigsaw(t *testing.T) {

}

func TestChatServer(t *testing.T) {

}

func TestOthello(t *testing.T) {

}

func TestCircularArray(t *testing.T) {

}

func TestMinesweeper(t *testing.T) {

}

func TestFileSystem(t *testing.T) {

}

func TestHashTable(t *testing.T) {
	cases := []struct {
		key   string
		value string
	}{
		{"foo", "bar"},
		{"piyo", "fuga"},
		{"color", "red"},
		{"hoge", "hoge"},
	}

	table := hashtable.New()
	for _, c := range cases {
		table.Put(c.key, c.value)
	}

	for _, c := range cases {
		v, ok := table.Get(c.key)
		if ok && v == c.value {
			continue
		}

		t.Errorf("v=%v, ok=%v", v, ok)
	}

	for _, c := range cases {
		table.Remove(c.key)
		v, ok := table.Get(c.key)
		if ok {
			t.Errorf("v=%v, ok=%v", v, ok)
		}
	}
}
