package handlers

import (
	"encoding/csv"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type Inverter struct {
	BaseHandler
	log hclog.Logger
}

func NewInverter(log hclog.Logger) *Inverter {
	return &Inverter{log: log}
}

func (t *Inverter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/plain")
	records, err := t.getFieldFromForm(t.log, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	var response string
	for i := 0; i < len(records); i++ {
		line := ""
		for j := 0; j < len(records); j++ {
			line += records[j][i] + ","
		}
		line = line[:len(line)-1]
		response += line + "\n"
	}
	fmt.Fprint(rw, response)
}
