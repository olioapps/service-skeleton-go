package util

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	BaseURL    string
	AuthToken  string
	AuthScheme string
}

func NewHttpClient(baseURL string) *HttpClient {
	httpClient := HttpClient{}
	httpClient.BaseURL = baseURL

	return &httpClient
}

func (client *HttpClient) Post(uri string, bodyStr string) (*http.Response, string, error) {
	return client.requestWithBody("POST", uri, bodyStr, client.AuthToken)
}

func (client *HttpClient) Put(uri string, bodyStr string) (*http.Response, string, error) {
	return client.requestWithBody("PUT", uri, bodyStr, client.AuthToken)
}

func (client *HttpClient) Get(uri string) (*http.Response, string, error)  {
	return client.requestNoBody("GET", uri, client.AuthToken)
}

func (client *HttpClient) Delete(uri string) (*http.Response, string, error) {
	return client.requestNoBody("DELETE", uri, client.AuthToken)
}

func (client *HttpClient) requestWithBody(verb string, uri string, bodyStr string, authToken string) (*http.Response, string, error) {
	url := client.BaseURL + uri
	fmt.Println(verb, url)

	var body io.Reader
	if bodyStr != "" {
		body = bytes.NewBuffer([]byte(bodyStr))
	} else {
		body = nil
	}

	req, err := http.NewRequest(verb, url, body)
	var authScheme = client.AuthScheme
	if authScheme == "" {
		authScheme = "Basic"
	}
	if authToken != "" {
		req.Header.Set("Authorization", authScheme + " " + authToken)
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

func (client *HttpClient) requestNoBody(verb string, uri string, authToken string) (*http.Response, string, error) {
	return client.requestWithBody(verb, uri, "", authToken)
}
