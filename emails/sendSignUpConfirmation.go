package emails

import (
	"context"
	"fmt"
	"os"
	"time"
)

func (emails Emails) SendSignUpConfirmation(recipientUser *RecipientUser) (string, error) {
	body, _ := emails.template.generateSignUpConfirmation(recipientUser)

	message := emails.mg.NewMessage(
		os.Getenv("MAILGUN_SENDER_EMAIL"),
		fmt.Sprintf("Welcome to Go-Twitter, %s %s! Please confirm your email to continue", recipientUser.Fname, recipientUser.Lname),
		body,
		getRecipientEmail(recipientUser),
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, id, err := emails.mg.Send(ctx, message)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return id, nil
}
