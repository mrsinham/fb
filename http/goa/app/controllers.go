// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "My fizzbuzz api": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/mrsinham/fb/http/design
// --out=$(GOPATH)/src/github.com/mrsinham/fb/http/goa
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// FizzController is the controller interface for the Fizz actions.
type FizzController interface {
	goa.Muxer
	Buzz(*BuzzFizzContext) error
	BuzzCache(*BuzzCacheFizzContext) error
	ExpireCache(*ExpireCacheFizzContext) error
}

// MountFizzController "mounts" a Fizz resource controller on the given service.
func MountFizzController(service *goa.Service, ctrl FizzController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewBuzzFizzContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Buzz(rctx)
	}
	service.Mux.Handle("GET", "/fizz/buzz", ctrl.MuxHandler("Buzz", h, nil))
	service.LogInfo("mount", "ctrl", "Fizz", "action", "Buzz", "route", "GET /fizz/buzz")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewBuzzCacheFizzContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.BuzzCache(rctx)
	}
	service.Mux.Handle("GET", "/fizz/buzz_cache", ctrl.MuxHandler("BuzzCache", h, nil))
	service.LogInfo("mount", "ctrl", "Fizz", "action", "BuzzCache", "route", "GET /fizz/buzz_cache")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewExpireCacheFizzContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ExpireCache(rctx)
	}
	service.Mux.Handle("GET", "/fizz/expire_cache", ctrl.MuxHandler("ExpireCache", h, nil))
	service.LogInfo("mount", "ctrl", "Fizz", "action", "ExpireCache", "route", "GET /fizz/expire_cache")
}

// SpecController is the controller interface for the Spec actions.
type SpecController interface {
	goa.Muxer
	goa.FileServer
}

// MountSpecController "mounts" a Spec resource controller on the given service.
func MountSpecController(service *goa.Service, ctrl SpecController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSpecOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swagger.yaml", ctrl.MuxHandler("preflight", handleSpecOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "http/goa/swagger/swagger.json")
	h = handleSpecOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Spec", "files", "http/goa/swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger.yaml", "http/goa/swagger/swagger.yaml")
	h = handleSpecOrigin(h)
	service.Mux.Handle("GET", "/swagger.yaml", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Spec", "files", "http/goa/swagger/swagger.yaml", "route", "GET /swagger.yaml")
}

// handleSpecOrigin applies the CORS response headers corresponding to the origin.
func handleSpecOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
