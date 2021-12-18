//go:generate mockgen -source=./interfaces.go -destination=./mock_member/mock_interfaces.go
package user

import (
	"context"
	"net/mail"
)

type Auth0ManagementClient interface {
	CreateUser(ctx context.Context, u *User) (string, error)
	CreateInvitationUrl(ctx context.Context, auth0UserID string) (string, error)
}

type MailClient interface {
	Send(ctx context.Context, from *mail.Address, to *mail.Address, subject string, message string) error
}
