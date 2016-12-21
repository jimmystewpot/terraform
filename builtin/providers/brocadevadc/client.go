package brocadevadc

import (
	"bytes"
	"net/http"
)

// The Supported HTTP Methods are outliend in the brocade-vtm-10.4-restapi.pdf document found
// at brocade.com. GET, PUT, DELETE
func (c *ClientConfig) Get(endpoint string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.URL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)
	return client.Do(req)

}

func (c *ClientConfig) Put(endpoint string, jsonpayload *bytes.Buffer) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", c.URL+endpoint, jsonpayload)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("content-type", "application/json")
	return client.Do(req)
}

func (c *ClientConfig) Delete(endpoint string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", c.URL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	return client.Do(req)
}
