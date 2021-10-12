package framework

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMailMessage(t *testing.T) {
	mailMessage := MailMessage{
		From:    "per@softteam.se",
		To:      "per@softteam.se",
		Subject: "test",
		Body:    "test message",
	}

	mailSender := &MailSender{}
	mail := mailSender.createNewMail(mailMessage)

	assert.Equal(t, "From: per@softteam.se\r\nTo: per@softteam.se\r\nSubject: test\r\n\r\ntest message\r\n", mail)
}

func TestGetCredentials(t *testing.T) {
	mailSender := &MailSender{}
	cred, err := mailSender.getCredentials()
	assert.Nil(t, err)
	assert.Equal(t, "akilles.oderland.com", cred.HostName)
	assert.Equal(t, "akilles.oderland.com:465", cred.HostNameAndPort())
	assert.Equal(t, 465, cred.Port)
	assert.Equal(t, "per@softteam.se", cred.FromAddress)
	assert.Equal(t, "per@softteam.se", cred.Name)
}
