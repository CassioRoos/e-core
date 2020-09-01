package handlers

import (
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"net/http"
)
//Creates a serverMux to handle all requests
var ServerMux = http.NewServeMux()


var log = hclog.New(&hclog.LoggerOptions{
	Name:       "e-core - ",
	Level:      hclog.LevelFromString("DEBUG"),
})

func init() {
	// creates all services
	echoService := services.NewEchoService()
	sumService := services.NewSumService()
	flattenService := services.NewFlattenService()
	multiplyService := services.NewMultiplyService()
	invertService := services.NewInvertService()

	// inject them into handlers
	echo := NewEcho(log, echoService)
	sum := NewSum(log,sumService)
	flatten := NewFlatten(log, flattenService)
	multiply := NewMultiply(log,multiplyService)
    invert := NewInvert(log, invertService)
	healthCheck := NewHealthCheck(log)

	// defining the routes
	ServerMux.Handle("/echo", echo)
	ServerMux.Handle("/sum", sum)
	ServerMux.Handle("/multiply", multiply)
	ServerMux.Handle("/flatten", flatten)
	ServerMux.Handle("/invert", invert)

	// a simple handler the obtain the health of our app
	ServerMux.Handle("/healthcheck", healthCheck)

}
