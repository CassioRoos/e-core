package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Flatten(t *testing.T) {
	names := []string{"Should return a single line with all the values",
		"Should return a single line with all the values using a 4 by 4"}
	expected := []string{"1,2,3,4,5,6,7,8,9", "9,8,7,0,6,5,4,0,3,2,1,0,0,0,0,0"}
	for _, c := range getCasesString(names, expected){
		t.Run(c.name, func(t *testing.T) {
			service := NewFlattenService()
			response := service.GetFlatten(c.data)
			assert.Equal(t, c.expected, response)
		})
	}
}
