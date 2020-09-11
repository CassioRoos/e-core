package handlers

import (
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

//Creates a serverMux to handle all requests
var ServerMux = http.NewServeMux()

var log = hclog.New(&hclog.LoggerOptions{
	Name:  "e-core - ",
	Level: hclog.LevelFromString("DEBUG"),
})

func init() {
	// creates all services
	echoService := services.NewEchoService()
	sumService := services.NewSumService()
	flattenService := services.NewFlattenService()
	multiplyService := services.NewMultiplyService()
	invertService := services.NewInvertService()

	// inject them into handlers
	echo := MethodHttp(NewEcho(log, echoService))
	sum := MethodHttp(NewSum(log, sumService))
	flatten := MethodHttp(NewFlatten(log, flattenService))
	multiply := MethodHttp(NewMultiply(log, multiplyService))
	invert := MethodHttp(NewInvert(log, invertService))
	healthCheck := MethodHttp(NewHealthCheck(log))
	// defining the routes
	ServerMux.Handle("/echo", echo)
	ServerMux.Handle("/sum", sum)
	ServerMux.Handle("/multiply", multiply)
	ServerMux.Handle("/flatten", flatten)
	ServerMux.Handle("/invert", invert)

	// a simple handler the obtain the health of our app
	ServerMux.Handle("/healthcheck", healthCheck)

}
