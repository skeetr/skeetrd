package worker

import (
	"net/http"
)

import . "launchpad.net/gocheck"

type PoolSuite struct{}

var _ = Suite(&PoolSuite{})

func (s *PoolSuite) TestCall(c *C) {

	pool := NewPool(&PoolConfig{})
	pool.Start()

	request := http.Request{Method: "GET"}
	for i := 0; i < 100; i++ {
		pool.Process(&request)
	}

	total := 0
	for _, value := range pool.GetUsage() {
		total += value
	}

	c.Assert(total, Equals, 100)
	pool.Kill()
}
