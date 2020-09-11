package handlers

import (
	"fmt"
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type Echo struct {
	service services.EchoService
	log     hclog.Logger
}

func NewEcho(log hclog.Logger, service services.EchoService) *Echo {
	return &Echo{log: log, service: service}
}

func (e *Echo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/plain")
	records, err := getFieldFromForm("file", e.log, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	response := e.service.GetEcho(records)
	e.log.Debug("Echo handler success", "message", response)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, response)
}
