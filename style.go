package main

import "github.com/charmbracelet/lipgloss"

var (
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))

	normalStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder())

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)
