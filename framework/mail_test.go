package framework

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMailMessage(t *testing.T) {
	mailMessage := MailMessage{
		From:    "test@test.se",
		To:      "test@test.se",
		Subject: "test",
		Body:    "test message",
	}

	mailSender := &MailSender{}
	mail := mailSender.createNewMail(mailMessage)

	assert.Equal(t, "From: test@test.se\r\nTo: test@test.se\r\nSubject: test\r\n\r\ntest message\r\n", mail)
}

func TestGetCredentials(t *testing.T) {
	mailSender := &MailSender{}
	cred, err := mailSender.getCredentials()
	assert.Nil(t, err)
	crypt := Crypto{}
	hostName, err := crypt.Decrypt(constHostName)
	assert.Nil(t, err)
	email, err := crypt.Decrypt(constEmail)
	assert.Nil(t, err)

	assert.Equal(t, hostName, cred.HostName)
	assert.Equal(t, hostName+":465", cred.HostNameAndPort())
	assert.Equal(t, 465, cred.Port)
	assert.Equal(t, email, cred.FromAddress)
	assert.Equal(t, email, cred.Name)
}
