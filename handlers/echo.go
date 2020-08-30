package handlers

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"strings"
)

type Echo struct {
	BaseHandler
	log hclog.Logger

}

func NewEcho(log hclog.Logger) *Echo {
	return &Echo{log: log}
}

func (e *Echo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/plain")
	records, err := e.getFieldFromForm(e.log, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	e.log.Debug("Echo handler success", "message", response)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, response)
}
