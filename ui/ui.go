package ui

import (
	"charming-todo/data"
	"charming-todo/ui/components/help"
	"charming-todo/ui/components/list"
	"charming-todo/ui/components/tabs"
	"charming-todo/utils"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

var todoLists []data.TodoList

type Model struct {
	keys      utils.KeyMap
	todoLists *[]data.TodoList

	tabs tabs.Model
	list list.Model
	help help.Model
}

func New() Model {
	todoLists = []data.TodoList{
		data.FetchTodoList(),
		data.TemplateTodoList(),
	}

	m := Model{
		keys:      utils.Keys,
		todoLists: &todoLists,

		tabs: tabs.NewModel(&todoLists),
		list: list.NewModel(&todoLists[0]),
		help: help.NewModel(),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	tabIndex := -1

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Quit app
		if key.Matches(msg, m.keys.Quit) {
			cmd = tea.Quit

			// Create a new Tab
		} else if key.Matches(msg, m.keys.TabNew) {
			tabIndex = m.tabs.AddTab()
			m.tabs.SetCurrSectionId(tabIndex)

			// Close the current Tab
		} else if key.Matches(msg, m.keys.TabClose) {
			tabIndex = m.tabs.CloseCurrentTab()

			// Navigate to next Tab
		} else if key.Matches(msg, m.keys.TabNext) {
			tabIndex = m.tabs.NextTab()

			// Navigate to previous Tab
		} else if key.Matches(msg, m.keys.TabPrev) {
			tabIndex = m.tabs.PrevTab()

			// Navigate to previous Tab
		} else if key.Matches(msg, m.keys.TabSave) {
			m.list.MarkDirty(false)
			// TODO: implement real save

			// Toggle currently selected Todo Item
		} else if key.Matches(msg, m.keys.ItemToggle) {
			m.list.ToggleCurrentItem()
			m.list.MarkDirty(true)

			// Select next todo item
		} else if key.Matches(msg, m.keys.ItemNext) {
			m.list.NextItem()

			// Select previous todo item
		} else if key.Matches(msg, m.keys.ItemPrev) {
			m.list.PrevItem()

			// Add indentation to currently selected Todo Item
		} else if key.Matches(msg, m.keys.ItemAddIndent) {
			m.list.AddIndent()
			m.list.MarkDirty(true)

			// Remove indentation from currently selected Todo Item
		} else if key.Matches(msg, m.keys.ItemRemIndent) {
			m.list.RemIndent()
			m.list.MarkDirty(true)

			// Create new todo item under current selected
		} else if key.Matches(msg, m.keys.ItemNew) {
			m.list.NewItem()
			m.list.MarkDirty(true)

			// Delete the currently selected todo item
		} else if key.Matches(msg, m.keys.ItemDelete) {
			m.list.DeleteItem()
			m.list.MarkDirty(true)

			// Navigate to top item
		} else if key.Matches(msg, m.keys.ItemTop) {
			m.list.NavToTopItem()

			// Navigate to bottom item
		} else if key.Matches(msg, m.keys.ItemBottom) {
			m.list.NavToBottomItem()
		}
	}

	// If tab has been changed and index is within bounds
	if tabIndex > -1 && tabIndex < len(*m.todoLists) {
		m.list.SetTodoList(&todoLists[tabIndex])
	}

	return m, cmd
}

func (m Model) View() string {
	s := strings.Builder{}

	s.WriteString(m.tabs.View())
	s.WriteString(m.list.View())
	s.WriteString("\n")
	s.WriteString("\n")
	s.WriteString(m.help.View())

	return s.String()
}

// TODO: edit selected item's text
// TODO: ctrl+o should create a different list with all TodoLists (hidden or not)
