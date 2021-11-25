package set_test

import (
	"testing"
	
	"github.com/nobishino/gotipcode/set"
)

func TestSet(t *testing.T) {
	s := set.New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	got := s.Includes(1)
	if !got {
		t.Errorf("expect %t, but got %t", true, got)
	}

	s.Remove(1)

	got = s.Includes(1)
	if got {
		t.Errorf("expect %t, but got %t", false, got)
	}
}