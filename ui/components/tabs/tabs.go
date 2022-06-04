package tabs

import (
	"charming-todo/data"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	lists     *[]data.TodoList
	CurrTabId int
}

func NewModel(lists *[]data.TodoList) Model {
	return Model{
		lists:     lists,
		CurrTabId: 0,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var tabs []string

	for i, list := range *m.lists {
		if !list.Displayed {
			continue
		}

		if m.CurrTabId == i {
			tabs = append(tabs, activeTab.Render(list.Title))
		} else {
			tabs = append(tabs, tab.Render(list.Title))
		}
	}

	width := 100
	renderedTabs := lipgloss.NewStyle().Width(width).MaxWidth(width).Render(lipgloss.JoinHorizontal(lipgloss.Top, tabs...))

	return tabsRow.Copy().Width(width).MaxWidth(width).Render(lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs))
}

func (m *Model) SetCurrSectionId(id int) {
	m.CurrTabId = id
}

func (m *Model) AddTab() int {
	tab := data.TemplateTodoList()

	if len(*m.lists) == 0 {
		*m.lists = append(*m.lists, tab)
		return 0
	}

	m.CurrTabId++
	*m.lists = append((*m.lists)[:m.CurrTabId], append([]data.TodoList{tab}, (*m.lists)[m.CurrTabId:]...)...)

	return m.CurrTabId
}

func (m *Model) CloseCurrentTab() int {
	if len(*m.lists) > 0 {
		(*m.lists)[m.CurrTabId].Displayed = false
		return m.PrevTab()
	}

	return 0
}

func (m *Model) NextTab() int {
	// start at next element in array
	loopIdx := m.CurrTabId + 1
	totalLoops := 0

	// While we haven't looped over all elements
	for totalLoops < len(*m.lists) {
		// Out of bounds - start from first elem
		if loopIdx >= len(*m.lists) {
			loopIdx = 0
		}

		// Found an elem that can be displayed
		if (*m.lists)[loopIdx].Displayed {
			break
		}

		loopIdx++
		totalLoops++
	}

	m.CurrTabId = loopIdx
	return m.CurrTabId
}

func (m *Model) PrevTab() int {
	// start at prev element in array
	loopIdx := m.CurrTabId - 1
	totalLoops := 0

	// While we haven't looped over all elements
	for totalLoops < len(*m.lists) {
		// Out of bounds - start from last elem
		if loopIdx < 0 {
			loopIdx = len(*m.lists) - 1
		}

		// Found an elem that can be displayed
		if (*m.lists)[loopIdx].Displayed {
			break
		}

		loopIdx--
		totalLoops++
	}

	m.CurrTabId = loopIdx
	return m.CurrTabId
}
