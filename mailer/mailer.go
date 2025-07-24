package mailer

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	FromName     string
	FromEmail    string
	ToName       string
	ToEmail      string
	Subject      string
	PlainContent string
	HTMLContent  string
}

// SendEmail sends an email using SendGrid.
func SendEmail(e Email) error {
	from := mail.NewEmail(e.FromName, e.FromEmail)
	to := mail.NewEmail(e.ToName, e.ToEmail)

	message := mail.NewSingleEmail(from, e.Subject, to, e.PlainContent, e.HTMLContent)

	apiKey := "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"
	if apiKey == "" {
		return fmt.Errorf("SENDGRID_API_KEY not set in environment")
	}

	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("failed to send email: status=%d, body=%s", response.StatusCode, response.Body)
	}

	return nil
}