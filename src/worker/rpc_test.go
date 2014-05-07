package worker

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

import . "launchpad.net/gocheck"
import "code.google.com/p/go-uuid/uuid"

type DBusSuite struct{}

var _ = Suite(&DBusSuite{})

func (s *DBusSuite) TestCall(c *C) {
	socket := fmt.Sprintf("/tmp/test.%s.sock", uuid.New()[0:8])

	go Serve(socket)
	rpc := NewRPC(socket)
	rpc.SetTimeout(1 * time.Second)
	rpc.Connect()

	var result int
	rpc.Call("Arith.Add", 1, &result)

	c.Assert(result, Equals, 16)
}

func (s *DBusSuite) TestCallWithError(c *C) {
	rpc := NewRPC("/tmp/non-exists.sock")
	rpc.SetTimeout(1 * time.Microsecond)
	err := rpc.Connect()

	c.Assert(err, NotNil)
}

type Arith int

func (t *Arith) Add(value int, reply *int) error {
	*reply = value + 15
	return nil
}

func Serve(socket string) {
	l, err := net.Listen("unix", socket)
	defer l.Close()

	if err != nil {
		panic(err)
	}

	rpc.Register(new(Arith))
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
