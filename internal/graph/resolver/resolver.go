package resolver

import "github.com/kaitoimai01/gqlgen_auth0_sample/internal/user"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Auth0Client user.Auth0ManagementClient
	MailClient  user.MailClient
}
