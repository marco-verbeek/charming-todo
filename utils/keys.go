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
	TabSave  key.Binding

	ItemToggle    key.Binding
	ItemNext      key.Binding
	ItemPrev      key.Binding
	ItemAddIndent key.Binding
	ItemRemIndent key.Binding
	ItemNew       key.Binding
	ItemDelete    key.Binding

	ItemTop    key.Binding
	ItemBottom key.Binding

	EditModeStart  key.Binding
	EditModeCancel key.Binding
	EditModeSave   key.Binding
	EditModeLeft   key.Binding
	EditModeRight  key.Binding
}

var Keys = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q", "quit"),
	),

	TabNew: key.NewBinding(
		key.WithKeys("ctrl+n"),
		key.WithHelp("ctrl+n", "new tab"),
	),
	TabClose: key.NewBinding(
		key.WithKeys("ctrl+w"),
		key.WithHelp("ctrl+w", "close tab"),
	),
	TabNext: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("/l", "next tab"),
	),
	TabPrev: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("/h", "previous tab"),
	),
	TabSave: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "save tab"),
	),

	ItemToggle: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "check item"),
	),
	ItemNext: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "next item"),
	),
	ItemPrev: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "prev item"),
	),
	ItemAddIndent: key.NewBinding(
		key.WithKeys("tab"),
	),
	ItemRemIndent: key.NewBinding(
		key.WithKeys("shift+tab"),
	),
	ItemNew: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new item"),
	),
	ItemDelete: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "delete item"),
	),

	ItemTop: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "navigate to top item"),
	),
	ItemBottom: key.NewBinding(
		key.WithKeys("b"),
		key.WithHelp("b", "navigate to bottom item"),
	),

	EditModeStart: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit todo item"),
	),
	EditModeCancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "cancel edit mode"),
	),
	EditModeSave: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "save current edits"),
	),
	EditModeRight: key.NewBinding(
		key.WithKeys("right"),
	),
	EditModeLeft: key.NewBinding(
		key.WithKeys("left"),
	),
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ItemPrev, k.ItemNext, k.TabNext, k.TabPrev},
		{k.TabNew, k.TabClose, k.Quit},
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.ItemPrev,
		k.ItemNext,
		k.TabNext,
		k.TabPrev,
		k.TabNew,
		k.TabClose,
		k.Quit,
	}
}
