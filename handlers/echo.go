package handlers

import (
	"encoding/csv"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"strings"
)

type Echo struct {
	log hclog.Logger
}

func NewEcho(log hclog.Logger) *Echo {
	return &Echo{log: log}
}

func (e *Echo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	rw.Header().Set("Content-Type", "text/plain")
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(rw, response)
}
