package skeetrd

import (
	"bytes"
	"net/http"
)

import . "launchpad.net/gocheck"

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestServeHTTP(c *C) {

	writer := &TestWriter{}
	request, _ := http.NewRequest("GET", "/foo", bytes.NewReader([]byte("")))

	pool := NewPool(&PoolConfig{})
	pool.Start()

	handler := NewHandler(pool)
	handler.ServeHTTP(writer, request)

	//	result := <-requestChannel
	//	c.Assert(result.Method, Equals, "GET")
}

type TestWriter struct {
}

func (self *TestWriter) Header() http.Header {
	return http.Header{}
}

func (self *TestWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (self *TestWriter) WriteHeader(int) {
}
