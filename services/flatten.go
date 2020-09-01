package services

import (
	"fmt"
	"strings"
)

type FlattenService interface {
	GetFlatten(records [][]string) string
}

type flatten struct {
}

func NewFlattenService() FlattenService {
	return &flatten{}
}

func (e *flatten) GetFlatten(records [][]string) string {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}
	response = response[:len(response)-1]
	return response
}