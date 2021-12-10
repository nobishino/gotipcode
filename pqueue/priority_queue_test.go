package pqueue_test

import (
	"fmt"
	"testing"

	"github.com/nobishino/gotipcode/pqueue"
)

func ExamplePriorityQueue() {
	pq := pqueue.New(func(x, y int) bool {
		return x < y
	})

	pq.Enqueue(2)
	pq.Enqueue(3)
	pq.Enqueue(1)
	pq.Enqueue(5)
	pq.Enqueue(4)

	for i := 1; i <= 5; i++ {
		v := pq.Dequeue()
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExamplePriorityQueueString() {
	pq := pqueue.New(func(x, y string) bool {
		return x < y
	})

	pq.Enqueue("hello")
	pq.Enqueue("gopher")
	fmt.Println(pq.Dequeue())

	pq.Enqueue("apple")
	fmt.Println(pq.Dequeue())

	pq.Enqueue("world")
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Dequeue())
	// Output:
	// gopher
	// apple
	// hello
	// world
}

func TestPriorityQueueInt(t *testing.T) {
	testcases := []struct {
		name  string
		input []int
		less  func(x, y int) bool
		want  []int
	}{
		{name: "asc", input: []int{2, 4, 3, 6, 1}, less: func(x, y int) bool { return x < y }, want: []int{1, 2, 3, 4, 6}},
		{name: "desc", input: []int{2, 4, 3, 6, 1}, less: func(x, y int) bool { return x > y }, want: []int{6, 4, 3, 2, 1}},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			pq := pqueue.New(tt.less)

			for _, in := range tt.input {
				pq.Enqueue(in)
			}

			for _, want := range tt.want {
				got := pq.Dequeue()
				if want != got {
					t.Errorf("want %d, but got %d", want, got)
				}
			}
		})
	}
}
