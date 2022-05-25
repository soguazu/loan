package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client model request
type Client struct {
	BaseURL string
	Header  map[string]string
}

// GET method for carrying out get request
func (r *Client) GET(method, link string, params map[string]string) ([]byte, error) {
	var (
		req *http.Request
		err error
	)

	endpoint := fmt.Sprintf("%v/%v", r.BaseURL, link)

	if len(params) > 0 {
		q := url.Values{}
		for _, param := range params {
			q.Add(param, params[param])
		}
		req, err = http.NewRequest(method, endpoint, strings.NewReader(q.Encode()))
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
	}

	if err != nil {
		return nil, err
	}

	if len(r.Header) != 0 {
		r.SetHeader(req)
	}

	client := http.Client{Timeout: time.Duration(10) * time.Second}

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// CHANGE method for carrying out get request
func (r *Client) CHANGE(method, url string, payload []byte) ([]byte, error) {
	endpoint := fmt.Sprintf("%v/%v", r.BaseURL, url)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	if len(r.Header) > 0 {
		r.SetHeader(req)
	}

	client := http.Client{Timeout: time.Duration(30) * time.Second}

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// SetHeader method for setting header
func (r *Client) SetHeader(request *http.Request) {
	for key, header := range r.Header {
		request.Header.Add(key, header)
	}
	return
}
