package services

import (
	"fmt"
	"strings"
)

type EchoService interface {
	GetEcho(records [][]string) string
}

type echo struct {}

func NewEchoService() EchoService {
	return &echo{}
}

func (e *echo) GetEcho(records [][]string) string {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}
