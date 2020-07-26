package main

// BitIntSet is a set of small non-negative integers. Its zero value represents the empty set
type BitIntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x
func (s *BitIntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (s *BitIntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll non-negative values to the set
func (s *BitIntSet) AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
}

// UnionWith sets s to the union of s and t
func (s *BitIntSet) UnionWith(t IntSet) {
	if bis, ok := t.(*BitIntSet); ok {
		for i, tword := range bis.words {
			if i < len(s.words) {
				s.words[i] |= tword
			} else {
				s.words = append(s.words, tword)
			}
		}
	} else {
		for _, i := range t.Ints() {
			s.Add(i)
		}
	}
}

func popcount(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

// Len returns the number of elements
func (s *BitIntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popcount(word)
	}
	return count
}
