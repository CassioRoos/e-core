package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sum(t *testing.T) {
	expected := []int64{45, 2, 666, 8}
	for _, c := range getCasesNumber(expected) {
		t.Run(c.name, func(t *testing.T) {
			service := NewSumService()
			response, err := service.GetSum(c.data)
			assert.Equal(t, c.err, err)
			assert.Equal(t, c.expected, response)
		})
	}
}
