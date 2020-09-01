package handlers

import (
	"encoding/csv"
	"github.com/hashicorp/go-hclog"
	"net/http"
)


// A base class for all the handler, the idea is to access only by the handlers
type BaseHandler struct{}

// Method common for all the handlers, it takes the file from the request
// Returns
// [][]string - Data extracted from file as csv
// error - any error that may occur during the execution
func (b *BaseHandler) getFieldFromForm(field string, log hclog.Logger, r *http.Request) ([][]string, error) {
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
