package email

import (
	"net/smtp"
)

type SMTPClient interface {
	SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}

type RealSMTPClient struct{}

func (rc *RealSMTPClient) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}

type MockSMTPClient struct {
	CapturedAddr string
	CapturedAuth smtp.Auth
	CapturedFrom string
	CapturedTo   []string
	CapturedMsg  []byte
}

func (mc *MockSMTPClient) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	mc.CapturedAddr = addr
	mc.CapturedAuth = a
	mc.CapturedFrom = from
	mc.CapturedTo = to
	mc.CapturedMsg = msg
	return nil
}
