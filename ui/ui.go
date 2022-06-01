package ui

import (
	"charming-todo/utils"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	keys utils.KeyMap
}

func New() Model {
	m := Model{
		keys: utils.Keys,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil;
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd 

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) { 
			cmd = tea.Quit
		}
	}

	return m, cmd
}

func (m Model) View() string {
	return "ui.View()"
}