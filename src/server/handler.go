package server

import (
	"encoding/json"
	"net/http"
	"skeetrd/intf"
	. "skeetrd/logger"
)

type Handler struct {
	requestChannel intf.RequestChannel
}

func NewHandler(requestChannel intf.RequestChannel) *Handler {
	handler := &Handler{requestChannel: requestChannel}
	return handler
}

func (self *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	self.requestChannel <- request
}
