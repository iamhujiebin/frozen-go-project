package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("en.toml")
	bundle.MustParseMessageFileBytes([]byte(`
HelloWorld = "Hola Mundo!"
`), "es.toml")

	{
		localizer := i18n.NewLocalizer(bundle, "en-US")
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "HelloWorld2"}))
	}
	{
		localizer := i18n.NewLocalizer(bundle, "es-ES")
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "HelloWorld"}))
	}
}
