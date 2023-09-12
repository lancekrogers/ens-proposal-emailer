package email

import (
	"fmt"
	"net/smtp"
)

type EmailSettings struct {
	DestinationAddress string
	Host               string
	Port               string
	SenderUsername     string
	SenderPassword     string
}

type EmailClient struct {
	Settings   EmailSettings
	SMTPClient SMTPClient
}

func NewEmailClient(settings EmailSettings) *EmailClient {
	return &EmailClient{
		Settings:   settings,
		SMTPClient: &RealSMTPClient{},
	}
}

func (ec *EmailClient) SendEmail(proposalID string, proposalTitle string) error {
	to := ec.Settings.DestinationAddress
	// Setup authentication
	auth := smtp.PlainAuth("", ec.Settings.SenderUsername, ec.Settings.SenderPassword, ec.Settings.Host)

	// Create email content
	proposalURL := fmt.Sprintf("https://www.tally.xyz/gov/ens/proposal/%s", proposalID)
	subject := "New ENS proposal"
	body := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\nTitle: %s\r\nLink: %s\r\n", to, subject, proposalTitle, proposalURL)

	// Send email
	return ec.SMTPClient.SendMail(ec.Settings.Host+":"+ec.Settings.Port, auth, to, []string{to}, []byte(body))
}
