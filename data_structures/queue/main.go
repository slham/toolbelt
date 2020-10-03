package main

import "log"

type Queue struct {
	items []int
}

// Enqueue
func (q *Queue) Enqueue(in int) {
	q.items = append(q.items, in)
}

// Dequeue
func (q *Queue) Dequeue() int {
	out := q.items[0]
	q.items = q.items[1:]
	return out
}

func main() {
	q := Queue{}
	log.Println(q)
	q.Enqueue(1)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	log.Println(q)
	out := q.Dequeue()
	log.Println(q, out)
}
