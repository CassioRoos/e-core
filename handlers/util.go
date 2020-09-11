package handlers

import (
	"encoding/csv"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

func getFieldFromForm(field string, log hclog.Logger, r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile(field)
	if err != nil {
		log.Error("Unable to find required field",
			"Field", field,
			"Error", err)
		return nil, err
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Error("Error while reading CSV file", "Error", err)
		return nil, err
	}
	//return the matrix and no error
	return records, nil
}
