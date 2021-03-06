// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "My fizzbuzz api": fizz TestHelpers
//
// Command:
// $ goagen
// --design=github.com/mrsinham/fb/http/design
// --out=$(GOPATH)/src/github.com/mrsinham/fb/http/goa
// --version=v1.3.0

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/mrsinham/fb/http/goa/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// BuzzFizzBadRequest runs the method Buzz of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func BuzzFizzBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.FizzController, int1 int, int2 int, limit int, string1 string, string2 string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		query["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		query["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		query["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		query["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		query["string2"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/fizz/buzz"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		prms["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		prms["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		prms["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		prms["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		prms["string2"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "FizzTest"), rw, req, prms)
	buzzCtx, _err := app.NewBuzzFizzContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Buzz(buzzCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}

	// Return results
	return rw
}

// BuzzFizzInternalServerError runs the method Buzz of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func BuzzFizzInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.FizzController, int1 int, int2 int, limit int, string1 string, string2 string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		query["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		query["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		query["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		query["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		query["string2"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/fizz/buzz"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		prms["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		prms["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		prms["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		prms["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		prms["string2"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "FizzTest"), rw, req, prms)
	buzzCtx, _err := app.NewBuzzFizzContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Buzz(buzzCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// BuzzFizzOK runs the method Buzz of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func BuzzFizzOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.FizzController, int1 int, int2 int, limit int, string1 string, string2 string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		query["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		query["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		query["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		query["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		query["string2"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/fizz/buzz"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		prms["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		prms["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		prms["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		prms["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		prms["string2"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "FizzTest"), rw, req, prms)
	buzzCtx, _err := app.NewBuzzFizzContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Buzz(buzzCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}

// BuzzCacheFizzBadRequest runs the method BuzzCache of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func BuzzCacheFizzBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.FizzController, int1 int, int2 int, limit int, string1 string, string2 string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		query["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		query["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		query["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		query["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		query["string2"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/fizz/buzz_cache"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		prms["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		prms["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		prms["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		prms["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		prms["string2"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "FizzTest"), rw, req, prms)
	buzzCacheCtx, _err := app.NewBuzzCacheFizzContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.BuzzCache(buzzCacheCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}

	// Return results
	return rw
}

// BuzzCacheFizzInternalServerError runs the method BuzzCache of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func BuzzCacheFizzInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.FizzController, int1 int, int2 int, limit int, string1 string, string2 string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		query["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		query["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		query["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		query["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		query["string2"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/fizz/buzz_cache"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		prms["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		prms["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		prms["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		prms["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		prms["string2"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "FizzTest"), rw, req, prms)
	buzzCacheCtx, _err := app.NewBuzzCacheFizzContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.BuzzCache(buzzCacheCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// BuzzCacheFizzOK runs the method BuzzCache of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func BuzzCacheFizzOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.FizzController, int1 int, int2 int, limit int, string1 string, string2 string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		query["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		query["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		query["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		query["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		query["string2"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/fizz/buzz_cache"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(int1)}
		prms["int1"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(int2)}
		prms["int2"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(limit)}
		prms["limit"] = sliceVal
	}
	{
		sliceVal := []string{string1}
		prms["string1"] = sliceVal
	}
	{
		sliceVal := []string{string2}
		prms["string2"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "FizzTest"), rw, req, prms)
	buzzCacheCtx, _err := app.NewBuzzCacheFizzContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.BuzzCache(buzzCacheCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}

// ExpireCacheFizzOK runs the method ExpireCache of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ExpireCacheFizzOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.FizzController) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/fizz/expire_cache"),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "FizzTest"), rw, req, prms)
	expireCacheCtx, _err := app.NewExpireCacheFizzContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.ExpireCache(expireCacheCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}
