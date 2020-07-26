package main

// IntSet bit-vector and map-based interface
type IntSet interface {
	Has(x int) bool
	Add(x int)
	AddAll(nums ...int)
	UnionWith(t IntSet)
	Len() int
	Remove(x int)
	Clear()
	Copy() IntSet
	String() string
	Ints() []int
}
