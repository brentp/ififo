package ififo

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type ISuite struct{}

var _ = Suite(&ISuite{})

func (s *ISuite) TestNewFifo(c *C) {
	f := NewIFifo(100, func() interface{} { return 22 })
	c.Assert(cap(f.cache), Equals, 100)
	c.Assert(len(f.cache), Equals, 0)
	c.Assert(f.New(), Equals, 22)
}

func (s *ISuite) TestFifoPut(c *C) {
	f := NewIFifo(100, func() interface{} { return 22 })
	f.Put(200)
	c.Assert(cap(f.cache), Equals, 100)
	c.Assert(len(f.cache), Equals, 1)
}

func (s *ISuite) TestFifoGet(c *C) {
	f := NewIFifo(100, func() interface{} { return 22 })
	f.Put(201)
	c.Assert(len(f.cache), Equals, 1)
	c.Assert(201, Equals, f.Get())
	c.Assert(len(f.cache), Equals, 0)
}

func (s *ISuite) TestFifoPutFull(c *C) {
	f := NewIFifo(1, func() interface{} { return 22 })
	f.Put(2)
	f.Put(3)
	c.Assert(len(f.cache), Equals, 1)
	c.Assert(2, Equals, f.Get())
	c.Assert(len(f.cache), Equals, 0)
}

func (s *ISuite) TestFifoGetEmpty(c *C) {
	f := NewIFifo(1, func() interface{} { return 22 })
	c.Assert(22, Equals, f.Get())
	c.Assert(22, Equals, f.Get())
}
