package goa

import (
	"net/http/httptest"
	"testing"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"

	"net/http"
	"net/url"

	"strconv"

	"io/ioutil"

	"strings"

	"github.com/mrsinham/fb/http/goa/app"
)

func TestHTTPFizzBuzz(t *testing.T) {
	service := goa.New("My fizzbuzz api")
	//service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	//service.Use(middleware.Recover())
	c := NewFizzController(service)
	app.MountFizzController(service, c)

	st := httptest.NewServer(service.Mux)
	defer st.Close()

	for _, v := range []struct {
		String1  string
		String2  string
		Int1     int
		Int2     int
		Limit    int
		Code     int
		Expected string
	}{
		{
			String1:  "fizz",
			String2:  "buzz",
			Int1:     3,
			Int2:     5,
			Limit:    20,
			Code:     200,
			Expected: `["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz"]`,
		},
		{
			String1:  "fizz",
			String2:  "buzz",
			Int1:     -1,
			Int2:     5,
			Limit:    20,
			Code:     400,
			Expected: `"code":"invalid_request","status":400,"detail":"int1 must be greater than or equal to 1 but got value -1","meta":{"attribute":"int1","comp":"greater than or equal to","expected":1,"value":-1}}`,
		},
	} {

		func() {
			q := make(url.Values)
			q.Set("string1", v.String1)
			q.Set("string2", v.String2)
			q.Set("int1", strconv.Itoa(v.Int1))
			q.Set("int2", strconv.Itoa(v.Int2))
			q.Set("limit", strconv.Itoa(v.Limit))
			req, err := http.NewRequest("GET", st.URL+"/fizz/buzz?"+q.Encode(), nil)
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-type", "application/json")

			r, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			defer func() {
				r.Body.Close()
			}()

			var buf []byte
			buf, err = ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}

			bodyStr := string(buf)
			// json marshal output an extra line after each encoding
			exp := v.Expected + "\n"

			if !strings.Contains(bodyStr, exp) {
				t.Fatalf(
					"with param string1 %q, string2 %q, int1 %v, int2 %v, limit %v - expected %v, received %v",
					v.String1,
					v.String2,
					v.Int1,
					v.Int2,
					v.Limit,
					exp,
					bodyStr,
				)
			}

			if v.Code != r.StatusCode {
				t.Fatalf("request received with code %v - was expecting %v", r.StatusCode, v.Code)
			}
		}()

	}
}
