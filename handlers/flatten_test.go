package handlers

import (
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_FlattenHandler_Success(t *testing.T) {
	matrix := [][]string{{"1", "2"}, {"3", "-4"}}
	req := getRequest("GET", "/flatten", "echo","file")
	service := &services.FlattenServiceMock{}
	handler := &Flatten{
		service: service,
		log:     hclog.New(&hclog.LoggerOptions{Level: hclog.LevelFromString("DEBUG")}),
	}
	service.On("GetFlatten", matrix).Return("1,2,3,-4\n")
	handler.ServeHTTP(req.Recorder, req.Request)
	assert.Equal(t, http.StatusOK, req.Recorder.Code)
	assert.Equal(t, "1,2,3,-4\n", req.Recorder.Body.String())
}

