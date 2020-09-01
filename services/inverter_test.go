package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Invert(t *testing.T) {
	names := []string{"Should print the matrix inverted",
		"Should print the matrix inverted with a 4 by 4"}
	expected := []string{"1,4,7\n2,5,8\n3,6,9\n", "9,6,3,0\n8,5,2,0\n7,4,1,0\n0,0,0,0\n"}
	for _, c := range getCasesString(names, expected){
		t.Run(c.name, func(t *testing.T) {
			service := NewInvertService()
			response := service.GetInvert(c.data)
			assert.Equal(t, c.expected, response)
		})
	}
}