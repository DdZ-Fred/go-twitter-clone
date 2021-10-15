package i18n

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func InitLocalizer() *i18n.Localizer {
	bundle := i18n.NewBundle(language.BritishEnglish)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("i18n/locales/en-GB.json")
	bundle.LoadMessageFile("i18n/locales/fr-FR.json")

	localizer := i18n.NewLocalizer(
		bundle,
		language.BritishEnglish.String(),
		language.French.String(),
	)
	return localizer
}
