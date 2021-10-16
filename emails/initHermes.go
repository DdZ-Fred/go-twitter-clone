package emails

import (
	"os"

	"github.com/matcornic/hermes/v2"
)

func initHermes() *hermes.Hermes {
	return &hermes.Hermes{
		Product: hermes.Product{
			Name: "Go Twitter Clone",
			Link: os.Getenv("WEBAPP_HOST"),
		},
	}
}
