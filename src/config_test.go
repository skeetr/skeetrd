package skeetrd

import (
	"testing"
)

import . "launchpad.net/gocheck"

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type CofigSuite struct{}

var _ = Suite(&CofigSuite{})

func (s *CofigSuite) TestLoad(c *C) {
	var raw = string(`
		[section]
		level = debug
	`)

	GetConfig().Load(raw)
	c.Assert(GetConfig().Section.Level, Equals, "debug")
}
