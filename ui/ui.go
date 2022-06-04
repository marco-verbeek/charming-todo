package ui

import (
	"charming-todo/ui/components/help"
	"charming-todo/ui/components/tabs"
	"charming-todo/utils"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keys utils.KeyMap
	help help.Model
	tabs tabs.Model
}

func New() Model {
	m := Model{
		keys: utils.Keys,
		help: help.NewModel(),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) {
			cmd = tea.Quit
		} else if key.Matches(msg, m.keys.TabNew) {
			tabIndex := m.tabs.AddTab("• Unsaved Tab")
			m.tabs.SetCurrSectionId(tabIndex)
		} else if key.Matches(msg, m.keys.TabClose) {
			m.tabs.CloseCurrentTab()
		} else if key.Matches(msg, m.keys.TabNext) {
			m.tabs.NextTab()
		} else if key.Matches(msg, m.keys.TabPrev) {
			m.tabs.PrevTab()
		} 
	}

	return m, cmd
}

func (m Model) View() string {
	s := strings.Builder{}

	s.WriteString(m.tabs.View())
	s.WriteString("\n")
	s.WriteString(m.help.View())

	return s.String()
}
