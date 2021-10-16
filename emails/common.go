package emails

import (
	"os"

	"github.com/mailgun/mailgun-go/v4"
)

type Emails struct {
	// Mailgun client
	mg *mailgun.MailgunImpl
	// Gathers all mail template generators
	template *Template
}

func InitEmails() *Emails {
	hermes := initHermes()
	return &Emails{
		mg: initMailgunClient(),
		template: &Template{
			hermes: hermes,
		},
	}
}

func getRecipientEmail(recipientUser *RecipientUser) string {
	if os.Getenv("MAILGUN_RECIPIENT_EMAIL") != "" {
		return os.Getenv("MAILGUN_RECIPIENT_EMAIL")
	}
	return recipientUser.Email
}
