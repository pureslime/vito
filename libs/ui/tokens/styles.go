package tokens

import "github.com/charmbracelet/lipgloss"

var (
	// General Styles
	TitleStyle = lipgloss.NewStyle().
			Foreground(AccentColor).
			Bold(true).
			MarginBottom(1)

	PillBadge = lipgloss.NewStyle().
			Background(AccentColor).
			Foreground(TextPrimary).
			Padding(0, 1).
			Bold(true)

	// Error Styles
	ErrorPrefixStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(TextPrimary).
				Background(ErrorBackground).
				Padding(0, 1).
				MarginRight(1)

	ErrorContextStyle = lipgloss.NewStyle().
				Italic(true).
				Foreground(ErrorContextText)

	ErrorBodyStyle = lipgloss.NewStyle().
			Foreground(ErrorBodyText).
			PaddingLeft(4)

	// Version Styles
	VersionNameStyle = lipgloss.NewStyle().
				Background(AccentColor).
				Foreground(Black).
				Padding(0, 1).
				Bold(true)

	VersionNumberStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				BorderForeground(BorderColor).
				PaddingLeft(1).
				Foreground(TextSecondary)

	SloganStyle = lipgloss.NewStyle().
			Foreground(SloganText).
			Italic(true)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(TextPrimary).
			Background(VitoGreen).
			Padding(0, 1).
			Bold(true)

	InfoStyle = lipgloss.NewStyle().
			Foreground(TextPrimary).
			Background(VitoBlue).
			Padding(0, 1).
			Bold(true)
)
