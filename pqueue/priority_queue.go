package pqueue

import "container/heap"

type PriorityQueue[T any] struct {
	q *priorityQueue[T]
}

func New[T any](less func(x, y T) bool) PriorityQueue[T] {
	q := new[T](less)
	return PriorityQueue[T]{
		q: &q,
	}
}

func (pq *PriorityQueue[T]) Enqueue(item T) {
	heap.Push(pq.q, &item)
}

func (pq *PriorityQueue[T]) Dequeue() T {
	item := heap.Pop(pq.q)
	v := item.(*T)
	return *v
}

type priorityQueue[T any] struct {
	items []*T
	less  func(x, y T) bool
}

func new[T any](less func(x, y T) bool) priorityQueue[T] {
	return priorityQueue[T]{
		less: less,
	}
}

func (pq *priorityQueue[T]) Len() int {
	return len(pq.items)
}

func (pq *priorityQueue[T]) Less(i, j int) bool {
	return pq.less(*pq.items[i], *pq.items[j])
}

func (pq *priorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *priorityQueue[T]) Push(item interface{}) {
	pq.items = append(pq.items, item.(*T))
}

func (pq *priorityQueue[T]) Pop() interface{} {
	n := len(pq.items)
	item := pq.items[n-1]
	pq.items[n-1] = nil
	pq.items = pq.items[:n-1]
	return item
}
