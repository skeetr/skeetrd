package intf

import (
	"net/http"
)

type RequestChannel chan *http.Request
