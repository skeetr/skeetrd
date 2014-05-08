package skeetrd

import . "launchpad.net/gocheck"

type LoggerSuite struct{}

var _ = Suite(&LoggerSuite{})

func (s *LoggerSuite) TestGetRecord(c *C) {

	Warning("foo")
	Info("foo")
}
