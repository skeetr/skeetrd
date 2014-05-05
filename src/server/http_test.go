package server

import (
	"net/http"
	"skeetrd/intf"
	"testing"
)

import . "launchpad.net/gocheck"

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type HTTPSuite struct{}

var _ = Suite(&HTTPSuite{})

func (s *HTTPSuite) TestLoad(c *C) {
	requestChannel := make(intf.RequestChannel, 1)

	server := NewHTTP(&HTTPConfig{
		Port: 1234,
		Host: "localhost",
	})

	server.SetProcessChannel(requestChannel)
	server.Start()

	http.Get("http://localhost:1234/robots.txt")
	request := <-requestChannel

	c.Assert(string(request), HasLen, 490)
	server.Stop()
}
