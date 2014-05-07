package worker

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	. "skeetrd/logger"
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
	client, err := self.buildClient()
	if err != nil {
		Info("%s", err)
		return err
	}

	self.client = client
	return nil
}

func (self *RPC) SetTimeout(timeout time.Duration) {
	self.timeout = timeout
}

func (self *RPC) buildClient() (*rpc.Client, error) {
	if err := self.isSockectAvailable(); err != nil {
		return nil, err
	}

	client, err := jsonrpc.Dial("unix", self.socket)
	if err != nil {
		return nil, err
	}

	self.connected = true

	return client, nil
}

func (self *RPC) Call(method string, args interface{}, reply interface{}) {
	if !self.connected {
		return
	}

	err := self.client.Call(method, args, &reply)
	if err != nil {
		Error("Error: %s, callingto %s", self.socket, method)
	}
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
