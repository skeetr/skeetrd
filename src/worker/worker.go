package worker

import (
	"fmt"
	"skeetrd/intf"
	. "skeetrd/logger"
	"time"
)

import "github.com/mcuadros/go-command"
import "github.com/mcuadros/go-defaults"
import "code.google.com/p/go-uuid/uuid"

const (
	ProcessMethod = "process"
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

func (self *Worker) Start() {
	self.generateIdIfNeeded()
	self.buildAndRunCommand()
	self.buildAndConnectRPC()
	Info("New worker %s started", self.id)
}

func (self *Worker) generateIdIfNeeded() {
	if self.id == "" {
		self.id = uuid.New()[0:8]
	}
}

func (self *Worker) buildAndConnectRPC() {
	socket := fmt.Sprintf(self.config.SocketPattern, self.id)

	self.rpc = NewRPC(socket)
	self.rpc.SetTimeout(1 * time.Second)
	self.rpc.Connect()
}

func (self *Worker) buildAndRunCommand() {
	cmd := fmt.Sprintf(self.config.Script, self.id)
	self.command = command.NewCommand(cmd)

	if err := self.command.Run(); err != nil {
		Error("Error %s, executing: %s", err, cmd)
	}

	go self.goWaitCommaint()
	Debug("Running %s", cmd)
}

func (self *Worker) goWaitCommaint() {
	if err := self.command.Wait(); err != nil {
		panic(err)
	}
}

func (self *Worker) Process(request *intf.Request) {
	var response string
	Debug("Calling method: %s", ProcessMethod)

	self.rpc.Call(ProcessMethod, request, &response)

	Debug("Response: %s", response)
}
