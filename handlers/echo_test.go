package handlers

import (
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)


func Test_EchoHandler_Success(t *testing.T) {
	matrix := [][]string{{"1", "2"}, {"3", "-4"}}
	req := getRequest("GET", "/echo", "echo","file")
	service := &services.EchoServiceMock{}
	handler := &Echo{
		service: service,
		log:     hclog.New(&hclog.LoggerOptions{Level: hclog.LevelFromString("DEBUG")}),
	}
	service.On("GetEcho", matrix).Return("1,2\n3,-4\n")
	handler.ServeHTTP(req.Recorder, req.Request)
	assert.Equal(t, http.StatusOK, req.Recorder.Code)
	assert.Equal(t, "1,2\n3,-4\n", req.Recorder.Body.String())
}

func Test_EchoHandler_Malformed(t *testing.T) {
	req := getRequest("GET", "/echo", "malformed","file")
	service := &services.EchoServiceMock{}
	handler := &Echo{
		service: service,
		log:     hclog.New(&hclog.LoggerOptions{Level: hclog.LevelFromString("DEBUG")}),
	}
	handler.ServeHTTP(req.Recorder, req.Request)
	assert.Equal(t, http.StatusBadRequest, req.Recorder.Code)
	assert.Equal(t, "record on line 2: wrong number of fields\n", req.Recorder.Body.String())
}

func Test_EchoHandler_NonexistField(t *testing.T) {
	req := getRequest("GET", "/echo", "echo","X")
	service := &services.EchoServiceMock{}
	handler := &Echo{
		service: service,
		log:     hclog.New(&hclog.LoggerOptions{Level: hclog.LevelFromString("DEBUG")}),
	}
	handler.ServeHTTP(req.Recorder, req.Request)
	assert.Equal(t, http.StatusBadRequest, req.Recorder.Code)
	assert.Equal(t, "http: no such file\n", req.Recorder.Body.String())
}