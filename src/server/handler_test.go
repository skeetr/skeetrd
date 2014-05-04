package server

import (
	"bytes"
	"net/http"
	"skeetrd/intf"
)

import . "launchpad.net/gocheck"

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestServeHTTP(c *C) {
	requestChannel := make(intf.RequestChannel, 1)

	writer := &TestWriter{}
	request, _ := http.NewRequest("GET", "/foo", bytes.NewReader([]byte("")))

	handler := NewHandler(requestChannel)
	handler.ServeHTTP(writer, request)

	result := <-requestChannel
	c.Assert(string(result), HasLen, 367)
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
