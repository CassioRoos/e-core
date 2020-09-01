package services

import (
	"errors"
	"fmt"
	"strconv"
)

type SumService interface {
	GetSum(records [][]string) (int64,error)
}

type sum struct {
}

func NewSumService() SumService {
	return &sum{}
}

func (s *sum) GetSum(records [][]string) (int64,error) {
	var response int64
	for _, row := range records {
		for _, value := range row {
			i, err := strconv.ParseInt(value,0, 64)
			if err != nil {
				return 0, errors.New(fmt.Sprintf("Error converting %s to integer", value))
			}
			response += i
		}
	}
	return response, nil
}