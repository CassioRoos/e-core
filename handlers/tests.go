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

func csvfile(filename string) io.Reader {
	file, err := os.Open(fmt.Sprintf("data/%s.csv", filename))
	if err != nil {
		panic(err)
	}
	return file
}

func getRequest(method, url, filename, field string) *requestObjecy{
	var requestBody bytes.Buffer
	mpw := multipart.NewWriter(&requestBody)
	defer mpw.Close()
	fw, err := mpw.CreateFormFile(field, fmt.Sprintf("%s.csv", filename))
	if err != nil {
		panic(err)
	}

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
