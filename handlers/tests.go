package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
)

type requestObjecy struct {
	Request  *http.Request
	Recorder *httptest.ResponseRecorder
}

// get the file as a io.reader
func csvfile(filename string) io.Reader {
	file, err := os.Open(fmt.Sprintf("data/%s.csv", filename))
	if err != nil {
		panic(err)
	}
	return file
}

// return the object with the request test
func getRequest(method, url, filename, field string) *requestObjecy{
	var requestBody bytes.Buffer
	// creates a multipar writer, we need to set the file to test properly
	mpw := multipart.NewWriter(&requestBody)
	// ensures that the file will close
	defer mpw.Close()

	// defines the name of the file
	fw, err := mpw.CreateFormFile(field, fmt.Sprintf("%s.csv", filename))
	if err != nil {
		panic(err)
	}

	// copy the content to the file inside our request
	_, err = io.Copy(fw, csvfile(filename))
	if err != nil {
		panic(err)
	}
	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(method, url, &requestBody)
	request.Header.Add("Content-Type", mpw.FormDataContentType())
	return &requestObjecy{
		Request: request,
		Recorder: recorder,
	}
}
