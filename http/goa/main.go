//go:generate goagen bootstrap -d github.com/mrsinham/fb/http/design

package goa

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/mrsinham/fb/http/goa/app"
)

func main() {
	// Create service
	service := goa.New("My fizzbuzz api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "fizz" controller
	c := NewFizzController(service)
	app.MountFizzController(service, c)
	// Mount "spec" controller
	c2 := NewSpecController(service)
	app.MountSpecController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
