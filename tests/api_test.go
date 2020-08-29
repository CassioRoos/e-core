package tests

import (
	"fmt"
	"gopkg.in/h2non/baloo.v3"
	"io"
	"net/http"
	"os"
	"testing"
)

var apiClient *baloo.Client

func init() {
	host := os.Getenv("TEST_SERVER_HOST")
	port := os.Getenv("PORT")
	apiClient = baloo.New(fmt.Sprintf("%s%s", host, port))
}

func csvfile(filename string) io.Reader{
	file, err := os.Open(fmt.Sprintf("../docker/e2e/data/%s.csv",filename))
	if err != nil{
		panic(err)
	}
	return file
}

func Test_Echo(t *testing.T) {
	apiClient.Get("/echo").
		File("file", csvfile("echo")).
		Expect(t).
		Status(http.StatusOK).
		Type("text/plain").
		BodyEquals("8,9,7,6\n1,2,3,4\n5,0,5,0").
		Done()
}

func Test_Flatten(t *testing.T) {
	apiClient.Get("/flatten").
		File("file", csvfile("flatten")).
		Expect(t).
		Status(http.StatusOK).
		Type("text/plain").
		BodyEquals("9,8,7,6,5,4,3,2,1").
		Done()
}
func Test_Sum(t *testing.T) {
	apiClient.Get("/sum").
		File("file", csvfile("sum")).
		Expect(t).
		Status(http.StatusOK).
		Type("text/plain").
		BodyEquals("210").
		Done()
}
func Test_Inverter(t *testing.T) {
	apiClient.Get("/inverter").
		File("file", csvfile("inverter")).
		Expect(t).
		Status(http.StatusOK).
		Type("text/plain").
		BodyEquals("1,3\n2,4").
		Done()
}
func Test_MultiplyWithZero(t *testing.T) {
	apiClient.Get("/multiply").
		File("file", csvfile("multiply_zero")).
		Expect(t).
		Status(http.StatusOK).
		Type("text/plain").
		BodyEquals("0").
		Done()
}

func Test_Multiply(t *testing.T) {
	apiClient.Get("/multiply").
		File("file", csvfile("multiply")).
		Expect(t).
		Status(http.StatusOK).
		Type("text/plain").
		BodyEquals("725760").
		Done()
}