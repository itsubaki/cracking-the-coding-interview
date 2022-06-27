package hashtable

import "fmt"

type Entry struct {
	Key   *string
	Value *string
	Next  *Entry
}

func (e *Entry) String() string {
	return fmt.Sprintf("%s: %s", *e.Key, *e.Value)
}
