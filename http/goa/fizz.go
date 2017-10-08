package goa

import (
	"time"

	"context"

	"strings"

	"github.com/goadesign/goa"
	"github.com/golang/groupcache"
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

	cctx, cancel := context.WithTimeout(ctx.Context, 5*time.Second)
	defer cancel()

	res, err := fizzbuzz.Naive(cctx, ctx.String1, ctx.String2, ctx.Int1, ctx.Int2, ctx.Limit)
	if err != nil {
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	return ctx.OK(res)
}

// Buzz runs the buzz cache action.
func (c *FizzController) BuzzCache(ctx *app.BuzzCacheFizzContext) error {

	cctx, cancel := context.WithTimeout(ctx.Context, 10*time.Second)
	defer cancel()

	var res []byte
	err := Group.Get(cctx, HTTPCacheKey(time.Hour, ctx.String1, ctx.String2, ctx.Int1, ctx.Int2, ctx.Limit), groupcache.AllocatingByteSliceSink(&res))
	if err != nil {
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	return ctx.OK(res)
}

// ExpiresAllKeys expires all the cache for the fizz buzz api
func (c *FizzController) ExpireCache(ctx *app.ExpireCacheFizzContext) error {
	ExpiresAllKeys()
	return ctx.OK(nil)

}

// 300 mb of cache for every nodes
var Group = groupcache.NewGroup("fizzbuzz", 64<<20, groupcache.GetterFunc(
	func(ctx groupcache.Context, key string, dest groupcache.Sink) error {

		s1, s2, i1, i2, l, err := ParamsFromCacheKey(key)
		if err != nil {
			return err
		}

		res, err := fizzbuzz.Naive(ctx.(context.Context), s1, s2, i1, i2, l)
		if err != nil {
			return err
		}

		dest.SetBytes([]byte(strings.Join(res, "|||")))
		return nil
	},
))
