package help

import (
	"charming-todo/utils"

	baseHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keys utils.KeyMap
	help baseHelp.Model
}

func NewModel() Model {
	return Model{
		keys: utils.Keys,
		help: baseHelp.NewModel(),
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
	return m.help.ShortHelpView(m.collectHelpBindings())
}

// returns the keybindings to display.
func (m Model) collectHelpBindings() []key.Binding {
	k := m.keys
	bindings := []key.Binding{}

	bindings = append(bindings, k.Quit)

	bindings = append(bindings, k.TabNew)
	bindings = append(bindings, k.TabClose)

	bindings = append(bindings, k.TabNext)
	bindings = append(bindings, k.TabPrev)

	return bindings
}
