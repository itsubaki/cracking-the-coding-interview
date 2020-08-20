package hashtable

type HashTable struct {
	Table    []*Entry
	Capacity int
}

func New() *HashTable {
	capacity := 1 << 4
	table := make([]*Entry, capacity)

	return &HashTable{
		Table:    table,
		Capacity: capacity,
	}
}

func Hash(key string, capacity int) uint {
	return uint(len(key) % capacity)
}

func (t *HashTable) Put(key, value string) {
	e := &Entry{Key: &key, Value: &value}
	i := uint((t.Capacity - 1)) & Hash(key, t.Capacity)

	p := t.Table[i]
	if p == nil || *p.Key == key {
		t.Table[i] = e
		return
	}

	n := p
	for {
		if n.Next == nil {
			n.Next = e
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

	if *p.Key == key {
		t.Table = append(t.Table[:i], t.Table[i+1:]...)
		return
	}

	prev := p
	next := p.Next
	for {
		if next == nil {
			break
		}

		if *next.Key == key {
			prev.Next = next.Next
			break
		}

		prev = next
		next = next.Next
	}
}
