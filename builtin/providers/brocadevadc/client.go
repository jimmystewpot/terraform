package brocadevadc

import (
	"bytes"
	"net/http"
)

const (
	apipath = "/api/tm/3.10/config/active"
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
