package brocadevadc

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
)

const (
	endpoint = "/api/tm/3.9/config/active"
)

// The Supported HTTP Methods are outliend in the brocade-vtm-10.4-restapi.pdf document found
// at brocade.com. GET, PUT, DELETE
func (c *ClientConfig) Get(uripath string) (*http.Response, error) {
	client := http_ssl(c)
	req, err := http.NewRequest("GET", c.URL+uripath, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	return client.Do(req)

}

func (c *ClientConfig) Put(uripath string, jsonpayload *bytes.Buffer) (*http.Response, error) {
	client := http_ssl(c)
	req, err := http.NewRequest("PUT", c.URL+uripath, jsonpayload)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("content-type", "application/json")
	return client.Do(req)
}

func (c *ClientConfig) Delete(uripath string) (*http.Response, error) {
	client := http_ssl(c)
	req, err := http.NewRequest("DELETE", c.URL+uripath, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	return client.Do(req)
}

func http_ssl(c *ClientConfig) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.SslVerify,
			},
		},
	}
}

// This is just a wrapper for the JSON Encoding of structs
func jsonEncoder(j interface{}) *bytes.Buffer {
	var jsonbuffer []byte

	jsonpayload := bytes.NewBuffer(jsonbuffer)
	enc := json.NewEncoder(jsonpayload)
	enc.Encode(j)
	return jsonpayload
}
