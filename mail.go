package gomailer

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"net/smtp"
	"reflect"
)

// Mail ...
type Mail struct {
	Auth      auth
	Remote    string
	Sender    string
	Recipient string
	Content   string
	Headers   headers
	Error     error
	config    *tls.Config
}

type auth struct {
	Identity string
	Username string
	Password string
	Host     string
}

type headers struct {
	MIMEVersion string `gomailer:"MIME-Version"`
	ContentType string `gomailer:"Content-Type"`
	Charset     string `gomailer:"Charset"`
	Subject     string `gomailer:"Subject"`
	Other       string
}

// New ...
func New() *Mail {
	mail := new(Mail)
	mail.Headers = headers{
		Charset:     "\"UTF-8\"",
		MIMEVersion: "1.0",
		ContentType: "text/html",
	}
	return mail
}

// Send ...
func (m *Mail) Send() error {
	// Connect to remote SMTP
	c, err := smtp.Dial(m.Remote)
	if err != nil {
		return err
	}
	defer c.Close()

	if m.Auth.Username != "" && m.Auth.Password != "" {
		if err := c.Auth(smtp.PlainAuth(m.Auth.Identity, m.Auth.Username, m.Auth.Password, m.Auth.Host)); err != nil {
			return err
		}
	}

	// Define sender and recipient
	if err := c.Mail(m.Sender); err != nil {
		return err
	}
	if err := c.Rcpt(m.Recipient); err != nil {
		return err
	}

	// Create a Writer
	wc, err := c.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	// Create a Buffer
	buffer := bytes.NewBufferString(m.getHeaders())
	buffer.WriteString(m.Content)

	// Buffer to Writer
	if _, err = buffer.WriteTo(wc); err != nil {
		return err
	}

	return nil
}

func (m *Mail) getHeaders() string {
	var headers string

	value := reflect.ValueOf(m.Headers)
	hType := reflect.TypeOf(m.Headers)

	for i := 0; i < value.NumField(); i++ {
		content := value.Field(i).Interface().(string)

		if content != "" {
			f, _ := hType.FieldByName(value.Type().Field(i).Name)
			name, _ := f.Tag.Lookup("gomailer")
			headers = headers + name + ": " + content + "\r\n"
		}
	}
	return headers
}

// ParseTemplate ...
func (m *Mail) ParseTemplate(file string, data interface{}) error {
	t, err := template.ParseFiles(file)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	m.Content = buf.String()
	return nil
}
