package myslices

import (
	"testing"
)

type S []int

func TestEqual(t *testing.T) {
	s1 := S{1,2,3}
	s2 := S{1,2,3}

	if !Equal(s1,s2) {
		t.Errorf("expect true but got false")
	}
}

