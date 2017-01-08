package brocadevadc

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"log"
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

// This is a http Client addon to verify the certificate or not based on the configuration.
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
	err := json.NewEncoder(jsonpayload).Encode(j)

	if err != nil {
		log.Printf("jsonEcoder error: %+v", err)
	}

	return jsonpayload
}

func jsonDecodeError(e error) bool {
	// If the JSON does not decode we want to know why and where.
	if e != nil {
		if serr, ok := e.(*json.UnmarshalTypeError); ok {
			line := serr.Offset
			log.Printf("GlobalSystemCreate JSON Decode at offset: %d: %+v", line, e)
		}
		return true
	} else {
		return false
	}
}
