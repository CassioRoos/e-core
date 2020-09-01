package handlers

import (
	"fmt"
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type Invert struct {
	BaseHandler
	log     hclog.Logger
	service services.InvertService
}

func NewInvert(log hclog.Logger, service services.InvertService) *Invert {
	return &Invert{log: log, service: service}
}

func (t *Invert) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/plain")
	records, err := t.getFieldFromForm("file", t.log, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	response := t.service.GetInvert(records)
	t.log.Debug("Echo invert success", "message", response)
	fmt.Fprint(rw, response)
}
