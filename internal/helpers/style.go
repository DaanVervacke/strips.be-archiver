package helpers

import (
	"github.com/charmbracelet/lipgloss"
)

var TitleStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("253")).
	Foreground(lipgloss.AdaptiveColor{Light: "#251a18", Dark: "#251a18"}).
	Bold(true)

var SuccessStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#00FF00", Dark: "#00FF00"})

var ErrorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#FF0000", Dark: "#FF0000"})

var WarningStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#FFA500", Dark: "#FFA500"})
