package email

import (
	"fmt"
	"testing"
)

func TestSendEmail(t *testing.T) {
	mockSMTP := &MockSMTPClient{}
	settings := EmailSettings{
		DestinationAddress: "recipient@example.com",
		Host:               "smtp.example.com",
		Port:               "587",
		SenderUsername:     "sender@example.com",
		SenderPassword:     "password",
	}

	client := &EmailClient{
		Settings:   settings,
		SMTPClient: mockSMTP,
	}

	proposalID := "12345"
	proposalTitle := "Test Proposal"
	expectedBody := fmt.Sprintf("To: %s\r\nSubject: New ENS proposal\r\n\r\nTitle: %s\r\nLink: https://www.tally.xyz/gov/ens/proposal/%s\r\n", settings.DestinationAddress, proposalTitle, proposalID)

	err := client.SendEmail(proposalID, proposalTitle)
	if err != nil {
		t.Errorf("SendEmail failed: %s", err)
	}

	// Assert that SendMail was called with the expected parameters
	if string(mockSMTP.CapturedMsg) != expectedBody {
		t.Errorf("Expected email body to be '%s' but got '%s'", expectedBody, string(mockSMTP.CapturedMsg))
	}
}
