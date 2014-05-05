package worker

import "github.com/godbus/dbus"

type DBus struct {
	destination string
	path        dbus.ObjectPath
	connection  *dbus.Conn
	object      *dbus.Object
}

func NewDBus(destination string, path string) *DBus {
	dbus := &DBus{destination: destination, path: dbus.ObjectPath(path)}

	return dbus
}

func (self *DBus) Connect() {
	self.connection = self.buildConnection()
	self.object = self.buildObject()
}

func (self *DBus) buildConnection() *dbus.Conn {
	connection, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	return connection
}

func (self *DBus) buildObject() *dbus.Object {
	return self.connection.Object(self.destination, self.path)
}

func (self *DBus) Call(method string, args ...interface{}) *dbus.Call {
	method = self.destination + string('.') + method

	return self.object.Call(method, 0, args...)
}
