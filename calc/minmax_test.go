package calc_test

import(
	"testing"

	"github.com/nobishino/gotipcode/calc"
)

func TestMaxInt(t *testing.T) {
	testcases := [...]struct{
		x int 
		y int
		want any
	}{
		{x: 1, y:2, want:2},
	}
	for _,tt:=range testcases{
		got := calc.Max(tt.x, tt.y)
		if tt.want != got {
			t.Errorf("want %v, but got %v", tt.want, got)
		}
	}
}
