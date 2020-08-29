package handlers

import (
	"encoding/csv"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type Inverter struct {
	log hclog.Logger
}

func NewInverter(log hclog.Logger) *Inverter {
	return &Inverter{log: log}
}

func (t *Inverter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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
