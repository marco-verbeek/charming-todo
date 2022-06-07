package help

import (
	"charming-todo/utils"

	baseHelp "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keys utils.KeyMap
	help baseHelp.Model
}

func NewModel() Model {
	return Model{
		keys: utils.Keys,
		help: baseHelp.New(),
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	}

	return m, nil
}

func (m Model) View() string {
	return m.help.View(m.keys)
}
