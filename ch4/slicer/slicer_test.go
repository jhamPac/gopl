package slicer_test

import (
	"reflect"
	"testing"

	"github.com/jhampac/gopl/ch4/slicer"
)

func TestReverse(t *testing.T) {
	dummyData := []int{1, 2, 3, 4, 5}
	expect := []int{5, 4, 3, 2, 1}

	slicer.Reverse(dummyData)

	if !reflect.DeepEqual(dummyData, expect) {
		t.Errorf("data %v does not equal %d", dummyData, expect)
	}
}
