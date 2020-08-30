package handlers

import (
	"encoding/csv"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"os"
	"strconv"
)

type Sum struct {
	BaseHandler
	log hclog.Logger
}

func NewSum(log hclog.Logger) *Sum {
	return &Sum{log: log}
}

func (s *Sum) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/plain")
	records, err := s.getFieldFromForm(s.log, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	var response int
	for _, row := range records {
		for _, value := range row {
			i, err := strconv.Atoi(value)
			if err != nil {
				s.log.Error("Error converting integer, this should not happen", "value", ss)
				os.Exit(1)
			}
			response += i
		}
	}
	fmt.Fprint(rw, response)
}
