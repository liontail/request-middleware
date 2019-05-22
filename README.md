# request-middleware

## How to use

Example of Post Request
======

```go
import (
  ...
	"github.com/liontail/request-middleware"
)
      body := []byte("Hello World")
	    url := "your URL"
			headers := make(map[string]string)
			headers["Content-Type"] = "application/json" // set content type to json in headers

			req, err := request.NewRequestMiddleware("POST", url, bytes.NewReader(body), headers) // Create Request Method Post
			if err != nil {
          //handle if there is error
			}
			
      req.IsLogDuration = true // if you want to log duration of each request set this
			res, err := req.Do() // Do the request

```
