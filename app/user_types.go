// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "jwt-signin": Application User Types
//
// Command:
// $ goagen
// --design=github.com/Microkubes/jwt-issuer/design
// --out=$(GOPATH)/src/github.com/Microkubes/jwt-issuer
// --version=v1.2.0-dirty

package app

// credentials user type.
type credentials struct {
	// Credentials: email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Credentials: password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// Access scope (api:read, api:write)
	Scope *string `form:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
}

// Publicize creates Credentials from credentials
func (ut *credentials) Publicize() *Credentials {
	var pub Credentials
	if ut.Email != nil {
		pub.Email = ut.Email
	}
	if ut.Password != nil {
		pub.Password = ut.Password
	}
	if ut.Scope != nil {
		pub.Scope = ut.Scope
	}
	return &pub
}

// Credentials user type.
type Credentials struct {
	// Credentials: email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Credentials: password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// Access scope (api:read, api:write)
	Scope *string `form:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
}
