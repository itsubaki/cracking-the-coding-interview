package hashtable

import (
	"testing"
)

func TestHashTable(t *testing.T) {
	table := New()

	table.Put("hoge", "foobar")
	v, ok := table.Get("hoge")
	if !ok || v != "foobar" {
		t.Errorf("v=%v, ok=%v", v, ok)
	}

	table.Remove("hoge")
	_, ng := table.Get("hoge")
	if ng {
		t.Errorf("ok=%v", ng)
	}
}
