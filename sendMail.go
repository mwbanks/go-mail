package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

// StartTLS Email Example

var auth smtp.Auth
var server string

// Init sets the credentials and server for the package
func Init(username, password, servername string) {
	// Accepts strings of the form "smtp.server.com:587"
	server = servername
	host, _, _ := net.SplitHostPort(server)
	auth = smtp.PlainAuth("", username, password, host)
}

// SendMail sends the message to the specified user from the specified other user.
func SendMail(to, from mail.Address, subject, body string) {
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Mime-Version"] = "1.0;"
	headers["Content-Type"] = "text/html; charset=\"ISO-8859-1\";"
	headers["Content-Transfer-Encoding"] = "7bit;"

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	host, _, _ := net.SplitHostPort(server)

	// TLS config
	tlsconfig := &tls.Config{
		ServerName: host,
	}

	c, err := smtp.Dial(server)
	if err != nil {
		log.Panic(err)
	}

	c.StartTLS(tlsconfig)

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

}
