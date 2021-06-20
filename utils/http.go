package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Header struct {
	ContentType   string
	Authorization string
}

type Request struct {
	Method string
	Url    string
	Header Header
	Body   []byte
}

type HttpClient struct{}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (hc *HttpClient) Get(request Request) string {
	res, _ := http.Get(request.Url)
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func (hc *HttpClient) Post(request Request) (string, error) {
	req, _ := http.NewRequest(
		request.Method,
		request.Url,
		bytes.NewReader(request.Body),
	)

	client := &http.Client{}
	req.Header.Add("Accept", request.Header.ContentType)
	req.Header.Add("Authorization", request.Header.Authorization)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}
