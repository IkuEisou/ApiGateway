package design

import (
	// . "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Authentication API", func() {
	Title("The Authentication API ")
	Description("An Authentication API")
	Scheme("http")
	Host("localhost:8080")
	BasePath("/api/")
	Origin("*", func() {
		Headers("Content-Type")
		Methods("GET", "POST", "PATCH", "DELETE", "PUT", "OPTION")
	})
	Consumes("application/json")
	Produces("application/json")
})
