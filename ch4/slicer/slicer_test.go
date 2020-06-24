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
		t.Errorf("data %v does not equal %v", dummyData, expect)
	}
}

func TestReverseP(t *testing.T) {
	dummyData := [6]int{1, 2, 3, 4, 5, 6}
	expect := [6]int{6, 5, 4, 3, 2, 1}

	slicer.ReverseP(&dummyData)

	if !reflect.DeepEqual(dummyData, expect) {
		t.Errorf("data %v does not equal expected %v", dummyData, expect)
	}
}

func TestPureReverse(t *testing.T) {
	dummyData := []int{1, 2, 3, 4, 5, 6, 7}
	expect := []int{7, 6, 5, 4, 3, 2, 1}
	got := slicer.PureReverse(dummyData)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("data %v does not equal expected %v", got, expect)
	}
}

func TestRotateLeft(t *testing.T) {
	dummyData := []int{1, 2, 3, 4, 5}
	expect := []int{2, 3, 4, 5, 1}
	got := slicer.RotateLeft(dummyData, 1)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got %v does not equal expected %v", got, expect)
	}
}

func TestRemoveDup(t *testing.T) {
	dummyData := []string{"diamonds", "diamonds", "hearts", "hearts", "hearts"}
	expect := []string{"diamonds", "hearts"}
	got := slicer.RemoveDup(dummyData)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("duplicated strings were not removed from %v to %v", dummyData, got)
	}
}
