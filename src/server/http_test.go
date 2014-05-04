package server

import (
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

	http := NewHTTP(&HTTPConfig{})
	http.SetProcessChannel(requestChannel)
	http.Start()
	request := <-requestChannel
	print(request)
	http.Stop()
}
