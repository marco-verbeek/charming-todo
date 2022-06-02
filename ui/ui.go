package ui

import (
	"charming-todo/ui/components/help"
	"charming-todo/utils"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	keys utils.KeyMap
	help help.Model
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
			fmt.Println("TabNew")
		} else if key.Matches(msg, m.keys.TabClose) {
			fmt.Println("TabClose")
		}
	}

	return m, cmd
}

func (m Model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		m.help.View(),
	)
}
