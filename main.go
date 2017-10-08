package main

//go:generate goagen bootstrap -d github.com/mrsinham/fb/http/design -o http/goa

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/logrus"
	"github.com/goadesign/goa/middleware"
	"github.com/golang/groupcache"
	fbgoa "github.com/mrsinham/fb/http/goa"
	"github.com/mrsinham/fb/http/goa/app"
	"github.com/rs/cors"
)

func main() {

	//TODO: cache

	peers := flag.String("pool", "http://localhost:8080", "server pool list")
	port := flag.String("port", ":8081", "port to start the api")

	flag.Parse()

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

	// random are less deterministic
	rand.Seed(time.Now().UnixNano())

	// cache
	listOfPeers := strings.Split(*peers, ",")
	hp := groupcache.NewHTTPPool(listOfPeers[0])
	hp.Set(listOfPeers...)

	// Mount "fizzbuzz" controller
	c := fbgoa.NewFizzController(service)
	app.MountFizzController(service, c)
	// Mount "spec" controller
	c2 := fbgoa.NewSpecController(service)
	app.MountSpecController(service, c2)

	s := &http.Server{
		Addr:         *port,
		Handler:      cors.Default().Handler(service.Mux),
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	s.SetKeepAlivesEnabled(true)

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
