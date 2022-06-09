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

	TodoItemToggle    key.Binding
	TodoItemNext      key.Binding
	TodoItemPrev      key.Binding
	TodoItemAddIndent key.Binding
	TodoItemRemIndent key.Binding
	TodoItemNew       key.Binding
	TodoItemDelete    key.Binding
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

	TodoItemToggle: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "check item"),
	),
	TodoItemNext: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "next item"),
	),
	TodoItemPrev: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "prev item"),
	),
	TodoItemAddIndent: key.NewBinding(
		key.WithKeys("tab"),
	),
	TodoItemRemIndent: key.NewBinding(
		key.WithKeys("shift+tab"),
	),
	TodoItemNew: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new item"),
	),
	TodoItemDelete: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "delete item"),
	),
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.TodoItemPrev, k.TodoItemNext, k.TabNext, k.TabPrev},
		{k.TabNew, k.TabClose, k.Quit},
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.TodoItemPrev,
		k.TodoItemNext,
		k.TabNext,
		k.TabPrev,
		k.TabNew,
		k.TabClose,
		k.Quit,
	}
}
