package stack_test

import (
	"testing"

	"github.com/nobishino/gotipcode/stack"
	"github.com/nobishino/gotipcode/assert"
)

func TestStackInt(t *testing.T) {
	type op = *int // nil means pop, non-nil means push
	push := func(v int) op {
		return op(&v)
	}
	pop := op(nil)
	testcases := [...]struct {
		name string
		ops  []op
		want []int
	}{
		{name: "1 push", ops: []op{push(1), pop}, want: []int{1}},
		{name: "2 push 1 pop", ops: []op{push(1),push(2), pop}, want: []int{2}},
		{name: "4 push 4 pop", ops: []op{push(1),push(2),push(3),pop,push(4),pop}, want: []int{3,4}},
	}

	for _, tt := range testcases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := stack.New[int]()
			var got []int
			for _, op := range tt.ops {
				if op == pop {
					got = append(got, s.Pop())
					continue
				}
				s.Push(*op)
			}
			assert.EqualSlice(t, tt.want, got)
		})
	}

	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	got := s.Pop()
	if got != 2 {
		t.Errorf("want 2, but got %d", got)
	}
}
