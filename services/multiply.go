package services

import (
	"errors"
	"fmt"
	"strconv"
)

type MultiplyService interface {
	GetMultiplication(records [][]string) (int64, error)
}

type multiply struct {
}

func NewMultiplyService() MultiplyService {
	return &multiply{}
}

func (m *multiply) GetMultiplication(records [][]string) (int64, error) {
	var response int64 = 1
	for _, row := range records {
		for _, value := range row {
			i, err := strconv.ParseInt(value,10, 64)
			if err != nil {
				return 0, errors.New(fmt.Sprintf("Error converting %s to integer", value))
			}
			response *= i
		}
	}
	return response, nil
}
