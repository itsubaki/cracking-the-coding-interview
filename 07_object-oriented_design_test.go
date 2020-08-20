package main

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/hashtable"
)

func TestDeckOfCards(t *testing.T) {
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
