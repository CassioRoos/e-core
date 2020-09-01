package handlers

import (
	"fmt"
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type Multiply struct {
	BaseHandler
	log     hclog.Logger
	service services.MultiplyService
}

func NewMultiply(log hclog.Logger, service services.MultiplyService) *Multiply {
	return &Multiply{log: log, service: service}
}

func (m *Multiply) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/plain")
	records, err := m.getFieldFromForm("file", m.log, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := m.service.GetMultiplication(records)
	if err != nil {
		m.log.Error("Error getting the multiplication from service", "error", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	m.log.Debug("Echo multiply success", "message", response)
	fmt.Fprint(rw, response)
}
