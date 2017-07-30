package util

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	BaseURL string
	JWT     string
}

func NewHttpClient(baseURL string) *HttpClient {
	httpClient := HttpClient{}
	httpClient.BaseURL = baseURL

	return &httpClient
}

func (client *HttpClient) Post(uri string, bodyStr string) (*http.Response, string, error) {
	return client.requestWithBody("POST", uri, bodyStr, client.JWT)
}

func (client *HttpClient) Put(uri string, bodyStr string) (*http.Response, string, error) {
	return client.requestWithBody("PUT", uri, bodyStr, client.JWT)
}

func (client *HttpClient) Get(uri string) (*http.Response, string, error)  {
	return client.requestNoBody("GET", uri, client.JWT)
}

func (client *HttpClient) Delete(uri string) (*http.Response, string, error) {
	return client.requestNoBody("DELETE", uri, client.JWT)
}

func (client *HttpClient) requestWithBody(verb string, uri string, bodyStr string, jwt string) (*http.Response, string, error) {
	url := client.BaseURL + uri
	fmt.Println(verb, url)

	var body io.Reader
	if bodyStr != "" {
		body = bytes.NewBuffer([]byte(bodyStr))
	} else {
		body = nil
	}

	req, err := http.NewRequest(verb, url, body)
	if jwt != "" {
		req.Header.Set("Authorization", "Bearer "+jwt)
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	return resp, string(responseBody), nil
}

func (client *HttpClient) requestNoBody(verb string, uri string, jwt string) (*http.Response, string, error) {
	return client.requestWithBody(verb, uri, "", jwt)
}
