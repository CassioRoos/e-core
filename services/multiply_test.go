package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Multiply(t *testing.T) {
	expected := []int64{362880, -24, 9003737871877668864, 0}
	for _, c := range getCasesNumber(expected) {
		t.Run(c.name, func(t *testing.T) {
			service := NewMultiplyService()
			response, err := service.GetMultiplication(c.data)
			assert.Equal(t, c.err, err)
			assert.Equal(t, c.expected, response)
		})
	}
}
