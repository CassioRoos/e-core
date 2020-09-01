package handlers

import (
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_InvertHandler_Success(t *testing.T) {
	matrix := [][]string{{"1", "2"}, {"3", "-4"}}
	req := getRequest("GET", "/invert", "echo","file")
	service := &services.InvertServiceMock{}
	handler := &Invert{
		service: service,
		log:     hclog.New(&hclog.LoggerOptions{Level: hclog.LevelFromString("DEBUG")}),
	}
	service.On("GetInvert", matrix).Return("1,3\n2,-4\n")
	handler.ServeHTTP(req.Recorder, req.Request)
	assert.Equal(t, http.StatusOK, req.Recorder.Code)
	assert.Equal(t, "1,3\n2,-4\n", req.Recorder.Body.String())
}
