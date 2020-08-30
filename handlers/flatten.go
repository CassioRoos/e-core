package handlers

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"strings"
)

type Flatten struct {
	BaseHandler
	log hclog.Logger
}

func NewFlatten(log hclog.Logger) *Flatten {
	return &Flatten{log: log}
}

func (f *Flatten) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/plain")
	records, err := f.getFieldFromForm(f.log, r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}
	response = response[:len(response)-1]
	fmt.Fprint(rw, response)
}
