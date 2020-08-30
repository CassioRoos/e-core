package handlers

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"os"
	"strconv"
)

type Multiply struct {
	BaseHandler
	log hclog.Logger
}

func NewMultiply(log hclog.Logger) *Multiply {
	return &Multiply{log: log}
}

func (m *Multiply) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/plain")
	records, err := m.getFieldFromForm(m.log, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	var response float64 = 1
	for _, row := range records {
		for _, ss := range row {
			i, err := strconv.ParseFloat(ss, 64)
			if err != nil {
				m.log.Error("Error converting integer, this should not happen", "value", ss)
				os.Exit(1)
			}
			response *= i
		}
		//response = fmt.Sprintf("%m%m\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(rw, response)
}
