package server

import (
	"net/http"
	"skeetrd/intf"
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
