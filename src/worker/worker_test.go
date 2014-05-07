package worker

import (
	"skeetrd/intf"
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

	request := intf.Request{
		Method: "GET",
		Host:   "example.org",
		Header: map[string][]string{
			"Accept-Encoding": {"gzip, deflate"},
			"Accept-Language": {"en-us"},
			"Connection":      {"keep-alive"},
		},
	}

	worker.Process(&request)

}

func BenchmarkBigLen(b *testing.B) {
	worker := NewWorker(&WorkerConfig{})
	worker.Start()

	request := intf.Request{
		Method: "GET",
		Host:   "example.org",
		Header: map[string][]string{
			"Accept-Encoding": {"gzip, deflate"},
			"Accept-Language": {"en-us"},
			"Connection":      {"keep-alive"},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		worker.Process(&request)
	}
}
