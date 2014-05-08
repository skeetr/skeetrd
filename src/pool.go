package skeetrd

import (
	"net/http"
)

import "github.com/mcuadros/go-defaults"

type workerRequest struct {
	httpRequest  *http.Request
	responseChan chan *Response
}

type PoolConfig struct {
	WorkerProcesses int `default:"4"`
	WorkerConfig
}

type Pool struct {
	config  *PoolConfig
	workers map[string]*Worker
	usage   map[string]int
	channel chan workerRequest
}

func NewPool(config *PoolConfig) *Pool {
	defaults.SetDefaults(config)

	pool := &Pool{config: config}
	return pool
}

func (self *Pool) Start() error {
	self.buildChannel()
	self.buildWorkers()

	return nil
}

func (self *Pool) buildChannel() {
	self.channel = make(chan workerRequest, self.config.WorkerProcesses)
}

func (self *Pool) buildWorkers() {
	self.workers = make(map[string]*Worker, self.config.WorkerProcesses)
	self.usage = make(map[string]int, self.config.WorkerProcesses)

	for i := 0; i < self.config.WorkerProcesses; i++ {
		worker := NewWorker(&WorkerConfig{})
		worker.Start()

		self.workers[worker.GetId()] = worker
		go self.waitForRequest(worker)
	}
}

func (self *Pool) waitForRequest(worker *Worker) {
	for workerRequest := range self.channel {
		self.usage[worker.GetId()]++

		response := worker.Process(workerRequest.httpRequest)
		workerRequest.responseChan <- response
	}
}

func (self *Pool) Process(request *http.Request) *Response {
	responseChan := make(chan *Response, 1)

	self.channel <- workerRequest{
		httpRequest:  request,
		responseChan: responseChan,
	}

	return <-responseChan
}

func (self *Pool) Kill() {
	for _, worker := range self.workers {
		worker.Kill()
	}
}

func (self *Pool) GetUsage() map[string]int {
	return self.usage
}
