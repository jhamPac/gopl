package column

import "fmt"

// Person represents a person
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type comparison int

const (
	lt comparison = iota
	eq
	gt
)

type columnCmp func(a, b *Person) comparison

// ByColumns sorts by a specified column
type ByColumns struct {
	p          []Person
	columns    []columnCmp
	maxColumns int
}

// NewByColumns creates a new ByColumns type
func NewByColumns(p []Person, maxColumns int) *ByColumns {
	return &ByColumns{p, nil, maxColumns}
}
