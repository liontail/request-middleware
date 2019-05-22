package request

import (
	"fmt"
	"io"
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
		fmt.Printf("| %s | %d | %s | %s\n", cli.Request.Method, statusCode, time.Since(now).String(), cli.Request.URL.String())
	}

	return res, err
}
