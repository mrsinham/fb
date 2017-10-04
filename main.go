package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"log"

	"time"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/logrus"
	"github.com/goadesign/goa/middleware"
	goa2 "github.com/mrsinham/fb/http/goa"
	"github.com/mrsinham/fb/http/goa/app"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

func main() {

	//TODO: context, connection handling (keepalive), cache, http testing

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Create service
	service := goa.New("My fizzbuzz api")

	logr := logrus.New()
	service.WithLogger(goalogrus.New(logr))

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "fizzbuzz" controller
	c := goa2.NewFizzController(service)
	app.MountFizzController(service, c)
	// Mount "spec" controller
	c2 := goa2.NewSpecController(service)
	app.MountSpecController(service, c2)

	// TODO: use logger

	s := &http.Server{
		// TODO: conf port
		Addr:           ":8081",
		Handler:        cors.Default().Handler(service.Mux),
		ReadTimeout:    0,
		WriteTimeout:   0,
		TLSConfig:      nil,
		MaxHeaderBytes: 0,
		TLSNextProto:   nil,
		ConnState:      nil,
		ErrorLog:       nil,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("shutting down")
	err := s.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
