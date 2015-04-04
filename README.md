IFifo
=====

[![GoDoc] (https://godoc.org/github.com/brentp/ififo?status.png)](https://godoc.org/github.com/brentp/fifo)


This is a fast, simple, thread-safe FIFO built with a channel. In my tests, this is faster than a Mutex
with a slice (even LIFO).
