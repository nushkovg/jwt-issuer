package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("jwt-signin", func() {
	Title("JWT Sign in")
	Description("Sign in and generate JWT token with claims")
	Version("1.0")
	Scheme("http")
	Host("localhost:8080")
	Consumes("application/x-www-form-urlencoded", func() {
		Package("github.com/goadesign/goa/encoding/form")
	})
})

var _ = Resource("jwt", func() {
	BasePath("jwt")
	Description("Sign in")

	Action("signin", func() {
		Description("Signs in the user and generates JWT token")
		Payload(CredentialsPayload)
		Routing(POST("/signin"))
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
		Response(Created)
	})

})

var CredentialsPayload = Type("Credentials", func() {
	Attribute("username", String, "Credentials: username")
	Attribute("password", String, "Credentials: password")
	Attribute("scope", String, "Access scope (api:read, api:write)")
})