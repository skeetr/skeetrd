package worker

import (
	"skeetrd/intf"
	"testing"
)

import . "launchpad.net/gocheck"
import "code.google.com/p/go-uuid/uuid"

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type WorkerSuite struct{}

var _ = Suite(&WorkerSuite{})

func (s *WorkerSuite) TestLoad(c *C) {
	worker := NewWorker(&WorkerConfig{})
	worker.SetId(uuid.New()[0:8])
	worker.Start()

	worker.Process(intf.Request("foo"))

}

func BenchmarkBigLen(b *testing.B) {
	worker := NewWorker(&WorkerConfig{})
	worker.SetId(uuid.New()[0:8])
	worker.Start()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		worker.Process(intf.Request("foo"))
	}
}
