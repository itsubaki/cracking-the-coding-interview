package hashtable

import (
	"testing"
)

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

	table := New()
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
