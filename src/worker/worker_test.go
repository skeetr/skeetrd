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
	worker.SetId("test")
	worker.Start()

	worker.Process(intf.Request("foo"))
}
