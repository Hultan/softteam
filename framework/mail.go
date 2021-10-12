package framework

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strconv"
)

type MailSender struct {
}

type MailMessage struct {
	From    string
	To      string
	Subject string
	Body    string
}

type credentials struct {
	HostName    string
	Port        int
	FromAddress string
	Name        string
	Password    string
}

const constPassword = "600189099c094d21a95c0ed6f25b70a8e5fbe25128d9e0af44c54086124849532e481633930f88148e4d7d32"

func (m *MailSender) getCredentials() (*credentials, error) {
	crypto := new(Crypto)
	password, err := crypto.Decrypt(constPassword)
	if err != nil {
		return nil, err
	}
	return &credentials{
		HostName:    "akilles.oderland.com",
		Port:        465,
		FromAddress: "per@softteam.se",
		Name:        "per@softteam.se",
		Password:    password,
	}, nil
}

func (m *MailSender) SendMail(mailMessage MailMessage) error {
	cred, err := m.getCredentials()
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", cred.Name, cred.Password, cred.HostName)

	// TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         cred.HostName,
	}

	mailConn, err := tls.Dial("tcp", cred.HostNameAndPort(), tlsConfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(mailConn, cred.HostName)
	if err != nil {
		return err
	}

	if err = client.Auth(auth); err != nil {
		return err
	}
	if err = client.Mail(mailMessage.From); err != nil {
		return err
	}
	if err = client.Rcpt(mailMessage.To); err != nil {
		return err
	}
	writeCloser, err := client.Data()
	if err != nil {
		return err
	}

	mail := m.createNewMail(mailMessage)

	_, err = writeCloser.Write([]byte(mail))
	if err != nil {
		return err
	}
	err = writeCloser.Close()
	if err != nil {
		return err
	}

	client.Quit()

	return nil
}

func (m *MailSender) createNewMail(mailMessage MailMessage) string {
	headers := make(map[string]string)
	headers["From"] = mailMessage.From
	headers["To"] = mailMessage.To
	headers["Subject"] = mailMessage.Subject

	message := ""
	for headerName, headerValue := range headers {
		message += fmt.Sprintf("%s: %s\r\n", headerName, headerValue)
	}
	message += "\r\n" + mailMessage.Body + "\r\n"

	return message
}

func (c *credentials) HostNameAndPort() string {
	return fmt.Sprintf("%v:%v", c.HostName, strconv.Itoa(c.Port))
}
