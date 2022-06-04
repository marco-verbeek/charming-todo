package utils

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Quit key.Binding

	TabNew   key.Binding
	TabClose key.Binding
	TabNext  key.Binding
	TabPrev  key.Binding
}

var Keys = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc", "q"),
		key.WithHelp("q", "quit"),
	),
	TabNew: key.NewBinding(
		key.WithKeys("ctrl+n"),
		key.WithHelp("ctrl+n", "new tab"),
	),
	TabClose: key.NewBinding(
		key.WithKeys("ctrl+w", "x"),
		key.WithHelp("ctrl+w", "close tab"),
	),
	TabNext: key.NewBinding(
		key.WithKeys("ctrl+tab"),
		key.WithHelp("ctrl+tab", "next tab"),
	),
	TabPrev: key.NewBinding(
		key.WithKeys("ctrl+shift+tab"),
		key.WithHelp("ctrl+shift+tab", "previous tab"),
	),
}
