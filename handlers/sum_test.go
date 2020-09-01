package handlers

import (
	"errors"
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_SumHandler_Success(t *testing.T) {
	matrix := [][]string{{"1", "2"}, {"3", "-4"}}
	req := getRequest("GET", "/sum", "echo","file")
	service := &services.SumServiceMock{}
	handler := &Sum{
		service: service,
		log:     hclog.New(&hclog.LoggerOptions{Level: hclog.LevelFromString("DEBUG")}),
	}
	service.On("GetSum", matrix).Return(int64(2), nil)
	handler.ServeHTTP(req.Recorder, req.Request)
	assert.Equal(t, http.StatusOK, req.Recorder.Code)
	assert.Equal(t, "2", req.Recorder.Body.String())
}

func Test_SumHandler_Error(t *testing.T) {
	matrix := [][]string{{"1", "2"}, {"3", "-4"}}
	req := getRequest("GET", "/sum", "echo","file")
	service := &services.SumServiceMock{}
	handler := &Sum{
		service: service,
		log:     hclog.New(&hclog.LoggerOptions{Level: hclog.LevelFromString("DEBUG")}),
	}
	service.On("GetSum", matrix).Return(int64(2), errors.New("Oops"))
	handler.ServeHTTP(req.Recorder, req.Request)
	assert.Equal(t, http.StatusInternalServerError, req.Recorder.Code)
	assert.Equal(t, "Oops\n", req.Recorder.Body.String())
}
