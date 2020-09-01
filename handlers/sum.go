package handlers

import (
	"fmt"
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type Sum struct {
	BaseHandler
	service services.SumService
	log     hclog.Logger
}

func NewSum(log hclog.Logger, service services.SumService) *Sum {
	return &Sum{log: log, service: service}
}

func (s *Sum) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/plain")
	records, err := s.getFieldFromForm("file", s.log, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := s.service.GetSum(records)
	if err != nil {
		s.log.Error("Error getting the sum from service", "error", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	s.log.Debug("Echo sum success", "message", response)
	fmt.Fprint(rw, response)
}
