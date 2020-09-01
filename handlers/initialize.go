package handlers

import (
	"github.com/CassioRoos/e-core/services"
	"github.com/hashicorp/go-hclog"
	"net/http"
)
//Creates a serverMux to handle all requests
var ServerMux = http.NewServeMux()


var log = hclog.New(&hclog.LoggerOptions{
	Name:       "e-core -> ",
	Level:      hclog.LevelFromString("DEBUG"),
	JSONFormat: true,
	TimeFormat: "02/01/2006 15:04:05",
})

func init() {
	echoService := services.NewEchoService()
	sumService := services.NewSumService()
	flattenService := services.NewFlattenService()
	multiplyService := services.NewMultiplyService()
	invertService := services.NewInvertService()

	echo := NewEcho(log, echoService)
	sum := NewSum(log,sumService)
	flatten := NewFlatten(log, flattenService)
	multiply := NewMultiply(log,multiplyService)

    invert := NewInvert(log, invertService)
	healthCheck := NewHealthCheck(log)

	ServerMux.Handle("/echo", echo)
	ServerMux.Handle("/sum", sum)
	ServerMux.Handle("/multiply", multiply)
	ServerMux.Handle("/flatten", flatten)
	ServerMux.Handle("/invert", invert)
	ServerMux.Handle("/healthcheck", healthCheck)

}
