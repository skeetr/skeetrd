package worker

import (
	"fmt"
	"skeetrd/intf"
	. "skeetrd/logger"
	"strings"
)

//import "github.com/mcuadros/go-command"
import "github.com/mcuadros/go-defaults"

const (
	ProcessMethod = "process"
)

type WorkerConfig struct {
	DBusDestination string `default:"com.github.skeetr.%s"`
}

type Worker struct {
	id     string
	config *WorkerConfig
	dbus   *DBus
}

func NewWorker(config *WorkerConfig) *Worker {
	defaults.SetDefaults(config)

	worker := &Worker{config: config}
	return worker
}

func (self *Worker) SetId(id string) {
	self.id = fmt.Sprintf(self.config.DBusDestination, id)
}

func (self *Worker) Start() {
	self.buildAndConnectDBus()

	Info("New worker %s started", self.id)
}

func (self *Worker) buildAndConnectDBus() {
	path := "/" + strings.Replace(self.id, ".", "/", -1)
	Info("%s %s at %s", path, self.id, ProcessMethod)

	self.dbus = NewDBus(self.id, path)
	self.dbus.Connect()
}

func (self *Worker) Process(request intf.Request) {
	var response string
	self.dbus.Call(ProcessMethod, string(request)).Store(&response)

	Debug("Response: %s", response)
}
