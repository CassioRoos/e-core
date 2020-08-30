package handlers

import (
	"encoding/csv"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type BaseHandler struct {}

func (b *BaseHandler) getFieldFromForm(log hclog.Logger, r *http.Request) ([][]string, error){
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Error("Unable to find required field",
			"Field", "file",
			"Error", err)
		//http.Error(rw, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Error("Error while reading CSV file", "Error", err)
		return nil, err
	}
	return records, nil
}