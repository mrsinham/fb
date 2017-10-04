package goa

import (
	"time"

	"context"

	"github.com/goadesign/goa"
	"github.com/mrsinham/fb/fizzbuzz"
	"github.com/mrsinham/fb/http/goa/app"
)

// FizzController implements the fizz resource.
type FizzController struct {
	*goa.Controller
}

// NewFizzController creates a fizz controller.
func NewFizzController(service *goa.Service) *FizzController {
	return &FizzController{Controller: service.NewController("FizzController")}
}

// Buzz runs the buzz action.
func (c *FizzController) Buzz(ctx *app.BuzzFizzContext) error {

	cctx, cancel := context.WithTimeout(ctx.Context, 10*time.Second)
	defer cancel()

	res, err := fizzbuzz.Naive(cctx, ctx.String1, ctx.String2, ctx.Int1, ctx.Int2, ctx.Limit)
	if err != nil {
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	return ctx.OK(res)
	//return ctx.OK(&app.FizzBuzz{
	//	List: res,
	//})
}
