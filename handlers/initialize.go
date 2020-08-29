package handlers

import (
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
	echo := NewEcho(log)
	sum := NewSum(log)
	multiply := NewMultiply(log)
	flatten := NewFlatten(log)
	inverter := NewInverter(log)
	healthCheck := NewHealthCheck(log)

	ServerMux.Handle("/echo", echo)
	ServerMux.Handle("/sum", sum)
	ServerMux.Handle("/multiply", multiply)
	ServerMux.Handle("/flatten", flatten)
	ServerMux.Handle("/inverter", inverter)
	ServerMux.Handle("/healthcheck", healthCheck)

}
