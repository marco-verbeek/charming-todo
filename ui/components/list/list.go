package list

import (
	"charming-todo/data"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	itemStyle          = lipgloss.NewStyle().PaddingLeft(4)
	checkedItemStyle   = lipgloss.NewStyle().Strikethrough(true).StrikethroughSpaces(false).Faint(true)
	checkedColorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff99"))
	selectedColorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#a3cded"))
)

type Model struct {
	todoList   *data.TodoList
	currItemId int
}

func NewModel(todoList *data.TodoList) Model {
	return Model{
		todoList: todoList,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var listAsText string

	for idx, item := range m.todoList.Items {
		indentation := strings.Repeat("   ", item.Indentation)

		var str string
		if item.Checked {
			striked := checkedItemStyle.Render(item.Description)
			checkmark := checkedColorStyle.Render("✓")

			str = fmt.Sprintf("%s%s %s", indentation, checkmark, striked)
		} else {
			str = fmt.Sprintf("%s☐ %s", indentation, item.Description)
		}

		if idx == m.currItemId {
			str = selectedColorStyle.Render("›› " + itemStyle.Render(str))
		} else {
			str = itemStyle.Render("   " + str)
		}

		listAsText += "\n" + str
	}

	return listAsText
}

func (m *Model) SetTodoList(todoList *data.TodoList) {
	m.todoList = todoList
	m.currItemId = 0
}

func (m *Model) ToggleCurrentItem() {
	currentItem := &(*m.todoList).Items[m.currItemId]
	currentItem.Checked = !currentItem.Checked
}

func (m *Model) AddIndent() {
	currentItem := &(*m.todoList).Items[m.currItemId]

	if currentItem.Indentation >= data.MAX_INDENTATION {
		return
	}

	currentItem.Indentation++
}

func (m *Model) RemIndent() {
	currentItem := &(*m.todoList).Items[m.currItemId]

	if currentItem.Indentation <= 0 {
		return
	}

	currentItem.Indentation--
}

func (m *Model) NextItem() {
	nextIdx := m.currItemId + 1

	if nextIdx >= len(m.todoList.Items) {
		nextIdx = 0
	}

	m.currItemId = nextIdx
}

func (m *Model) PrevItem() {
	prevIdx := m.currItemId - 1

	if prevIdx < 0 {
		prevIdx = len(m.todoList.Items) - 1
	}

	m.currItemId = prevIdx
}
