package handlers

import (
	"encoding/csv"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"strings"
)

type Flatten struct {
	log hclog.Logger
}

func NewFlatten(log hclog.Logger) *Flatten {
	return &Flatten{log: log}
}

func (e *Flatten) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}
	response = response[:len(response)-1]
	fmt.Fprint(w, response)
}
