package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/franela/goblin"
	request "github.com/liontail/request-middleware"
)

type Response struct {
	Message string `json:"msg"`
}

func Test(t *testing.T) {
	g := Goblin(t)
	mockResponse := `{"msg": "Hello World"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, mockResponse)
	}))
	defer ts.Close()

	g.Describe("DoBind", func() {
		g.It("Should Make Request and Bind Response", func() {
			req, err := request.NewRequestMiddleware("GET", ts.URL, nil, nil)
			if err != nil {
				t.Error(err)
			}
			v := Response{}
			if err := req.DoBind(&v); err != nil {
				t.Error(err)
			}
			g.Assert(v.Message).Equal("Hello World")
		})
	})
}
