package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	tabs          []string
	CurrSectionId int
}

func NewModel() Model {
	return Model{
		tabs:          make([]string, 0),
		CurrSectionId: 0,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var tabs []string

	for i, tabsTitle := range m.tabs {
		if m.CurrSectionId == i {
			tabs = append(tabs, activeTab.Render(tabsTitle))
		} else {
			tabs = append(tabs, tab.Render(tabsTitle))
		}
	}

	width := 100
	renderedTabs := lipgloss.NewStyle().Width(width).MaxWidth(width).Render(lipgloss.JoinHorizontal(lipgloss.Top, tabs...))

	return tabsRow.Copy().Width(width).MaxWidth(width).Render(lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs))
}

func (m *Model) SetCurrSectionId(id int) {
	m.CurrSectionId = id
}

func (m *Model) AddTab(name string) int {
	if len(m.tabs) == 0 {
		m.tabs = append(m.tabs, name)
		return 0
	}

	m.tabs = append(m.tabs[:m.CurrSectionId], append([]string{name}, m.tabs[m.CurrSectionId:]...)...)
	m.CurrSectionId = m.CurrSectionId + 1

	return m.CurrSectionId
}

func (m *Model) CloseCurrentTab() {
	if len(m.tabs) > 0 {
		m.tabs = append(m.tabs[:m.CurrSectionId], m.tabs[m.CurrSectionId+1:]...)
		m.PrevTab()
	}
}

func (m *Model) NextTab() {
	if m.CurrSectionId+1 == len(m.tabs) {
		m.CurrSectionId = 0
	} else {
		m.CurrSectionId++
	}
}

func (m *Model) PrevTab() {
	if m.CurrSectionId == 0 {
		m.CurrSectionId = len(m.tabs) - 1
	} else {
		m.CurrSectionId--
	}
}
