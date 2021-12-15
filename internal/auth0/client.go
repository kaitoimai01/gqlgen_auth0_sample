package auth0

import (
	"context"
	"fmt"

	"github.com/kaitoimai01/gqlgen_auth0_sample/internal/user"
	"github.com/sethvargo/go-password/password"
	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

const (
	// NOTE: connection 定数は Auth0 のダッシュボードより、
	//       Authentication > Database > Database Connections にて確認できます。
	connection       = "Username-Password-Authentication"
	invitationTTLSec = 86400
)

type Client struct {
	management *management.Management
}

func NewClient(mgmt *management.Management) *Client {
	return &Client{management: mgmt}
}

func (c *Client) CreateMember(ctx context.Context, u *user.User) (string, error) {
	// 招待メールからパスワードを設定するまでの間の仮パスワードを生成します。
	pass, err := generateRandomPassword()
	if err != nil {
		return "", fmt.Errorf("could not generate password: %w", err)
	}

	// NOTE: Auth0 が用意しているデータベースにリクエストできたユーザの情報を登録します。
	//       この段階ではまだユーザのメールアドレスの認証とパスワードの設定が未完了のため、
	//       EmailVerified と VerifyEmail のフィールドはともに false にする必要があります。
	managementUser := &management.User{
		Connection:    auth0.String(connection),
		Email:         auth0.String(u.Email),
		Name:          auth0.String(u.FullName()),
		Password:      auth0.String(pass),
		EmailVerified: auth0.Bool(false),
		VerifyEmail:   auth0.Bool(false),
	}
	if err := c.management.User.Create(managementUser, management.Context(ctx)); err != nil {
		return "", fmt.Errorf("could not create a user on auth0: %w", err)
	}

	// パスワード再設定 URL を作成する際に生成されたユーザ ID が必要になります。
	return auth0.StringValue(managementUser.ID), nil
}

func generateRandomPassword() (string, error) {
	const (
		length      = 32
		numDigits   = 10
		numSymbols  = 10
		noUpper     = false
		allowRepeat = false
	)

	return password.Generate(length, numDigits, numSymbols, noUpper, allowRepeat)
}

func (c *Client) CreateInvitationUrl(ctx context.Context, auth0UserID string) (string, error) {
	// NOTE: パスワード再設定用の URL を作成します。
	//       MarkEmailAsVerified フィールドを true にし、保存されることでメールアドレスの認証が完了となります。
	//       また management.Ticket 構造体には ResultURL フィールドがあり、そこに設定した URL にリダイレクトさせることも可能です。
	t := &management.Ticket{
		UserID:              auth0.String(auth0UserID),
		TTLSec:              auth0.Int(invitationTTLSec),
		MarkEmailAsVerified: auth0.Bool(true),
	}
	if err := c.management.Ticket.ChangePassword(t, management.Context(ctx)); err != nil {
		return "", fmt.Errorf("could not create ticket for change password: %w", err)
	}

	return auth0.StringValue(t.Ticket) + "invite", nil
}
