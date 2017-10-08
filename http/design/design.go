package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("My fizzbuzz api", func() { // "My API" is the name of the API used in docs
	Title("Fizzbuzz")                          // Documentation title
	Description("The next step into fizzbuzz") // Longer documentation description
	Scheme("http")                             // HTTP scheme used by Swagger and clients
	Consumes("application/json")               // Media types supported by the API
	Produces("application/json")
})

var _ = Resource("spec", func() {

	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.yaml", "http/goa/swagger/swagger.yaml")
	Files("/swagger.json", "http/goa/swagger/swagger.json")
})

var _ = Resource("fizz", func() {
	BasePath("fizz")
	DefaultMedia("application/json")
	Action("buzz", func() {
		Routing(GET("buzz"))
		Description("given the two string and the fizzbuzz param, " +
			"you will have the fizzbuzz algorithm applied : a json string array from 1 to limit where the " +
			"offset is replaced with string1 when " +
			"the offset is a multiplier of int1 and replaced with string2 if the offset is a multiplier of int2." +
			" If the offset is a multiplier of both int1 and int2, the offset will be replaced by string1+string2")
		Params(func() {
			Param("string1", String, "the first string, this is the fizz string", func() {
				Example("fizz")
				MinLength(2)
			})
			Param("string2", String, "the second string, this is the buzz string", func() {
				Example("buzz")
				MinLength(2)
			})
			Param("int1", Integer, "if the offset of the response array is a multiplier of int1, replace with string1", func() {
				Example(3)
				Minimum(1)
			})
			Param("int2", Integer, "if the offset of the response array is a multiplier of int2, replace with string2", func() {
				Example(5)
				Minimum(1)
			})
			Param("limit", Integer, "the fizzbuzz algorithm will produce an array up to the given limit", func() {
				Example(15)
				Minimum(1)
			})

			Required("string1", "string2", "int1", "int2", "limit")
		})

		Response(OK, ArrayOf(String), func() {})
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)

	})

	Action("buzz_cache", func() {
		Routing(GET("buzz_cache"))
		Description("same as buzz, but with cache")
		Params(func() {
			Param("string1", String, "the first string, this is the fizz string", func() {
				Example("fizz")
				MinLength(2)
			})
			Param("string2", String, "the second string, this is the buzz string", func() {
				Example("buzz")
				MinLength(2)
			})
			Param("int1", Integer, "if the offset of the response array is a multiplier of int1, replace with string1", func() {
				Example(3)
				Minimum(1)
			})
			Param("int2", Integer, "if the offset of the response array is a multiplier of int2, replace with string2", func() {
				Example(5)
				Minimum(1)
			})
			Param("limit", Integer, "the fizzbuzz algorithm will produce an array up to the given limit", func() {
				Example(15)
				Minimum(1)
			})

			Required("string1", "string2", "int1", "int2", "limit")
		})

		Response(OK)
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest)

	})

	Action("expire_cache", func() {
		Routing(GET("expire_cache"))
		Description("expire all cached entries")
		Response(OK)
	})

})
