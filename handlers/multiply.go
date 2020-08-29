package handlers

import (
	"encoding/csv"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"os"
	"strconv"
)

type Multiply struct {
	log hclog.Logger
}

func NewMultiply(log hclog.Logger) *Multiply {
	return &Multiply{log: log}
}

func (s *Multiply) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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
	var response float64 = 1
	for _, row := range records {
		for _, ss := range row {
			i, err := strconv.ParseFloat(ss, 64)
			if err != nil {
				s.log.Error("Error converting integer, this should not happen", "value", ss)
				os.Exit(1)
			}
			response *= i
		}
		//response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(rw, response)
}
