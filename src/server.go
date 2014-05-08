package skeetrd

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

import "github.com/mcuadros/go-defaults"

type ServerConfig struct {
	Port int    `default:"8080"`
	Host string `default:""`
}

type Server struct {
	config   *ServerConfig
	server   *http.Server
	listener net.Listener
	handler  *Handler
	pool     *Pool
}

func NewServer(config *ServerConfig) *Server {
	defaults.SetDefaults(config)
	server := &Server{config: config}

	return server
}

func (self *Server) Stop() {
	err := self.listener.Close()
	if err != nil {
		Critical("Error closing port: %s", err)
	}

	self.pool.Kill()
}

func (self *Server) Start() {
	self.configure()
	go self.goStart()
}

func (self *Server) configure() {
	self.buildPool()
	self.buildHandler()
	self.buildListener()
	self.buildServer()
}

func (self *Server) buildPool() {
	self.pool = NewPool(&PoolConfig{})
	self.pool.Start()
}

func (self *Server) buildHandler() {
	self.handler = NewHandler(self.pool)
}

func (self *Server) buildListener() {
	addr := fmt.Sprintf("%s:%d", self.config.Host, self.config.Port)
	Debug("Listening at %s", addr)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		Critical("Error opening port: %s", err)
	}

	self.listener = listener
}

func (self *Server) buildServer() {
	self.server = &http.Server{
		Handler:        self.handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func (self *Server) goStart() {
	self.server.Serve(self.listener)
}
