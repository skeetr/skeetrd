package server

import (
	"fmt"
	"net"
	"net/http"
	"skeetrd/intf"
	. "skeetrd/logger"
	"time"
)

import "github.com/mcuadros/go-defaults"

type HTTPConfig struct {
	Port int    `default:"8080"`
	Host string `default:""`
}

type HTTP struct {
	config         *HTTPConfig
	server         *http.Server
	listener       net.Listener
	handler        *Handler
	requestChannel intf.RequestChannel
}

func NewHTTP(config *HTTPConfig) *HTTP {
	defaults.SetDefaults(config)
	server := &HTTP{config: config}

	return server
}

func (self *HTTP) SetProcessChannel(requestChannel intf.RequestChannel) {
	self.requestChannel = requestChannel
}

func (self *HTTP) Stop() {
	err := self.listener.Close()
	if err != nil {
		Critical("Error closing port: %s", err)
	}
}

func (self *HTTP) Start() {
	self.configure()
	go self.goStart()
}

func (self *HTTP) configure() {
	self.handler = self.buildHandler()
	self.listener = self.buildListener()
	self.server = self.buildServer()
}

func (self *HTTP) buildHandler() *Handler {
	return NewHandler(self.requestChannel)
}

func (self *HTTP) buildListener() net.Listener {
	addr := fmt.Sprintf("%s:%d", self.config.Host, self.config.Port)
	Debug("Listening at %s", addr)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		Critical("Error opening port: %s", err)
	}

	return listener
}

func (self *HTTP) buildServer() *http.Server {
	return &http.Server{
		Handler:        self.handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func (self *HTTP) goStart() {
	err := self.server.Serve(self.listener)
	Critical("Error on HTTP server: %s", err)
}
