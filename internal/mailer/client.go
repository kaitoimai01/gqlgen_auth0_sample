package mailer

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"
	"os"

	"github.com/sendgrid/sendgrid-go"
	sgmail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Client struct {
	sg *sendgrid.Client
}

func NewClient(c *sendgrid.Client) *Client {
	return &Client{sg: c}
}

func (c *Client) Send(ctx context.Context, from *mail.Address, to *mail.Address, subject string, message string) error {
	msg := sgmail.NewSingleEmail(
		sgmail.NewEmail(from.Name, from.Address),
		subject,
		sgmail.NewEmail(to.Name, to.Address),
		subject,
		message)

	res, err := c.sg.SendWithContext(ctx, msg)
	if err != nil {
		return err
	}

	// OK(200) ではなく、Accepted(201) が返却されます。
	if res.StatusCode != http.StatusAccepted {
		fmt.Fprintf(os.Stderr, res.Body)
		return fmt.Errorf("unexpected status from sendgrid api: %d", res.StatusCode)
	}

	return nil
}
