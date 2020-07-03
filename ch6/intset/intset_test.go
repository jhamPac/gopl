package main

import "testing"

func TestLenZeroInitially(t *testing.T) {
	s := &IntSet{}
	if s.Len() != 0 {
		t.Logf("%d != 0", s.Len())
		t.Fail()
	}
}

func TestLenAfterAddingElements(t *testing.T) {
	s := &IntSet{}
	s.Add(0)
	s.Add(2000)
	if s.Len() != 2 {
		t.Logf("%d != 2", s.Len())
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	s := &IntSet{}
	s.Add(0)
	s.Remove(0)
	if s.Has(0) {
		t.Log(s)
		t.Fail()
	}
}
