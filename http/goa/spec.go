package goa

import (
	"github.com/goadesign/goa"
)

// SpecController implements the spec resource.
type SpecController struct {
	*goa.Controller
}

// NewSpecController creates a spec controller.
func NewSpecController(service *goa.Service) *SpecController {
	return &SpecController{Controller: service.NewController("SpecController")}
}
