package pool

import (
	"axia/go-axia/chain/chainable"
	"axia/go-axia/structures/pq"
	"container/heap"
)

// A Pool is a group of chainables which have not yet been gathered
// into a block
type Pool interface {
	Add(chainable.Chainable)
	Next() chainable.Chainable
	Size() int
}

// BasicPool implementation uses a simple priority queue
type BasicPool struct {
	queue pq.PriorityQueue
}

// Add to pool
func (p *BasicPool) Add(c chainable.Chainable) {
	if p == nil {
		heap.Init(&p.queue)
	}
	p.queue.Push(c)
}

// Next from pool
func (p *BasicPool) Next() chainable.Chainable {
	return p.queue.Pop().(chainable.Chainable)
}

// Size of pool
func (p *BasicPool) Size(c chainable.Chainable) int {
	return p.queue.Len()
}
