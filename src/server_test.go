package skeetrd

import (
	"net/http"
)

import . "launchpad.net/gocheck"

type ServerSuite struct{}

var _ = Suite(&ServerSuite{})

func (s *ServerSuite) TestLoad(c *C) {
	requestChannel := make(RequestChannel, 1)

	server := NewServer(&ServerConfig{
		Port: 1234,
		Host: "localhost",
	})

	server.Start()

	http.Get("http://localhost:1234/robots.txt")
	request := <-requestChannel

	c.Assert(request.Method, Equals, "GET")
	server.Stop()
}
