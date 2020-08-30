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

//Just to make sure the application is running
func (e *HealthCheck) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	var message = SimpleResponse{http.StatusOK, "Healthy"}
	jsonEncoder := json.NewEncoder(rw)
	jsonEncoder.Encode(message)
}
