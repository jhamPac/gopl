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

// RotateLeft rotates a slice n times provided by the second argument
func RotateLeft(slice []int, numRot int) []int {
	lenA := len(slice)
	if numRot <= 0 || numRot >= lenA {
		return slice
	}

	temp := make([]int, lenA)
	for i, j := 0, numRot; i < lenA; i, j = i+1, j+1 {
		if j == len(a) {
			j = 0
		}
		temp[i] = slice[j]
	}
	return temp
}
