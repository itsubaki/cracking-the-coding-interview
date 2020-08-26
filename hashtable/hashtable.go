package hashtable

import (
	"fmt"
	"strings"
)

type HashTable struct {
	Table    []*Entry
	Capacity int
	Size     int
}

func New() *HashTable {
	return NewWith(1 << 4)
}

func NewWith(capacity int) *HashTable {
	table := make([]*Entry, capacity)

	return &HashTable{
		Table:    table,
		Capacity: capacity,
		Size:     0,
	}
}

func Hash(key string, capacity int) uint {
	return uint(len(key) % capacity)
}

func (t *HashTable) Put(key, value string) {
	defer func() {
		if t.Size > t.Capacity {
			t.Resize()
		}
	}()

	i := uint((t.Capacity - 1)) & Hash(key, t.Capacity)
	e := &Entry{Key: &key, Value: &value}

	p := t.Table[i]
	if p == nil {
		t.Table[i] = e
		t.Size++
		return
	}

	if *p.Key == key {
		t.Table[i] = e
		return
	}

	n := p
	for {
		if n.Next == nil {
			n.Next = e
			t.Size++
			break
		}

		n = n.Next
	}
}

func (t *HashTable) Get(key string) (string, bool) {
	i := uint((t.Capacity - 1)) & Hash(key, t.Capacity)

	p := t.Table[i]
	if p == nil {
		return "", false
	}

	if *p.Key == key {
		return *p.Value, true
	}

	n := p.Next
	for {
		if n == nil {
			return "", false
		}

		if *n.Key == key {
			return *n.Value, true
		}

		n = n.Next
	}
}

func (t *HashTable) Remove(key string) {
	i := uint((t.Capacity - 1)) & Hash(key, t.Capacity)

	p := t.Table[i]
	if p == nil {
		return
	}

	if *p.Key == key && p.Next == nil {
		t.Table[i] = nil
		t.Size--
		return
	}

	if *p.Key == key && p.Next != nil {
		t.Table[i] = p.Next
		t.Size--
	}

	prev := p
	next := p.Next
	for {
		if next == nil {
			break
		}

		if *next.Key == key {
			prev.Next = next.Next
			t.Size--
			break
		}

		prev = next
		next = next.Next
	}
}

func (t *HashTable) Resize() {
	newCapacity := t.Capacity << 1
	newTable := NewWith(newCapacity)

	for i := 0; i < t.Capacity; i++ {
		p := t.Table[i]
		if p == nil {
			continue
		}

		newTable.Put(*p.Key, *p.Value)

		n := p.Next
		for {
			if n == nil {
				break
			}

			newTable.Put(*n.Key, *n.Value)
			n = n.Next
		}
	}

	t.Table = newTable.Table
	t.Capacity = newTable.Capacity
	t.Size = newTable.Size
}

func (t *HashTable) String() string {
	var sb strings.Builder
	for _, e := range t.Table {
		if e == nil {
			continue
		}
		sb.WriteString(fmt.Sprintf("%s, ", e))

		n := e.Next
		for {
			if n == nil {
				break
			}

			sb.WriteString(fmt.Sprintf("%s, ", e))
			n = n.Next
		}
	}

	str := sb.String()
	return str[:len(str)-2] // remove ", "
}
