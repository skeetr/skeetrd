package worker

import . "launchpad.net/gocheck"

type DBusSuite struct{}

var _ = Suite(&DBusSuite{})

func (s *DBusSuite) TestCall(c *C) {
	dbus := NewDBus("org.freedesktop.DBus", "/org/freedesktop/DBus")
	dbus.Connect()

	var result []string
	dbus.Call("ListNames").Store(&result)

	c.Assert(len(result) > 0, Equals, true)
}
