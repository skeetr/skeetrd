package skeetrd

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"time"
)

type RPC struct {
	timeout   time.Duration
	socket    string
	client    *rpc.Client
	connected bool
}

func NewRPC(socket string) *RPC {
	jsonrpc := &RPC{socket: socket}

	return jsonrpc
}

func (self *RPC) Connect() error {
	if err := self.buildClient(); err != nil {
		return err
	}

	return nil
}

func (self *RPC) SetTimeout(timeout time.Duration) {
	self.timeout = timeout
}

func (self *RPC) buildClient() error {
	if err := self.isSockectAvailable(); err != nil {
		return err
	}

	client, err := jsonrpc.Dial("unix", self.socket)
	if err != nil {
		return err
	}

	self.connected = true
	self.client = client

	return nil
}

func (self *RPC) Call(method string, args interface{}, reply interface{}) error {
	if !self.connected || self.client == nil {
		return nil
	}

	return self.client.Call(method, args, &reply)
}

func (self *RPC) isSockectAvailable() error {
	sleep := 1 * time.Microsecond
	elapsed := 0 * time.Microsecond
	for {
		if _, err := os.Stat(self.socket); err == nil {
			Debug("Socket weakup in %f seconds", elapsed.Seconds())
			return nil
		}

		if elapsed >= self.timeout {
			return fmt.Errorf("unable to find socket %s", self.socket)
		}

		time.Sleep(sleep)
		elapsed += sleep
	}

	return nil
}
