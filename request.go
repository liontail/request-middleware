package request

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type RequestMiddleware struct {
	Client        http.Client
	Request       *http.Request
	IsLogDuration bool
}

func NewRequestMiddleware(method, url string, body io.Reader, headers map[string]string) (*RequestMiddleware, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return &RequestMiddleware{Client: http.Client{}, Request: req}, nil
}

func (cli *RequestMiddleware) Do() (*http.Response, error) {
	now := time.Now()
	statusCode := 520
	res, err := cli.Client.Do(cli.Request)
	if cli.IsLogDuration {
		if err == nil {
			statusCode = res.StatusCode
		}
		fmt.Printf("%s | %s | %d | %s | %s\n", now.Format("2006-01-02 15:04:05"), cli.Request.Method, statusCode, time.Since(now).String(), cli.Request.URL.String())
	}

	return res, err
}

func (cli *RequestMiddleware) DoBind(value interface{}) error {
	res, err := cli.Do()
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &value)
}
