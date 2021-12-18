package user

import (
	"context"
	"fmt"
	"net/mail"
)

// 送信元
var (
	invitaionFrom = &mail.Address{
		Name:    "gqlgen-auth0",
		Address: "noreply@example.com",
	}
	invitaionSubject = "仮登録完了のお知らせ"
)

type RegisterUserInput struct {
	User *User
	Auth0ManagementClient
	MailClient
}

func RegisterUser(ctx context.Context, input *RegisterUserInput) (string, error) {
	var (
		amc = input.Auth0ManagementClient
		mc  = input.MailClient
	)

	auth0UserID, err := amc.CreateUser(ctx, input.User)
	if err != nil {
		return "", fmt.Errorf("failed to create user on auth0: %w", err)
	}

	invitaionUrl, err := amc.CreateInvitationUrl(ctx, auth0UserID)
	if err != nil {
		return "", fmt.Errorf("failed to send invitation: %w", err)
	}

	var (
		body = invitaionUrl
		to   = &mail.Address{
			Name:    input.User.FullName(),
			Address: input.User.Email,
		}
	)
	if err := mc.Send(ctx, invitaionFrom, to, invitaionSubject, body); err != nil {
		return "", fmt.Errorf("failed to send invitation: %w", err)
	}

	return auth0UserID, nil
}
