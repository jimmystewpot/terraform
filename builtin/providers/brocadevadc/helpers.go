package brocadevadc

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Handle the Error codes returned from the API
// 200 - Correctly Uploaded
// 201 - Created New
// 204 - No Content which is only for file uploads
func handleHttpCodes(h *http.Response) bool {
	okay := []int{200, 201, 204}
	log.Printf("handleHttpCodes status_code: %+v", h.StatusCode)
	for _, status_code := range okay {
		if h.StatusCode == status_code {
			return true
		}
	}
	return false
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
			log.Printf("JSON Decode at offset: %d: %+v", line, e)
		}
		return true
	} else {
		return false
	}
}

// this func will validate that the input is either one, two, four.
func validateDataPlaneCores(v interface{}, k string) (ws []string, errors []error) {
	var cores = []string{"one", "two", "four"}
	value := v.(string)

	// does the value exist in the slice.
	for _, c := range cores {
		if value == c {
			return
		}
	}
	errors = append(errors, fmt.Errorf("%q must be one of the following %+v", k, cores))
	return
}

// this func will validate that the input is either rfc5746, always, never or safe.
func validateSsl3Handshake(v interface{}, k string) (ws []string, errors []error) {
	var handshakes = []string{"rfc5746", "always", "never", "safe"}
	value := v.(string)

	// does the value exist in the slice.
	for _, h := range handshakes {
		if value == h {
			return
		}
	}
	errors = append(errors, fmt.Errorf("%q must be one of the following %+v", k, handshakes))
	return
}

// this func will validate that the input is either dh_1024, dh_2048, dh_3072, dh_4096.
func validateSS3diffieHellmanKl(v interface{}, k string) (ws []string, errors []error) {
	var dhkl = []string{"dh_1024", "dh_2048", "dh_3072", "dh_4096"}
	value := v.(string)

	// does the value exist in the slice.
	for _, kl := range dhkl {
		if value == kl {
			return
		}
	}
	errors = append(errors, fmt.Errorf("%q must be one of the following %+v", k, dhkl))
	return
}

// validate that the heartbeat is either unicast or multicast
func validateHeartBeatMethod(v interface{}, k string) (ws []string, errors []error) {
	var heartbeat = []string{"unicast", "multicast"}
	value := v.(string)

	for _, h := range heartbeat {
		if value == h {
			return
		}
	}
	errors = append(errors, fmt.Errorf("%q must be one of the following %+v", k, heartbeat))
	return
}
