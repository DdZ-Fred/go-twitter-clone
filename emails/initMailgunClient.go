package emails

import (
	"os"

	"github.com/mailgun/mailgun-go/v4"
)

func initMailgunClient() *mailgun.MailgunImpl {
	return mailgun.NewMailgun(
		os.Getenv("MAILGUN_DOMAIN"),
		os.Getenv("MAILGUN_PRIVATE_API_KEY"),
	)
}
