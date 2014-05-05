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
	json := self.marshalToJSON(request)
	self.requestChannel <- json
}

func (self *Handler) marshalToJSON(request *http.Request) intf.Request {
	json, err := json.Marshal(request)
	if err != nil {
		Warning("Unable to marshal request to JSON: %s", err)
		return []byte("")
	}

	return json
}
