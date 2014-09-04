package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func OsRequest(method, url string, body, out interface{}, token string) (err error) {
	var req *http.Request
	if body != nil {
		bodyReader := &bytes.Buffer{}
		err = json.NewEncoder(bodyReader).Encode(body)
		if err != nil {
			return
		}
		req, err = http.NewRequest(method, url, bodyReader)
		if err != nil {
			return
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return
		}
	}
	req.Header.Add("Content-Type", "application/json")
	if token != "" {
		req.Header.Add("X-Auth-Token", token)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	if resp.Status[0] == '2' {
		if out != nil {
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(out)
		}
		return
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = errors.New(resp.Status + ": " + string(bodyBytes))
	return
}
