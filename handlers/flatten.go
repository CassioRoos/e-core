package handlers

import (
	"fmt"
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type Flatten struct {
	BaseHandler

	log     hclog.Logger
	service services.FlattenService
}

func NewFlatten(log hclog.Logger, service services.FlattenService) *Flatten {
	return &Flatten{log: log, service: service}
}

func (f *Flatten) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/plain")
	records, err := f.getFieldFromForm("file", f.log, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	response := f.service.GetFlatten(records)
	f.log.Debug("Flatten handler success", "message", response)
	fmt.Fprint(rw, response)
}
