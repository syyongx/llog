package handler

import (
	"fmt"
	"github.com/syyongx/llog/types"
	"net/smtp"
	"strings"
)

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
	return mail
}

// Handles a record.
func (m *Mail) Handle(record *types.Record) bool {
	if !m.IsHandling(record) {
		return false
	}
	if m.processors != nil {
		m.ProcessRecord(record)
	}
	err := m.GetFormatter().Format(record)
	if err != nil {
		return false
	}
	m.Write(record)

	return false == m.GetBubble()
}

// Handles a set of records.
func (m *Mail) HandleBatch(records []*types.Record) {
	for _, record := range records {
		m.Handle(record)
	}
}

// Write to network.
func (m *Mail) Write(record *types.Record) {
	message := fmt.Sprintf("To: %v\r\nFrom: %v\r\nSubject: %v\r\nContent-Type: %v; charset=%v\r\n\r\n%v",
		strings.Join(m.To, ";"),
		m.From,
		m.Subject,
		m.GetContentType(),
		m.GetEncoding(),
		record.Formatted.String(),
	)
	err := smtp.SendMail(m.Addr, m.auth, m.From, m.To, []byte(message))
	if err != nil {
		// ...
	}
}

// The content type of the email - Defaults to text/plain. Use text/html for HTML
func (m *Mail) SetContentType(contentType string) {
	m.contentType = contentType
}

// The content type of the email - Defaults to text/plain. Use text/html for HTML
func (m *Mail) GetContentType() string {
	if m.contentType == "" {
		return "text/plain"
	}
	return m.contentType
}

// The encoding for the message - Defaults to UTF-8
func (m *Mail) SetEncoding(encoding string) {
	m.encoding = encoding
}

// The encoding for the message - Defaults to UTF-8
func (m *Mail) GetEncoding() string {
	if m.encoding == "" {
		return "UTF-8"
	}
	return m.encoding
}
