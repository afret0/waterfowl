package _http

func HttpTem() string {
	t := `
package _http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

var request *Request

type Request struct {
	client *http.Client
	logger *logrus.Logger
}

func GetReq() *Request {
	if request == nil {
		request = new(Request)
		request.client = new(http.Client)
		request.logger = log.GetLogger()
	}
	return request
}

func (r *Request) Post(url string, body interface{}, headers ...http.Header) ([]byte, error) {
	hd := make(http.Header)
	if len(headers) != 0 {
		hd = headers[0]
	}
	hd.Add("Content-Type", "application/json")
	payloadJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	payload := bytes.NewReader(payloadJson)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header = hd
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		r.logger.Errorf("statusCode is %v, url: %s", resp.StatusCode, url)
	}

	defer func() {
		err = resp.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func (r *Request) MarshallUrlParams(url string, params map[string]string) string {
	l := make([]string, 0)
	for k, v := range params {
		s := fmt.Sprintf("%s=%s", k, v)
		l = append(l, s)
	}
	s := strings.Join(l, "&")
	return fmt.Sprintf("%s?%s", url, s)
}

func (r *Request) Get(url string, headers ...http.Header) ([]byte, error) {
	hd := make(http.Header)
	if len(headers) != 0 {
		hd = headers[0]
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = hd


	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}


`
	return t
}
