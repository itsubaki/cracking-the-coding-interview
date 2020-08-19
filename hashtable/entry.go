package hashtable

import "fmt"

type Entry struct {
	Key   *string
	Value *string
	Next  *Entry
}

func (e *Entry) IsEmpty() bool {
	return e.Key == nil
}

func (e *Entry) String() string {
	return fmt.Sprintf("%s: %s", *e.Key, *e.Value)
}

func Hash(key string, capacity int) uint {
	return uint(len(key) % capacity)
}
