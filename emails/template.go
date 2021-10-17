package emails

import (
	"fmt"
	"os"

	"github.com/matcornic/hermes/v2"
)

type Template struct {
	hermes *hermes.Hermes
}

type RecipientUser struct {
	Fname            string
	Lname            string
	Email            string
	ConfirmationCode string
}

func (template Template) generateSignUpConfirmation(recipientUser *RecipientUser) (string, error) {
	return template.hermes.GeneratePlainText(hermes.Email{
		Body: hermes.Body{
			Name: fmt.Sprintf("%s %s", recipientUser.Fname, recipientUser.Lname),
			Intros: []string{
				"Welcome to Go-Twitter! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Go-Twitter, please click here:",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Confirm your account",
						Link: fmt.Sprintf("%s/%s?token=%s",
							os.Getenv("WEBAPP_HOST"),
							os.Getenv("WEBAPP_EMAIL_CONFIRMATION_PAGE_ROUTE"),
							recipientUser.ConfirmationCode,
						),
					},
				},
			},
		},
	})
}
