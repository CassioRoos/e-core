package services

import (
	"errors"
	"fmt"
	"strconv"
)

type ProcessInt64 func(j, i int64)int64

func processInt64(records [][]string, startsin int64, f ProcessInt64) (int64,error) {
	var response int64 = startsin
	for _, row := range records {
		for _, value := range row {
			i, err := strconv.ParseInt(value,0, 64)
			if err != nil {
				return 0, errors.New(fmt.Sprintf("Error converting %s to integer", value))
			}
			response = f(response,i)
		}
	}
	return response, nil
}