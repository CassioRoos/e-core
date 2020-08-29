package handlers

import (
	"encoding/json"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

type HealthCheck struct {
	log hclog.Logger
}

func NewHealthCheck(log hclog.Logger) *HealthCheck {
	return &HealthCheck{log: log}
}

func (e *HealthCheck) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	var message = struct {
		Code int `json:"code"`
		Message string `json:"message"`
	}{http.StatusOK, "Healthy"}
	jsonEncoder := json.NewEncoder(rw)
	jsonEncoder.Encode(message)
	r.Header.Set("Content-Type", "application/json")
	rw.Header().Set("Content-Type", "application/json")
	//r.Header.Add("Content-Type", "application/json")

	//fmt.Fprint(rw,map[string]interface{}{"code":http.StatusOK, "message":"Healthy"})
}
