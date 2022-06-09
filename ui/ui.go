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

			// Toggle currently selected Todo Item
		} else if key.Matches(msg, m.keys.TodoItemToggle) {
			m.list.ToggleCurrentItem()
			m.list.MarkDirty(true)

			// Select next todo item
		} else if key.Matches(msg, m.keys.TodoItemNext) {
			m.list.NextItem()

			// Select previous todo item
		} else if key.Matches(msg, m.keys.TodoItemPrev) {
			m.list.PrevItem()

			// Add indentation to currently selected Todo Item
		} else if key.Matches(msg, m.keys.TodoItemAddIndent) {
			m.list.AddIndent()
			m.list.MarkDirty(true)

			// Remove indentation from currently selected Todo Item
		} else if key.Matches(msg, m.keys.TodoItemRemIndent) {
			m.list.RemIndent()
			m.list.MarkDirty(true)

			// Create new todo item under current selected
		} else if key.Matches(msg, m.keys.TodoItemNew) {
			m.list.NewItem()
			m.list.MarkDirty(true)

			// Delete the currently selected todo item
		} else if key.Matches(msg, m.keys.TodoItemDelete) {
			m.list.DeleteItem()
			m.list.MarkDirty(true)
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
// TODO: go to elem 0 (s)tart
// TODO: go to elem last (b)ottom
