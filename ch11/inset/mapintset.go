package main

// MapIntSet is a IntSet as a map data structure
type MapIntSet struct {
	m map[int]bool
}

// NewMapIntSet returns a pointer to a zero value MapIntSet
func NewMapIntSet() *MapIntSet {
	return &MapIntSet{map[int]bool{}}
}
