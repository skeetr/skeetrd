package skeetrd

import (
	"net/http"
)

import . "launchpad.net/gocheck"

type ServerSuite struct{}

var _ = Suite(&ServerSuite{})

func (s *ServerSuite) TestLoad(c *C) {
	server := NewServer(&ServerConfig{
		Port: 1234,
		Host: "localhost",
	})

	server.Start()

	http.Get("http://localhost:1234/robots.txt")

	server.Stop()
}
