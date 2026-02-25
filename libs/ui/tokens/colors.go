package tokens

import "github.com/charmbracelet/lipgloss"

var (
	// Brand Colors
	VitoBlue      = lipgloss.Color("#00ADD8")
	VitoGreen     = lipgloss.Color("#04B575")
	VitoRed       = lipgloss.Color("#FF4136")
	VitoLightRed  = lipgloss.Color("#FF7670")
	VitoGrey      = lipgloss.Color("#444444")
	VitoLightGrey = lipgloss.Color("#888888")
	VitoDarkGrey  = lipgloss.Color("#333333")
	White         = lipgloss.Color("#FFFFFF")
	Black         = lipgloss.Color("#000000")

	// Semantic Assign (Use tokens)
	AccentColor   = VitoBlue
	SuccessColor  = VitoGreen
	TextPrimary   = White
	TextSecondary = VitoLightGrey

	// Error tokens
	ErrorBackground  = VitoRed
	ErrorContextText = VitoLightGrey
	ErrorBodyText    = VitoLightRed

	// UI Components
	SloganText  = VitoDarkGrey
	BorderColor = VitoDarkGrey
)
