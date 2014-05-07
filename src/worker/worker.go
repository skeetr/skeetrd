package worker

import (
	"fmt"
	"net/http"
	"skeetrd/intf"
	. "skeetrd/logger"
	"time"
)

import "github.com/mcuadros/go-command"
import "github.com/mcuadros/go-defaults"
import "code.google.com/p/go-uuid/uuid"

const (
	ProcessMethod = "process"
	KillMethod    = "kill"
)

type WorkerConfig struct {
	Script        string `default:"/usr/local/bin/php /tmp/client.php %s"`
	SocketPattern string `default:"/tmp/skeetrd.%s.sock"`
}

type Worker struct {
	id      string
	config  *WorkerConfig
	rpc     *RPC
	command *command.Command
}

func NewWorker(config *WorkerConfig) *Worker {
	defaults.SetDefaults(config)

	worker := &Worker{config: config}
	return worker
}

func (self *Worker) SetId(id string) {
	self.id = id
}

func (self *Worker) GetId() string {
	return self.id
}

func (self *Worker) Start() error {
	self.generateIdIfNeeded()

	if err := self.buildAndRunCommand(); err != nil {
		return err
	}

	if err := self.buildAndConnectRPC(); err != nil {
		return err
	}

	Info("New worker %s started", self.id)
	return nil
}

func (self *Worker) generateIdIfNeeded() {
	if self.id == "" {
		self.id = uuid.New()[0:8]
	}
}

func (self *Worker) buildAndConnectRPC() error {
	socket := fmt.Sprintf(self.config.SocketPattern, self.id)

	self.rpc = NewRPC(socket)
	self.rpc.SetTimeout(1 * time.Second)

	return self.rpc.Connect()
}

func (self *Worker) buildAndRunCommand() error {
	cmd := fmt.Sprintf(self.config.Script, self.id)
	self.command = command.NewCommand(cmd)

	if err := self.command.Run(); err != nil {
		return fmt.Errorf("Error %s, executing: %s", err, cmd)
	}

	go self.goWaitCommaint()
	Debug("Running %s", cmd)

	return nil
}

func (self *Worker) goWaitCommaint() {
	if err := self.command.Wait(); err != nil {
		panic(err)
	}
}

func (self *Worker) Process(request *http.Request) *intf.Response {
	var response intf.Response
	self.rpc.Call(ProcessMethod, request, &response)

	Debug("[%s] Response: %s", ProcessMethod, response)

	return &response
}

func (self *Worker) Kill() {
	var response bool
	self.rpc.Call(KillMethod, nil, &response)

	Debug("[%s] Response: %s", KillMethod, response)
}
