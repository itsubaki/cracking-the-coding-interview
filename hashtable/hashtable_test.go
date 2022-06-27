package hashtable_test

import (
	"testing"

	"github.com/itsubaki/cracking-the-coding-interview/hashtable"
)

func TestHashTable(t *testing.T) {
	cases := []struct {
		key   string
		value string
		size  int
	}{
		{"foo", "bar", 1},
		{"piyo", "fuga", 2},
		{"color", "red", 3},
		{"hoge", "hoge", 4},
	}

	table := hashtable.New()
	for _, c := range cases {
		table.Put(c.key, c.value)
		if table.Size == c.size {
			continue
		}
		t.Errorf("size=%d", table.Size)
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

	if table.Size != 0 {
		t.Errorf("size=%d", table.Size)
	}
}

func TestResize(t *testing.T) {
	cases := []struct {
		key   string
		value string
		size  int
	}{
		{"foo", "bar", 1},
		{"piyo", "fuga", 2},
		{"color", "red", 3},
		{"hoge", "hoge", 4},
	}

	table := hashtable.New()
	for _, c := range cases {
		table.Put(c.key, c.value)
		if table.Size == c.size {
			continue
		}
		t.Errorf("size=%d", table.Size)
	}

	table.Resize()
	if table.Capacity != 32 {
		t.Errorf("capacity=%d", table.Capacity)
	}
	if table.Size != 4 {
		t.Errorf("size=%d", table.Size)
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

	if table.Size != 0 {
		t.Errorf("size=%d", table.Size)
	}
}
