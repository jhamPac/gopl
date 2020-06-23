package slicer

// Reverse a slice of int
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// ReverseP reverse an array using a pointer
func ReverseP(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// PureReverse reverses an array and returns a new copy
func PureReverse(s []int) []int {
	slice := make([]int, len(s))
	for i, j := 0, len(s)-1; j >= 0; i, j = i+1, j-1 {
		slice[i] = s[j]
	}
	return slice
}
