package worker

import (
	"bytes"
	"net/http"

	"testing"
)

import . "launchpad.net/gocheck"

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type WorkerSuite struct{}

var _ = Suite(&WorkerSuite{})

func (s *WorkerSuite) TestLoad(c *C) {
	worker := NewWorker(&WorkerConfig{})
	worker.Start()

	request := http.Request{
		Method: "GET",
		Host:   "example.org",
		Header: map[string][]string{
			"Accept-Encoding": {"gzip, deflate"},
			"Accept-Language": {"en-us"},
			"Connection":      {"keep-alive"},
		},
	}

	result := worker.Process(&request)
	c.Assert(result.String(), Equals, "GET")

}

func BenchmarkBigLen(b *testing.B) {
	worker := NewWorker(&WorkerConfig{})
	worker.Start()

	request, _ := http.NewRequest("GET", "/foo", bytes.NewReader([]byte("")))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		worker.Process(request)
	}
}
