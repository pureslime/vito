package ui

import (
	"fmt"
	"libs/ui/tokens"
	"os"
)

func PrintErr(err error, context string) {
	header := fmt.Sprintf("%s Error in %s", tokens.ErrorPrefixStyle.Render("VITO"), tokens.ErrorContextStyle.Render(context))
	body := tokens.ErrorBodyStyle.Render(fmt.Sprintf("↳ %s", err.Error()))

	fmt.Fprintf(os.Stderr, "\n%s\n%s\n\n", header, body)
}

func PrintSuccess(msg string) {
	prefix := tokens.SuccessStyle.Render("SUCCESS👌")

	fmt.Printf("\n%s %s", prefix, msg)
}

func PrintInfo(msg string) {
	prefix := tokens.InfoStyle.Render("INFO👉")

	fmt.Printf("\n%s %s", prefix, msg)
}

func PrintVersion(appName string, version string) {
	fmt.Printf("%s version %s\n", tokens.VersionNameStyle.Render(appName), tokens.VersionNumberStyle.Render(version))
	fmt.Printf("%s\n\n", tokens.SloganStyle.Render("The extensible pill-based CLI engine."))
}

func MakeTitle(t string) string {
	return tokens.TitleStyle.Render(t)
}
func BodyText(t string) string {
	style := tokens.ErrorContextStyle
	return style.Italic(false).Render(t)
}

func Badge(t string) string {
	return tokens.PillBadge.Render(t)
}
