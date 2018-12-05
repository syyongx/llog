package handler

import (
	"fmt"
	"github.com/syyongx/llog/types"
	"net/smtp"
	"strings"
)

// Mail handler struct definition
type Mail struct {
	Processing

	Addr     string   // SMTP server address
	Username string   // SMTP server login username
	Password string   // SMTP server login password
	Subject  string   // The subject of the email
	From     string   // The sender of the mail
	To       []string // The email addresses to which the message will be sent

	contentType string // The Content-type for the message
	encoding    string // The encoding for the message
	auth        smtp.Auth
}

// NewMail new mail handler
func NewMail(address, username, password, from, subject string, to []string, level int, bubble bool) *Mail {
	auth := smtp.PlainAuth(
		"",
		username,
		password,
		strings.Split(address, ":")[0],
	)
	mail := &Mail{
		Addr:     address,
		Username: username,
		Password: password,
		Subject:  subject,
		From:     from,
		To:       to,
		auth:     auth,
	}
	mail.SetLevel(level)
	mail.SetBubble(bubble)
	mail.Writer = mail.Write
	return mail
}

// Write to network.
func (m *Mail) Write(record *types.Record) {
	message := fmt.Sprintf("To: %v\r\nFrom: %v\r\nSubject: %v\r\nContent-Type: %v; charset=%v\r\n\r\n%v",
		strings.Join(m.To, ";"),
		m.From,
		m.Subject,
		m.ContentType(),
		m.Encoding(),
		record.Formatted.String(),
	)
	err := smtp.SendMail(m.Addr, m.auth, m.From, m.To, []byte(message))
	if err != nil {
		// ...
	}
}

// SetContentType Set the content type of the email - Defaults to text/plain. Use text/html for HTML
func (m *Mail) SetContentType(contentType string) {
	m.contentType = contentType
}

// ContentType Get the content type of the email - Defaults to text/plain. Use text/html for HTML
func (m *Mail) ContentType() string {
	if m.contentType == "" {
		return "text/plain"
	}
	return m.contentType
}

// SetEncoding Set the encoding for the message - Defaults to UTF-8
func (m *Mail) SetEncoding(encoding string) {
	m.encoding = encoding
}

// Encoding Get the encoding for the message - Defaults to UTF-8
func (m *Mail) Encoding() string {
	if m.encoding == "" {
		return "UTF-8"
	}
	return m.encoding
}
