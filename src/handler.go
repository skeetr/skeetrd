package skeetrd

import (
	"fmt"
	"net/http"
)

type Handler struct {
	pool *Pool
}

func NewHandler(pool *Pool) *Handler {
	handler := &Handler{pool: pool}
	return handler
}

func (self *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	response := self.pool.Process(request)

	fmt.Fprintf(writer, "Hello, %q", response.Body)
}
