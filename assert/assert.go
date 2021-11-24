package assert

import "testing"

func EqualSlice[T comparable](t *testing.T, want, got []T) {
	if len(want) != len(got) {
		t.Errorf("want length of %d, but got length of %d", len(want), len(got))
		return
	}
	for i:=0;i<len(want);i++ {
		if want[i] != got[i] {
			t.Errorf("want[%[1]d] %[2]v but got[%[1]d] %[3]v",
			i, want[i], got[i],
		)
		}
	}

}