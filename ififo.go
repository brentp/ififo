// Package ififo provides a fast, sized, generic thread-safe FIFO.
// This is useful for re-using structs to reduce load on the garbarge collector
// and, unlink sync.Pool, it is longer-lived.
package ififo

// IFifo holds the queue and a default constructor when the queue is empty.
type IFifo struct {
	cache       chan interface{}
	constructor func() interface{}
}

// NewIfifo accepts a queue size and an initialization function that is called
// when the queue is empty.
func NewIFifo(max int, init func() interface{}) *IFifo {
	return &IFifo{make(chan interface{}, max), init}
}

// Get takes an item from the queue if possible or calls the constructor.
func (s *IFifo) Get() interface{} {
	select {
	case iv := <-s.cache:
		return iv
	default:
		return s.constructor()
	}
}

// Put is called to return an item to the queue.
func (s *IFifo) Put(i interface{}) {
	select {
	case s.cache <- i:
	default:
	}
}
