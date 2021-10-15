package emails

import (
	"github.com/matcornic/hermes/v2"
)

type Template struct {
	H *hermes.Hermes
}

type RecipientUser struct {
	fname string
	lname string
	email string
}

// func (template Template) GenerateEmailConfirmationEmail(recipientUser RecipientUser) (string, error) {
// 	return template.H.GenerateHTML(hermes.Email{
// 		Body: hermes.Body{
// 			Name: fmt.Sprintf("%s %s", recipientUser.fname, recipientUser.lname),
// 			Intros: []string {
// 				"Welcome to Go-Twitter! We're very excited to have you on board.",
// 			},
// 			Actions: []hermes.Action{
// 				{
// 					Instructions: "To get started with Go-Twitter, please click here:",
// 					Button: hermes.Button{
// 						Color: "#22BC66",
// 						Text: "Confirm your account",
// 						Link: fmt.Sprintf("%s/confirm-account?token=%s", os.Getenv("WEBAPP_HOST"),)
// 					},
// 				}
// 			}
// 		},
// 	})
// }
