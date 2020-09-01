package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Echo(t *testing.T) {
	names := []string{"Example as received",
		"Echo matrix 4 by 4"}
	expected := []string{"1,2,3\n4,5,6\n7,8,9\n", "9,8,7,0\n6,5,4,0\n3,2,1,0\n0,0,0,0\n"}
	for _, c := range getCasesString(names, expected){
		t.Run(c.name, func(t *testing.T) {
			service := NewEchoService()
			response := service.GetEcho(c.data)
			assert.Equal(t, c.expected, response)
		})
	}
}
